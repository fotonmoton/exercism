package paasio

import (
	"io"
	"sync"
)

type counter struct {
	mutex sync.Mutex
	count int64
	nops  int
}

type writeCounter struct {
	counter counter
	writer  io.Writer
}

type readCounter struct {
	counter counter
	reader  io.Reader
}

type readWriteCounter struct {
	writeCounter
	readCounter
}

func (c *counter) Count(n int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.count += n
	c.nops++
}

func (c *counter) GetCount() (count int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.count, c.nops
}

func (c *writeCounter) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)

	c.counter.Count(int64(n))

	return n, err
}

func (c *writeCounter) WriteCount() (n int64, nops int) {
	return c.counter.GetCount()
}

func (c *readCounter) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)

	c.counter.Count(int64(n))

	return n, err
}

func (c *readCounter) ReadCount() (n int64, nops int) {
	return c.counter.GetCount()
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{reader: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter: readCounter{
			reader: rw,
		},
		writeCounter: writeCounter{
			writer: rw,
		},
	}
}
