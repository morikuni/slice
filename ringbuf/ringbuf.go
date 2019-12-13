package ringbuf

import (
	"fmt"
)

// Buffer represents ring buffer used with any types of slices.
// It calculates an index of slice for each operation.
type Buffer struct {
	n     int
	start int
	l     int
}

// New creates a new Buffer from size and options.
// The size configures the maximum number of the elements.
// Basically, the size is the length of the slice managed by buffer.
func New(size int, opts ...Option) (*Buffer, error) {
	if size < 0 {
		return nil, fmt.Errorf("size must be greater than 0 but got %d", size)
	}
	return &Buffer{n: size}, nil
}

// PushHead returns the head index of the slice where element must be put.
// If the second value was false, it means there is no room
// for the element.
func (q *Buffer) PushHead() (int, bool) {
	if q.l == q.n {
		return 0, false
	}

	q.start = (q.start + q.n - 1) % q.n
	q.l++

	return q.start, true
}

// PushTail returns the tail index of the slice where element must be put.
// If the second value was false, it means there is no room
// for the element.
func (q *Buffer) PushTail() (int, bool) {
	if q.l == q.n {
		return 0, false
	}

	idx := (q.start + q.l) % q.n
	q.l++

	return idx, true
}

// Len returns the length of the buffer.
func (q *Buffer) Len() int {
	return q.l
}

// PopHead returns an index of the head element.
// If the second value was false, it means there is
// no idle element in the pool.
func (q *Buffer) PopHead() (int, bool) {
	idx, ok := q.PeekHead()
	if !ok {
		return 0, false
	}

	q.start = (q.start + 1) % q.n
	q.l--

	return idx, true
}

// PopTail returns an index of the tail element.
// If the second value was false, it means there is
// no idle element in the pool.
func (q *Buffer) PopTail() (int, bool) {
	idx, ok := q.PeekTail()
	if !ok {
		return 0, false
	}

	q.l--

	return idx, true
}

// PopHead is the same as PopHead but it does not modify the buffer.
func (q *Buffer) PeekHead() (int, bool) {
	if q.l == 0 {
		return 0, false
	}

	return q.start, true
}

// PopHead is the same as PopTail but it does not modify the buffer.
func (q *Buffer) PeekTail() (int, bool) {
	if q.l == 0 {
		return 0, false
	}

	idx := (q.start + q.l - 1) % q.n

	return idx, true
}
