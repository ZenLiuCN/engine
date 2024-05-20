#!/bin/sh
CGO_LDFLAGS="-L.libs/win/"  go build -tags=duckdb_use_lib,all -o engine.exe -ldflags="-s -w" . && \
upx --lzma engine.exe