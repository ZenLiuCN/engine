#!/bin/sh
mkdir -p build && \
cmake -G "MinGW Makefiles" \
 -DENABLE_EXTENSION_AUTOLOADING=1 \
 -DENABLE_EXTENSION_AUTOINSTALL=1 \
 -DBUILD_EXTENSIONS=parquet \
 -DDUCKDB_EXTENSION_CONFIGS="./.github/config/bundled_extensions.cmake" \
 -DBUILD_SHELL=0 \
 -DBUILD_BENCHMARK=0 \
 -DBUILD_JDBC=0 \
 -DBUILD_TPCH=0 \
 -DBUILD_TPCDS=0 \
 -DBUILD_ODBC=0 \
 -DBUILD_PYTHON=0 \
 -DDISABLE_UNITY=1 \
 -DBUILD_AUTOCOMPLETE=1 \
 -DBUILD_HTTPFS=1 \
 -DBUILD_JSON=1 \
 -DBUILD_INET=1 \
 -DBUILD_FTS=1 \
 -DCMAKE_CXX_FLAGS="-DDUCKDB_CUSTOM_PLATFORM=windows_amd64" \
 -DCMAKE_BUILD_TYPE=Release -B build &&\
cd build &&\
 MAKEFLAGS=-j6 cmake --build . --config Release &&\
cd ../ &&\
mkdir -p lib && \
find ./build -type f -name '*.obj' ! -path './extension/visualizer/*' ! -path './extension/tpcds/*' ! -path './extension/sqlsmith/*' |xargs -I {} gcc-ar rvs --thin libduckdb.a {}