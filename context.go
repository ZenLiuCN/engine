package engine

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	execPath   string
	root       string
	executable string
	ext        string
	name       string

	pathSeparator     = string(os.PathSeparator)
	pathListSeparator = string(os.PathListSeparator)
)

func GetExecPath() string {
	return execPath
}
func GetRoot() string {
	return root
}
func GetExecutable() string {
	return executable
}
func GetExt() string {
	return ext
}
func GetName() string {
	return name
}
func GetPathSeparator() string {
	return pathSeparator
}
func GetPathListSeparator() string {
	return pathListSeparator
}
func init() {
	var err error
	execPath, err = os.Executable()
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		execPath, _ = filepath.Abs(os.Args[0])
	}
	root = filepath.Dir(execPath)
	_, executable = filepath.Split(execPath)
	ext = filepath.Ext(execPath)
	name = strings.Replace(executable, ext, "", -1)

}
