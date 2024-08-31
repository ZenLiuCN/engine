package console

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

type Handle windows.Handle

func New() Handle {
	return Handle(windows.Stdin)
}

func (handle Handle) Close() error {
	return nil
}

func (handle Handle) GetConsoleMode() uint32 {
	var mode uint32
	windows.GetConsoleMode(windows.Handle(handle), &mode)
	return mode
}

func (handle Handle) SetConsoleMode(flag uint32) {
	windows.SetConsoleMode(windows.Handle(handle), flag)
}

var flushConsoleInputBuffer = kernel32.NewProc("FlushConsoleInputBuffer")

func (handle Handle) FlushConsoleInputBuffer() error {
	status, _, err := flushConsoleInputBuffer.Call(uintptr(handle))
	if status != 0 {
		return nil
	} else {
		return fmt.Errorf("FlushConsoleInputBuffer: %s", err)
	}
}

var getNumberOfConsoleInputEvents = kernel32.NewProc("GetNumberOfConsoleInputEvents")

func (handle Handle) GetNumberOfEvent() (int, error) {
	var count uint32 = 0
	status, _, err := getNumberOfConsoleInputEvents.Call(uintptr(handle),
		uintptr(unsafe.Pointer(&count)))
	if status != 0 {
		return int(count), nil
	} else {
		return 0, fmt.Errorf("GetNumberOfConsoleInputEvents: %w", err)
	}
}

var waitForSingleObject = kernel32.NewProc("WaitForSingleObject")

func (handle Handle) WaitForSingleObject(msec uintptr) (uintptr, error) {
	status, _, err := waitForSingleObject.Call(uintptr(handle), msec)
	if err != nil {
		return status, fmt.Errorf("WaitForSingleObject: %w", err)
	}
	return status, nil
}

var readConsoleInput = kernel32.NewProc("ReadConsoleInputW")

type InputRecord struct {
	EventType uint16
	_         uint16
	Info      [8]uint16
}

func (handle Handle) Read(events []InputRecord) uint32 {
	var n uint32
	readConsoleInput.Call(
		uintptr(windows.Stdin),
		uintptr(unsafe.Pointer(&events[0])),
		uintptr(len(events)),
		uintptr(unsafe.Pointer(&n)))
	return n
}

type KeyEventRecord struct {
	KeyDown         int32
	RepeatCount     uint16
	VirtualKeyCode  uint16
	VirtualScanCode uint16
	UnicodeChar     uint16
	ControlKeyState uint32
}

func (e *InputRecord) KeyEvent() *KeyEventRecord {
	return (*KeyEventRecord)(unsafe.Pointer(&e.Info[0]))
}

type MouseEventRecord struct {
	X          int16
	Y          int16
	Button     uint32
	ControlKey uint32
	Event      uint32
}

func (e *InputRecord) MouseEvent() *MouseEventRecord {
	return (*MouseEventRecord)(unsafe.Pointer(&e.Info[0]))
}

func (m MouseEventRecord) String() string {
	return fmt.Sprintf("X:%d,Y:%d,Button:%d,ControlKey:%d,Event:%d",
		m.X, m.Y, m.Button, m.ControlKey, m.Event)
}

type windowBufferSizeRecord struct {
	X int16
	Y int16
}

func (e *InputRecord) ResizeEvent() (int16, int16) {
	p := (*windowBufferSizeRecord)(unsafe.Pointer(&e.Info[0]))
	return p.X, p.Y
}

var writeConsoleInput = kernel32.NewProc("WriteConsoleInputW")

func (handle Handle) Write(events []InputRecord) uint32 {
	var count uint32
	writeConsoleInput.Call(uintptr(handle), uintptr(unsafe.Pointer(&events[0])), uintptr(len(events)), uintptr(unsafe.Pointer(&count)))

	return count
}
