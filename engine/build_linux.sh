#!/bin/sh
CGO_LDFLAGS="-L.libs/linux/" go build -tags=duckdb_use_lib,chrome,ducdb,gse,all -o engine -ldflags="-s -w" .
