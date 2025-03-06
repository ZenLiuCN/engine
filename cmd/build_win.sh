#!/bin/sh
CGO_LDFLAGS="-L$(pwd)/.libs/win/" \
go build  -tags="all duckdb_use_lib" -o engine.exe \
-ldflags="-X main.Version=$(Git describe --abbrev=0 --tags) -s -w" . && \
upx --lzma engine.exe

# go list -f '| {{.ImportPath}} | {{join .ImportedBuildTags ", "}} | {{join .GoFiles ", "}} |'  -ldflags="-X main.Version=$(Git describe --abbrev=0 --tags)"   -tags="all duckdb_use_lib" -deps ./... |  awk 'BEGIN {print "| Package | Tags | Files |\n| --- | --- | --- |"} {print}'
# go list -deps -f '{{.ImportPath}} | {{join .ImportedBuildTags ", "}}'    -tags="all duckdb_use_lib" > temp.txt
# go list -deps -f '{{range .GoFiles}}{{$.ImportPath}}: {{.}} | BuildTags: {{context.BuildTags}}{{"\n"}}{{end}}' -tags="duckdb_use_lib all" | sort >temp.txt
