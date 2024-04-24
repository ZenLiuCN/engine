#!/bin/sh
CGO_LDFLAGS="-L.libs/linux/" go build -tags=duckdb_use_lib -o engine -ldflags="-s -w" .
