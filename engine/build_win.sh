#!/bin/sh
CGO_LDFLAGS="-L.libs/win/"  go build -tags=duckdb_use_lib,chrome,ducdb,gse,all -o engine.exe -ldflags="-s -w" .
