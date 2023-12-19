#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go generate ./pkg/ent
go mod tidy

exit 0