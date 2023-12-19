#!/bin/sh
# build postgres_scanner
# set(CMAKE_SHARED_LINKER_FLAGS  "-static-libgcc -static-libstdc++") # add to cmakelists
cd postgres_scanner &&\
cd postgres &&\
./configure \
  CC=gcc CXX=g++ \
  --host=x86_64-w64-mingw32 \
    --without-llvm        \
    --without-icu         \
    --without-tcl         \
    --without-perl        \
    --without-python      \
    --without-gssapi      \
    --without-pam         \
    --without-bsd-auth    \
    --without-ldap        \
    --without-bonjour     \
    --without-selinux     \
    --without-systemd     \
    --without-readline    \
    --without-libxml      \
    --without-libxslt     \
    --without-zlib        \
    --without-lz4         \
    --without-openssl > /dev/null
cd ../ &&\
mkdir -p build/release && \
OPENSSL_ROOT_DIR='D:\Dev\env\mingw\pkgconfig\openssl\' \
OPENSSL_USE_STATIC_LIBS=true \
cmake -G "MinGW Makefiles" \
-DEXTENSION_STATIC_BUILD=1 \
-DDUCKDB_EXTENSION_NAMES="postgres_scanner" \
-DDUCKDB_EXTENSION_POSTGRES_SCANNER_PATH="D:/dev/store/postgres_scanner/" \
-DDUCKDB_EXTENSION_POSTGRES_SCANNER_SHOULD_LINK=0 \
-DDUCKDB_EXTENSION_POSTGRES_SCANNER_LOAD_TESTS=0 \
-DDUCKDB_EXTENSION_POSTGRES_SCANNER_TEST_PATH="D:/dev/store/postgres_scanner/test/sql" \
-DVCPKG_MANIFEST_DIR='D:/dev/store/postgres_scanner/' \
-DCMAKE_BUILD_TYPE=Release \
-S ./duckdb/ -B build/release && \
cmake --build build/release --config Release &&\
cd ../duckdb_mysql &&\
