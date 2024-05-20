package engine

import (
	"bytes"
	"sync"
)

// !! pool not effective

func Get() *Engine {
	return NewEngine()
}
func Put(e *Engine) {

}

var (
	buffers = sync.Pool{New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 64))
	}}
)

func GetBytesBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}
func PutBytesBuffer(buf *bytes.Buffer) {
	buf.Reset()
	if buf.Cap() > 4*1024 {
		return
	}
	buffers.Put(buf)
}
