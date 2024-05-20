//go:build windows

package engine

import (
	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"syscall"
)

func Locate(pth string) (string, error) {
	if !strings.Contains(pth, ":\\") {
		return Lookup(pth)
	}
	return pth, nil
}
func NoWindow(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func init() {
	patch(windows.GetACP())
}

// https://learn.microsoft.com/en-us/windows/win32/intl/code-page-identifiers
func patch(cm uint32) {
	switch cm {
	case 936:
		Encoding = simplifiedchinese.GBK
	}
}
