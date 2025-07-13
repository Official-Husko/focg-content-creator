#!/bin/bash
# Launch Air from the cmd/focg-content-creator directory
cd "$(dirname "$0")/cmd/focg-content-creator" || exit 1
exec air
