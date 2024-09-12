#!/bin/sh
CGO_ENABLED=1 CGO_LDFLAGS="-LD:/Dev/store/duckdb_lib/windows" go build -tags=duckdb_use_lib .
./testdata
