#!/usr/bin/env bash

set -xeuo pipefail

dir="$(dirname $(dirname $(dirname $(dirname "${0}"))))"
source "$dir/teamcity-bazel-support.sh"

bazel build //pkg/cmd/bazci //pkg/cmd/github-post //pkg/cmd/testfilter --config=ci
BAZEL_BIN=$(bazel info bazel-bin --config=ci)
GO_TEST_JSON_OUTPUT_FILE=/artifacts/test.json.txt
exit_status=0
$BAZEL_BIN/pkg/cmd/bazci/bazci_/bazci -- test --config=ci \
    //pkg/sql/tests:tests_test \
    --test_arg -rsg=5m --test_arg -rsg-routines=8 --test_arg -rsg-exec-timeout=1m \
    --test_timeout 3600 --test_filter 'TestRandomSyntax' \
    --test_sharding_strategy=disabled \
    --test_env=GO_TEST_JSON_OUTPUT_FILE=$GO_TEST_JSON_OUTPUT_FILE || exit_status=$?
process_test_json \
        $BAZEL_BIN/pkg/cmd/testfilter/testfilter_/testfilter \
        $BAZEL_BIN/pkg/cmd/github-post/github-post_/github-post \
        /artifacts \
        $GO_TEST_JSON_OUTPUT_FILE \
        $exit_status
exit $exit_status
