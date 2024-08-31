package console

const FOCUS_EVENT = 16
const KEY_EVENT = 1
const MENU_EVENT = 8
const MOUSE_EVENT = 2
const WINDOW_BUFFER_SIZE_EVENT = 4
const ENABLE_WINDOW_INPUT = 8
const WAIT_ABANDONED = 128
const WAIT_OBJECT_0 = 0
const WAIT_TIMEOUT = 258
const WAIT_FAILED = -1
const ENABLE_MOUSE_INPUT = 16

func SendKeyEvent(handle Handle, events ...*KeyEventRecord) uint32 {
	records := make([]InputRecord, len(events))
	for i, e := range events {
		records[i].EventType = KEY_EVENT
		keyEvent := records[i].KeyEvent()
		*keyEvent = *e
		if keyEvent.RepeatCount <= 0 {
			keyEvent.RepeatCount = 1
		}
	}
	return handle.Write(records[:])
}

func Rune(handle Handle, c rune) uint32 {
	return SendKeyEvent(handle,
		&KeyEventRecord{UnicodeChar: uint16(c), KeyDown: 1},
		&KeyEventRecord{UnicodeChar: uint16(c)},
	)
}

func String(handle Handle, s string) {
	for _, c := range s {
		Rune(handle, c)
	}
}

func VirtualKey(handle Handle, v int) uint32 {
	return SendKeyEvent(handle,
		&KeyEventRecord{VirtualKeyCode: uint16(v), KeyDown: 1},
		&KeyEventRecord{VirtualKeyCode: uint16(v)},
	)
}
