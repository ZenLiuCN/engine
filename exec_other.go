//go:build !windows

package engine

import (
	"os/exec"
	"strings"
)

func Locate(pth string) (string, error) {
	if !strings.Contains(pth, "/") {
		return Lookup(pth)
	}
	return pth, nil
}
func NoWindow(cmd *exec.Cmd) *exec.Cmd {
	return cmd //TODO
}
