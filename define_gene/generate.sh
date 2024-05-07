#!/bin/sh
#go run . -g=wasm -ro ../golang "${GOROOT}/src/syscall"
goRun(){
  go run . -rt  -o ../golang ${GOROOT}/src/$*
}
#goRun io
#goRun io/fs
#goRun fmt
#goRun context
#goRun time
#goRun bufio
#goRun math
#goRun math/rand
#goRun math/big
#goRun math/cmplx
#goRun math/bits
#goRun mime
#goRun mime/multipart
#goRun os
#goRun hash
#goRun hash/adler32
#goRun hash/crc32
#goRun hash/crc64
#goRun hash/fnv
#goRun hash/maphash
#goRun crypto
#goRun net
#goRun net/textproto
goRun net/url
goRun net/http
