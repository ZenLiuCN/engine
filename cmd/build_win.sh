#!/bin/sh
CGO_LDFLAGS="-L.libs/win/"  go build -tags=duckdb_use_lib -o engine.exe -ldflags="-s -w" .
