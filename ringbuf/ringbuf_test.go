package ringbuf

import (
	"testing"

	"github.com/morikuni/slice/internal/assert"
)

func TestBuffer(t *testing.T) {
	is := make([]int, 3)
	q := New(len(is))

	pushTail(is, q, 1)
	pushHead(is, q, 2)
	pushTail(is, q, 3)
	pushHead(is, q, 4)

	assert.Equal(t, 3, q.Len())

	assert.Equal(t, 2, peekHead(is, q))
	assert.Equal(t, 3, peekTail(is, q))
	assert.Equal(t, 2, popHead(is, q))
	assert.Equal(t, 3, popTail(is, q))
	assert.Equal(t, 1, popHead(is, q))
	assert.Equal(t, 0, popTail(is, q))

	pushTail(is, q, 5)
	pushTail(is, q, 6)

	assert.Equal(t, 2, q.Len())

	assert.Equal(t, 6, popTail(is, q))
	assert.Equal(t, 5, popTail(is, q))

	assert.Equal(t, 0, q.Len())
}

func pushHead(is []int, q *Buffer, val int) {
	idx, ok := q.PushHead()
	if ok {
		is[idx] = val
	}
}

func pushTail(is []int, q *Buffer, val int) {
	idx, ok := q.PushTail()
	if ok {
		is[idx] = val
	}
}

func popHead(is []int, q *Buffer) int {
	idx, ok := q.PopHead()
	if ok {
		return is[idx]
	}
	return 0
}

func popTail(is []int, q *Buffer) int {
	idx, ok := q.PopTail()
	if ok {
		return is[idx]
	}
	return 0
}

func peekHead(is []int, q *Buffer) int {
	idx, ok := q.PeekHead()
	if ok {
		return is[idx]
	}
	return 0
}

func peekTail(is []int, q *Buffer) int {
	idx, ok := q.PeekTail()
	if ok {
		return is[idx]
	}
	return 0
}
