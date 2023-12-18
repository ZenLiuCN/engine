#!/bin/sh
mkdir -p build && \
cmake -G "MinGW Makefiles" \
 -DENABLE_EXTENSION_AUTOLOADING=1 \
 -DENABLE_EXTENSION_AUTOINSTALL=1 \
 -DDUCKDB_EXTENSION_CONFIGS="./.github/config/bundled_extensions.cmake" \
 -DBUILD_SHELL=0 \
 -DBUILD_BENCHMARK=0 \
 -DBUILD_JDBC=0 \
 -DBUILD_TPCH=0 \
 -DBUILD_TPCDS=0 \
 -DBUILD_ODBC=0 \
 -DBUILD_PYTHON=0 \
 -DDISABLE_UNITY=1 \
 -DCMAKE_BUILD_TYPE=Release -B build &&\
cd build &&\
 MAKEFLAGS=-j6 cmake --build . --config Release &&\
cd ../ &&\
mkdir -p lib