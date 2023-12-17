//go:build !duckdb_use_lib && windows && amd64

package duckdb

/*
#cgo LDFLAGS: -lduckdb_static -lws2_32
#cgo LDFLAGS: -lduckdb_fsst -lduckdb_fmt -lduckdb_pg_query -lduckdb_re2 -lduckdb_miniz -lduckdb_utf8proc -lduckdb_hyperloglog -lduckdb_fastpforlib -lduckdb_mbedtls
#cgo LDFLAGS: -linet_extension
#cgo LDFLAGS: -lautocomplete_extension
#cgo LDFLAGS: -ltpch_extension
#cgo LDFLAGS: -lfts_extension
#cgo LDFLAGS: -licu_extension
#cgo LDFLAGS: -ljson_extension
#cgo LDFLAGS: -lexcel_extension
#cgo LDFLAGS: -lparquet_extension -lduckdb_static
#cgo LDFLAGS: -Wl,-Bstatic -lstdc++ -Wl,-Bstatic -lpthread -lm -L${SRCDIR}/deps/windows_amd64
#include <duckdb.h>
*/
import "C"
