package main

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	p := os.ExpandEnv("${GOROOT}/src")
	b := new(bytes.Buffer)
	fn.Panic(filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if strings.Contains(path, "cmd") ||
				strings.HasSuffix(path, "go") ||
				strings.Contains(path, "debug") ||
				strings.Contains(path, "runtime") ||
				strings.Contains(path, "internal") ||
				strings.Contains(path, "plugin") ||
				strings.Contains(path, "embed") ||
				strings.Contains(path, "database") ||
				strings.Contains(path, "vendor") ||
				strings.Contains(path, "testing") ||
				strings.Contains(path, "syscall\\js") ||
				strings.Contains(path, "tzdata") ||
				strings.Contains(path, "testdata") {
				return filepath.SkipDir
			}
			e := fn.Panic1(os.ReadDir(path))
			cnt := 0
			for _, entry := range e {
				if !entry.IsDir() {
					cnt++
				}
			}
			if cnt > 0 {
				b.WriteByte('\n')
				fmt.Fprintf(b, "goRun %s", fn.Panic1(filepath.Rel(p, path)))
			}
		}
		return nil
	}))
	b.WriteByte('\n')
	b.WriteTo(os.Stdout)
}
