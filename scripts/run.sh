#!/usr/bin/env bash
# run.sh is a thin wrapper around the harness. Pass any flags you would pass to
# the harness; arguments after "--" are forwarded directly to `go test`.
set -euo pipefail
exec go run ./cmd/harness "$@"
