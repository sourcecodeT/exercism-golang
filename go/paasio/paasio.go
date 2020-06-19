package paasio

import (
	"io"
	"sync"
)

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w, counter: newCounter()}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r, counter: newCounter()}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}

//------------------------------------------------------------------
type writeCounter struct {
	w io.Writer
	counter
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	m, err := wc.w.Write(p)
	wc.addBytes(m)
	return m, err
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	return wc.Count()
}

//------------------------------------------------------------------

type readCounter struct {
	r io.Reader
	counter
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.r.Read(p)
	rc.addBytes(m)
	return m, err
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.Count()
}

//------------------------------------------------------------------

type rwCounter struct {
	WriteCounter
	ReadCounter
}

//------------------------------------------------------------------

func newCounter() counter {
	return counter{mux: new(sync.Mutex)}
}

type counter struct {
	bytes int64
	ops   int
	mux   *sync.Mutex
}

func (c *counter) addBytes(n int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.bytes = c.bytes + int64(n)
	c.ops = c.ops + 1
}

func (c *counter) Count() (int64, int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.bytes, c.ops
}
