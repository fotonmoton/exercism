package circular

import (
	"container/list"
	"errors"
	"io"
)

type Bufferable interface {
	io.ByteReader
	io.ByteWriter
	Overwrite(c byte)
	Reset()
}

type Buffer struct {
	list.List
	capacity int
}

var fullBufferError = errors.New("Buffer is full")
var emptyBufferError = errors.New("Buffer is empty")

func NewBuffer(capacity int) *Buffer {
	return &Buffer{capacity: capacity}
}

func (b *Buffer) WriteByte(c byte) error {
	if b.Len() == b.capacity {
		return fullBufferError
	}

	b.PushFront(c)

	return nil
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.Len() == 0 {
		return 0, emptyBufferError
	}

	result := b.Back()
	b.Remove(result)

	return result.Value.(byte), nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.Len() == b.capacity {
		b.Remove(b.Back())
	}

	b.PushFront(c)
}

func (b *Buffer) Reset() {
	b.Init()
}
