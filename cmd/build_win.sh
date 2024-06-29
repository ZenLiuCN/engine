#!/bin/sh
CGO_LDFLAGS="-L$(pwd)/../../Loader/.libs/win/" \
go build -tags=duckdb_use_lib,all -o engine.exe \
-ldflags="-X main.Version=$(Git describe --abbrev=0 --tags) -s -w" . && \
upx --lzma engine.exe
