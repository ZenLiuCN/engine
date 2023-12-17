#!/bin/sh
find ./build -type f -name '*.obj' ! -path './extension/visualizer/*' ! -path './extension/tpcds/*' ! -path './extension/sqlsmith/*' |xargs -I {} gcc-ar rvs --thin ./lib/libduckdb_all.a {}
