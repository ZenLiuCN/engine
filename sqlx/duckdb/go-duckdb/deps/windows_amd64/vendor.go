// Package windows_amd64 is required to provide support for vendoring modules
// DO NOT REMOVE
/**
#!/bin/sh
# sample script to build lib
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
 -DCMAKE_BUILD_TYPE=Release -B build &&\
cd build &&\
 MAKEFLAGS=-j8 cmake --build . --config Release &&\
cd ../ &&\
mkdir -p lib && \
find ./build -name '*.obj'|xargs cp {} -t ./lib && \
cd lib && \
gcc-ar rvs libduckdb.a *.obj
*/
package windows_amd64
