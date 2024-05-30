#!/bin/sh
CGO_LDFLAGS="-L.libs/linux/" go build -tags=duckdb_use_lib,chrome,ducdb,gse,all -o engine -ldflags="-X main.Version=$(git describe --abbrev=0 --tags) -s -w" .
