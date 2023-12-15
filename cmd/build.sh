#!/bin/sh
go build -o engine.exe -ldflags="-s -w" .
