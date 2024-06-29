#!/bin/sh
#go run . -ro ../modules $1

root='D:\Dev\tmp\go\pkg\mod\go.etcd.io\bbolt@v1.3.9\'

go run . -ro ../modules -i golang.org/x/sys/windows $root

