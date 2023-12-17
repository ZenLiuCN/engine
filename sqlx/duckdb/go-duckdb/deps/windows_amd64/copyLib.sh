#!/bin/sh
find ./build -name '*.a'|xargs -I {} cp {} -t ./lib
