// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/cockroach/pkg/util/timeutil"
	"github.com/pkg/errors"
)

func registerBackup(r *registry) {
	backup2TBSpec := makeClusterSpec(10)
	r.Add(testSpec{
		Name:    fmt.Sprintf("backup2TB/%s", backup2TBSpec),
		Cluster: backup2TBSpec,
		Run: func(ctx context.Context, t *test, c *cluster) {
			rows := 65104166
			dest := c.name

			if local {
				rows = 100
				dest += fmt.Sprintf("%d", timeutil.Now().UnixNano())
			}

			c.Put(ctx, cockroach, "./cockroach")
			c.Put(ctx, workload, "./workload")
			c.Start(ctx, t)
			c.Run(ctx, c.Node(1), "./workload", "fixtures", "import", "bank",
				"--db=bank", "--payload-bytes=10240", "--ranges=0",
				fmt.Sprintf("--rows=%d", rows), "--seed=1", "{pgurl:1}")

			// NB: without this delay, the BACKUP operation sometimes claims that
			// bank.bank doesn't exist, probably due to some gossip propagation
			// delay.
			//
			// See https://github.com/cockroachdb/cockroach/issues/36841.
			for i := 0; i < 5; i++ {
				_, err := c.Conn(ctx, 1).ExecContext(ctx, "SELECT * FROM bank.bank LIMIT 1")
				if err != nil {
					c.l.Printf("%s", err)
					time.Sleep(time.Second)
					continue
				}
				c.l.Printf("found the table")
				break
			}

			m := newMonitor(ctx, c)
			m.Go(func(ctx context.Context) error {
				t.Status(`running backup`)
				c.Run(ctx, c.Node(1), `./cockroach sql --insecure -e "
				BACKUP bank.bank TO 'gs://cockroachdb-backup-testing/`+dest+`'"`)
				return nil
			})
			m.Wait()
		},
	})

	// backupTPCC continuously runs TPCC, takes a full backup after some time,
	// and incremental after more time. It then restores the two backups and
	// verifies them with a fingerprint.
	r.Add(testSpec{
		Name:    `backupTPCC`,
		Cluster: makeClusterSpec(3),
		Timeout: 1 * time.Hour,
		Run: func(ctx context.Context, t *test, c *cluster) {
			c.Put(ctx, cockroach, "./cockroach")
			c.Put(ctx, workload, "./workload")
			c.Start(ctx, t)
			conn := c.Conn(ctx, 1)

			duration := 5 * time.Minute
			if local {
				duration = 5 * time.Second
			}
			warehouses := 1
			backupDir := "gs://cockroachdb-backup-testing/" + c.name
			fullDir := backupDir + "/full"
			incDir := backupDir + "/inc"

			t.Status(`workload initialization`)
			cmd := fmt.Sprintf(
				"./workload init tpcc --warehouses=%d {pgurl:1-%d}",
				warehouses, c.nodes,
			)
			c.Run(ctx, c.Node(1), cmd)

			m := newMonitor(ctx, c)
			m.Go(func(ctx context.Context) error {
				_, err := conn.ExecContext(ctx, `
					CREATE DATABASE restore_full;
					CREATE DATABASE restore_inc;
				`)
				return err
			})
			m.Wait()

			t.Status(`run tpcc`)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			cmdDone := make(chan error)
			go func() {
				cmd := fmt.Sprintf(
					"./workload run tpcc --warehouses=%d {pgurl:1-%d}",
					warehouses, c.nodes,
				)

				cmdDone <- c.RunE(ctx, c.Node(1), cmd)
			}()

			select {
			case <-time.After(duration):
			case <-ctx.Done():
				return
			}

			t.Status(`full backup`)
			// Use a time slightly in the past to avoid "cannot specify timestamp in the future" errors.
			tFull := fmt.Sprint(timeutil.Now().Add(time.Second * -2).UnixNano())
			m = newMonitor(ctx, c)
			m.Go(func(ctx context.Context) error {
				_, err := conn.ExecContext(ctx,
					`BACKUP tpcc.* TO $1 AS OF SYSTEM TIME `+tFull,
					fullDir,
				)
				return err
			})
			m.Wait()

			t.Status(`continue tpcc`)
			select {
			case <-time.After(duration):
			case <-ctx.Done():
				return
			}

			t.Status(`incremental backup`)
			tInc := fmt.Sprint(timeutil.Now().Add(time.Second * -2).UnixNano())
			m = newMonitor(ctx, c)
			m.Go(func(ctx context.Context) error {
				_, err := conn.ExecContext(ctx,
					`BACKUP tpcc.* TO $1 AS OF SYSTEM TIME `+tInc+` INCREMENTAL FROM $2`,
					incDir,
					fullDir,
				)
				if err != nil {
					return err
				}

				// Backups are done, make sure workload is still running.
				select {
				case err := <-cmdDone:
					// Workload exited before it should have.
					return err
				default:
					return nil
				}
			})
			m.Wait()

			m = newMonitor(ctx, c)
			m.Go(func(ctx context.Context) error {
				t.Status(`restore full`)
				if _, err := conn.ExecContext(ctx,
					`RESTORE tpcc.* FROM $1 WITH into_db='restore_full'`,
					fullDir,
				); err != nil {
					return err
				}

				t.Status(`restore incremental`)
				if _, err := conn.ExecContext(ctx,
					`RESTORE tpcc.* FROM $1, $2 WITH into_db='restore_inc'`,
					fullDir,
					incDir,
				); err != nil {
					return err
				}

				t.Status(`fingerprint`)
				fingerprint := func(db string, asof string) (string, error) {
					var b strings.Builder

					var tables []string
					rows, err := conn.QueryContext(ctx, fmt.Sprintf("SHOW TABLES FROM %s", db))
					if err != nil {
						return "", err
					}
					defer rows.Close()
					for rows.Next() {
						var name string
						if err := rows.Scan(&name); err != nil {
							return "", err
						}
						tables = append(tables, name)
					}

					for _, table := range tables {
						fmt.Fprintf(&b, "table %s\n", table)
						query := fmt.Sprintf("SHOW EXPERIMENTAL_FINGERPRINTS FROM TABLE %s.%s", db, table)
						if asof != "" {
							query = fmt.Sprintf("SELECT * FROM [%s] AS OF SYSTEM TIME %s", query, asof)
						}
						rows, err = conn.QueryContext(ctx, query)
						if err != nil {
							return "", err
						}
						defer rows.Close()
						for rows.Next() {
							var name, fp string
							if err := rows.Scan(&name, &fp); err != nil {
								return "", err
							}
							fmt.Fprintf(&b, "%s: %s\n", name, fp)
						}
					}

					return b.String(), rows.Err()
				}

				tpccFull, err := fingerprint("tpcc", tFull)
				if err != nil {
					return err
				}
				tpccInc, err := fingerprint("tpcc", tInc)
				if err != nil {
					return err
				}
				restoreFull, err := fingerprint("restore_full", "")
				if err != nil {
					return err
				}
				restoreInc, err := fingerprint("restore_inc", "")
				if err != nil {
					return err
				}

				if tpccFull != restoreFull {
					return errors.Errorf("got %s, expected %s", restoreFull, tpccFull)
				}
				if tpccInc != restoreInc {
					return errors.Errorf("got %s, expected %s", restoreInc, tpccInc)
				}

				return nil
			})
			m.Wait()
		},
	})
}
