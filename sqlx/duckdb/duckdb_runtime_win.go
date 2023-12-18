//go:build windows && amd64

package duckdb

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
)

var (
	//go:embed libgcc_s_seh-1.dll
	seh []byte
	//go:embed libstdc++-6.dll
	stdcxx []byte
	//go:embed libwinpthread-1.dll
	pthread []byte
)

func CheckDll() {

	var err error
	execPath, err := os.Executable()
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	root := filepath.Dir(execPath)
	if _, err = os.Stat(filepath.Join(root, "libstdc++-6.dll")); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(filepath.Join(root, "libstdc++-6.dll"), stdcxx, 0700)

	}
	if _, err = os.Stat(filepath.Join(root, "libwinpthread-1.dll")); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(filepath.Join(root, "libwinpthread-1.dll"), pthread, 0700)

	}
	if _, err = os.Stat(filepath.Join(root, "libgcc_s_seh-1.dll")); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(filepath.Join(root, "libgcc_s_seh-1.dll"), seh, 0700)
	}
}
