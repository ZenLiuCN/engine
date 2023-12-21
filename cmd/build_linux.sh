#!/bin/sh
CGO_ENABLE=1 go build -o engine -ldflags="-s -w" .
