// expect is a module port from https://github.com/hymkor/expect
package expect

import (
	"context"
	"errors"
	"fmt"
	"github.com/ZenLiuCN/engine/modules/go/expect/console"
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	escEcho  = "\x1B[49;31;1m"
	escSend  = "\x1B[49;35;1m"
	escSpawn = "\x1B[49;32;1m"
	escEnd   = "\x1B[49;39;1m"
)
var conIn console.Handle
var output = colorable.NewColorableStdout()
var echo = io.Discard

var useStderrOnGetRecentOutput = false

func getRecentOutputByStdoutOrStderr(captureLines int) ([]string, error) {
	for {
		if useStderrOnGetRecentOutput {
			result, err := console.GetRecentOutputByStderr(captureLines)
			return result, err
		}
		result, err := console.GetRecentOutputByStdout(captureLines)
		if err == nil {
			return result, nil
		}
		useStderrOnGetRecentOutput = true
	}
}

type Matching struct {
	Position  int
	Line      string
	Match     string
	PreMatch  string
	PostMatch string
}

var waitGroup sync.WaitGroup

func spawn(newCmd func(string, ...string) *exec.Cmd, args []string, log io.Writer) (int, error) {
	var cmd *exec.Cmd
	for _, s := range args {
		_, _ = fmt.Fprintf(log, "%s\"%s\"%s ", escSpawn, s, escEnd)
	}
	_, _ = fmt.Fprintln(log)
	if len(args) <= 1 {
		cmd = newCmd(args[0])
	} else {
		cmd = newCmd(args[0], args[1:]...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return 0, fmt.Errorf("spawn: %w", err)
	}
	waitGroup.Add(1)
	pid := cmd.Process.Pid
	go func() {
		_ = cmd.Wait()
		waitGroup.Done()
	}()
	return pid, nil
}
func SpawnContext(ctx context.Context, args ...string) (int, error) {
	return spawn(func(s string, s2 ...string) *exec.Cmd {
		return exec.CommandContext(ctx, s, s2...)
	}, args, echo)
}
func Spawn(args ...string) (int, error) {
	return spawn(exec.Command, args, echo)
}
func Wait(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	_, err = process.Wait()
	return err
}
func Kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Kill()
}
func Send(str string, wait int) {
	fmt.Fprintf(echo, "%s%s%s", escSend, strings.Replace(str, "\r", "\n", -1), escEnd)
	for _, c := range str {
		console.Rune(conIn, c)
		if wait > 0 {
			time.Sleep(time.Second * time.Duration(wait) / 1000)
		}
	}
}
func Shot(n int) ([]string, error) {
	if !useStderrOnGetRecentOutput {
		result, err := console.GetRecentOutputByStdout(n)
		if err == nil {
			return result, nil
		}
		useStderrOnGetRecentOutput = true
	}
	return console.GetRecentOutputByStderr(n)
}
func expect(ctx context.Context, keywords []string, timeout time.Duration, captureLines int) (int, *Matching, error) {
	tick := time.NewTicker(time.Millisecond * 100)
	defer tick.Stop()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for {
		outputs, err := getRecentOutputByStdoutOrStderr(captureLines)
		if err != nil {
			return -1, nil, fmt.Errorf("expect: %w", err)
		}
		for _, out := range outputs {
			for i, str := range keywords {
				if pos := strings.Index(out, str); pos >= 0 {
					return i, &Matching{
						Position:  pos,
						Line:      out,
						Match:     out[pos : pos+len(str)],
						PreMatch:  out[:pos],
						PostMatch: out[pos+len(str):],
					}, nil
				}
			}
		}
		select {
		case <-ctx.Done():
			return -1, nil, fmt.Errorf("expect: %w", ctx.Err())
		case <-timer.C:
			return -1, nil, context.DeadlineExceeded
		case <-tick.C:
		}
	}
}

const (
	ErrnoExpectGetRecentOutput = -1
	ErrnoExpectTimeOut         = -2
	ErrnoExpectContextDone     = -3
)

func Expect(timeout int, captureLines int, keywords ...string) (int, *Matching) {
	ctx, cc := context.WithTimeout(context.Background(), time.Hour)
	defer cc()
	return ExpectContext(ctx, timeout, captureLines, keywords...)
}
func ExpectContext(ctx context.Context, timeout int, captureLines int, keywords ...string) (rc int, info *Matching) {
	var err error
	var to = time.Duration(timeout) * time.Millisecond
	if to < time.Millisecond*100 {
		to = time.Second
	}
	if captureLines <= 0 {
		captureLines = 1
	}
	rc, info, err = expect(ctx, keywords, to, captureLines)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			rc = ErrnoExpectContextDone
		} else if errors.Is(err, context.DeadlineExceeded) {
			rc = ErrnoExpectTimeOut
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			rc = ErrnoExpectGetRecentOutput
		}
	}
	return
}
func Sleep(value int) {
	time.Sleep(time.Second * time.Duration(value))
}
func USleep(value int) {
	time.Sleep(time.Millisecond * time.Duration(value))
}
func Echo(L any) {
	switch t := L.(type) {
	case bool:
		if t {
			echo = output
		} else {
			echo = io.Discard
		}
	default:
		_, _ = fmt.Fprintf(output, "%s%s%s\r\n", escEcho, t, escEnd)
	}

}
