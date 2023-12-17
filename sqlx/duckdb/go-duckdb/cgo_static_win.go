//go:build !duckdb_use_lib && windows && amd64

package duckdb

/*
#cgo LDFLAGS: -lduckdb_static -lws2_32 -lduckdb_fsst -lduckdb_fmt -lduckdb_pg_query -lduckdb_re2 -lduckdb_miniz -lduckdb_utf8proc -lduckdb_hyperloglog -lduckdb_fastpforlib -lduckdb_mbedtls -lparquet_extension -lduckdb_static -Wl,-Bstatic -lstdc++ -Wl,-Bstatic -lpthread -lm -L${SRCDIR}/deps/windows_amd64
#include <duckdb.h>
*/
import "C"
