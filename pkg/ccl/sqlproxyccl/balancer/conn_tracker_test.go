// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package balancer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/stop"
	"github.com/stretchr/testify/require"
)

func TestConnTracker(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	ctx := context.Background()
	stopper := stop.NewStopper()
	defer stopper.Stop(ctx)

	tracker, err := NewConnTracker(ctx, stopper, nil /* timeSource */)
	require.NoError(t, err)

	tenantID := roachpb.MakeTenantID(20)
	sa := &ServerAssignment{addr: "127.0.0.10:8090", owner: &testConnHandle{}}

	// Run twice for idempotency.
	tracker.registerAssignment(tenantID, sa)
	tracker.registerAssignment(tenantID, sa)

	connsMap := tracker.GetConnsMap(tenantID)
	require.Len(t, connsMap, 1)
	h, ok := connsMap[sa.Addr()]
	require.True(t, ok)
	require.Equal(t, []ConnectionHandle{sa.Owner()}, h)

	tenantIDs := tracker.getTenantIDs()
	require.Len(t, tenantIDs, 1)
	require.Equal(t, tenantID, tenantIDs[0])

	// Non-existent.
	connsMap = tracker.GetConnsMap(roachpb.MakeTenantID(42))
	require.Empty(t, connsMap)

	// Run twice for idempotency.
	tracker.unregisterAssignment(tenantID, sa)
	tracker.unregisterAssignment(tenantID, sa)

	// Once the assignment gets unregistered, we shouldn't return that tenant
	// since there are no active connections.
	tenantIDs = tracker.getTenantIDs()
	require.Empty(t, tenantIDs)

	// Ensure methods are thread-safe.
	var wg sync.WaitGroup
	const clients = 50
	wg.Add(clients)
	for i := 0; i < clients; i++ {
		go func() {
			defer wg.Done()
			tenantID := roachpb.MakeTenantID(uint64(1 + rand.Intn(5)))
			sa := &ServerAssignment{
				addr:  fmt.Sprintf("127.0.0.10:%d", rand.Intn(5)),
				owner: &testConnHandle{},
			}
			tracker.registerAssignment(tenantID, sa)
			time.Sleep(250 * time.Millisecond)
			tracker.unregisterAssignment(tenantID, sa)
		}()
	}

	wg.Wait()

	// Ensure that once all clients have returned, the tenant entries are empty.
	tenantIDs = tracker.getTenantIDs()
	require.Empty(t, tenantIDs)
	for _, entry := range tracker.mu.tenants {
		require.Empty(t, entry.mu.assignments)
	}
}

func TestConnTracker_GetConnsMap(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)

	ctx := context.Background()
	stopper := stop.NewStopper()
	defer stopper.Stop(ctx)

	tracker, err := NewConnTracker(ctx, stopper, nil /* timeSource */)
	require.NoError(t, err)

	makeConn := func(tenID int, addr string) (roachpb.TenantID, *ServerAssignment) {
		return roachpb.MakeTenantID(uint64(tenID)), &ServerAssignment{
			addr:  addr,
			owner: &testConnHandle{},
		}
	}

	tenant10, sa1 := makeConn(10, "127.0.0.10:1010")
	_, sa2 := makeConn(10, "127.0.0.10:1020")
	tracker.registerAssignment(tenant10, sa1)
	tracker.registerAssignment(tenant10, sa2)

	// Ensure that map contains two handles for tenant 10, one per pod.
	initialConnsMap := tracker.GetConnsMap(tenant10)
	require.Equal(t, map[string][]ConnectionHandle{
		sa1.Addr(): {sa1.Owner()},
		sa2.Addr(): {sa2.Owner()},
	}, initialConnsMap)

	// Add a new assignment, and unregister after.
	tenant20, sa3 := makeConn(20, "127.0.0.10:1020")
	tracker.registerAssignment(tenant20, sa3)
	tracker.unregisterAssignment(tenant20, sa3)

	// Ensure that tenants with no connections do not show up.
	connsMap := tracker.GetConnsMap(tenant20)
	require.Empty(t, connsMap)

	// Add a new assignment for tenant 10.
	_, sa4 := makeConn(10, "127.0.0.10:1020")
	tracker.registerAssignment(tenant10, sa4)

	// Existing initialConnsMap does not change. This shows snapshotting.
	require.Equal(t, map[string][]ConnectionHandle{
		sa1.Addr(): {sa1.Owner()},
		sa2.Addr(): {sa2.Owner()},
	}, initialConnsMap)

	// Fetch again, and we should have the updated entry.
	connsMap = tracker.GetConnsMap(tenant10)
	require.Equal(t, map[string][]ConnectionHandle{
		sa1.Addr(): {sa1.Owner()},
		sa2.Addr(): {sa2.Owner(), sa4.Owner()},
	}, connsMap)

	// Disconnect everything.
	tracker.unregisterAssignment(tenant10, sa1)
	tracker.unregisterAssignment(tenant10, sa2)
	tracker.unregisterAssignment(tenant10, sa4)
	connsMap = tracker.GetConnsMap(tenant10)
	require.Empty(t, connsMap)
}

func TestTenantEntry(t *testing.T) {
	defer leaktest.AfterTest(t)()

	entry := newTenantEntry()

	// Create a ServerAssignment without a tracker since that's not necessary
	// here.
	const addr = "10.0.0.1:12345"
	h1 := &testConnHandle{}
	s := &ServerAssignment{addr: addr, owner: h1}

	// Add a new assignment.
	require.True(t, entry.addAssignment(s))
	require.False(t, entry.addAssignment(s))
	require.Equal(t, 1, entry.getConnsCount())
	connsMap := entry.getConnsMap()
	require.Len(t, connsMap, 1)
	arr, ok := connsMap[addr]
	require.True(t, ok)
	require.Equal(t, []ConnectionHandle{h1}, arr)

	// Remove the assignment.
	require.True(t, entry.removeAssignment(s))
	require.False(t, entry.removeAssignment(s))
	require.Equal(t, 0, entry.getConnsCount())
	require.Empty(t, entry.getConnsMap())

	// Show that connsMap is a snapshot/copy.
	require.Len(t, connsMap, 1)
}
