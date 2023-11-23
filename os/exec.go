package os

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding"
	"io"
	"os/exec"
	"path"
	"strings"
	"time"
)

func fixPath(p string) string {
	return path.Join(strings.Split(p, pathSeparator)...)
}

type ExecOption struct {
	Sleep    time.Duration
	Optional bool
	Await    bool
	*ProcOption
}
type ProcOption struct {
	Cmd        string
	Args       []string
	WorkingDir string
	ShowWindow bool
	PathPatch  bool
}
type SubProcess struct {
	*exec.Cmd
	ProcStdout io.Reader
	ProcStderr io.Reader
	ProcStdin  io.Writer
}

func Lookup(cmd string) (p string, err error) {
	p, err = exec.LookPath(cmd)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s locate %s", cmd, err.Error()))
		return
	}
	return
}
func Execute(opt *ExecOption) (err error) {
	if len(opt.Cmd) == 0 {
		if opt.Sleep > 0 {
			time.Sleep(opt.Sleep)
			return nil
		}
		if !opt.Optional {
			panic("cmd must not empty")
		}
		return errors.New("cmd must not empty")
	}
	var pth = opt.Cmd
	if strings.ContainsAny(pth, "@%$") {
		pth = expand(pth)
	}
	pth, err = Locate(pth)
	if opt.Optional && err != nil {
		panic(err)
	} else if err != nil {
		return
	}
	if opt.PathPatch && opt.Args != nil {
		for i, p := range opt.Args {
			if strings.ContainsAny(p, "@%$") {
				opt.Args[i] = expand(p)
			}
		}
	}
	cmd := exec.Command(pth, opt.Args...)
	if !opt.ShowWindow {
		cmd = NoWindow(cmd)
	}
	if opt.WorkingDir != "" {
		cmd.Dir = opt.WorkingDir
	}
	var errs bytes.Buffer
	cmd.Stderr = &errs
	if opt.Await {
		if err = cmd.Run(); err != nil {
			if err = ToNativeConsole(&errs); err != nil {
				panic(err)
			}
			err = fmt.Errorf("error execute: \n %s %s \n erro:%s \n output:%s ",
				cmd.Path,
				strings.Join(cmd.Args, " "),
				err.Error(),
				errs.String(),
			)
			if opt.Optional {
				panic(err)
			}
			return
		}
	} else {
		err = cmd.Start()
	}
	if err = FromNativeConsole(&errs); err != nil {
		panic(err)
	}

	if opt.Sleep > 0 {
		time.Sleep(opt.Sleep)
	}
	return nil
}
func OpenProc(opt *ProcOption) (proc *SubProcess) {
	if len(opt.Cmd) == 0 {
		return
	}
	cmd, err := BuildCommand(opt)
	if err != nil {
		panic(err)
	}
	errOut := new(bytes.Buffer)
	stdOut := new(bytes.Buffer)
	stdIn := new(bytes.Buffer)
	cmd.Stderr = errOut
	cmd.Stdout = stdOut
	cmd.Stdin = stdIn
	return &SubProcess{
		Cmd:        cmd,
		ProcStdout: stdOut,
		ProcStderr: errOut,
		ProcStdin:  stdIn,
	}
}
func BuildCommand(opt *ProcOption) (cmd *exec.Cmd, err error) {
	var pth = opt.Cmd
	if strings.ContainsAny(pth, "@%$") {
		pth = expand(pth)
	}
	if pth, err = Locate(pth); err != nil {
		return
	}
	if opt.PathPatch && opt.Args != nil {
		for i, p := range opt.Args {
			if strings.ContainsAny(p, "@%$") {
				opt.Args[i] = expand(p)
			}
		}
	}
	cmd = exec.Command(pth, opt.Args...)
	if !opt.ShowWindow {
		cmd = NoWindow(cmd)
	}
	if opt.WorkingDir != "" {
		cmd.Dir = opt.WorkingDir
	}
	return
}

var (
	Encoding encoding.Encoding
)

func ToNativeConsole(b *bytes.Buffer) error {
	if b.Len() == 0 || Encoding == nil {
		return nil
	}
	bx, err := Encoding.NewEncoder().Bytes(b.Bytes())
	if err != nil {
		return err
	}
	b.Reset()
	b.Write(bx)
	return nil
}
func FromNativeConsole(b *bytes.Buffer) error {
	if b.Len() == 0 || Encoding == nil {
		return nil
	}
	bx, err := Encoding.NewDecoder().Bytes(b.Bytes())
	if err != nil {
		return err
	}
	b.Reset()
	b.Write(bx)
	return nil
}
