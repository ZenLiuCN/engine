#!/bin/sh
CGO_LDFLAGS="-L.libs/linux/" GOOS=linux\
go build -tags="duckdb_use_lib chrome duckdb gse all" -o engine -ldflags="-X main.Version=$(git describe --abbrev=0 --tags) -s -w" .
