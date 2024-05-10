#!/bin/sh
CGO_LDFLAGS="-L.libs/win/"  go build -tags=duckdb_use_lib,all,duckdb,chrome,gse -o engine.exe -ldflags="-s -w" . && upx_max engine.exe
