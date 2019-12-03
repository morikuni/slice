package slice

import (
	"testing"
)

func TestQueue(t *testing.T) {
	is := make([]int, 3)
	q := NewQueue(len(is))

	pushTail(is, q, 1)
	pushHead(is, q, 2)
	pushTail(is, q, 3)
	pushHead(is, q, 4)

	assertEqual(t, 3, q.Len())

	assertEqual(t, 2, popHead(is, q))
	assertEqual(t, 3, popTail(is, q))
	assertEqual(t, 1, popHead(is, q))
	assertEqual(t, 0, popTail(is, q))

	pushTail(is, q, 5)
	pushTail(is, q, 6)

	assertEqual(t, 2, q.Len())

	assertEqual(t, 6, popTail(is, q))
	assertEqual(t, 5, popTail(is, q))

	assertEqual(t, 0, q.Len())
}

func pushHead(is []int, q *Queue, val int) {
	idx, ok := q.PushHead()
	if ok {
		is[idx] = val
	}
}

func pushTail(is []int, q *Queue, val int) {
	idx, ok := q.PushTail()
	if ok {
		is[idx] = val
	}
}

func popHead(is []int, q *Queue) int {
	idx, ok := q.PopHead()
	if ok {
		return is[idx]
	}
	return 0
}

func popTail(is []int, q *Queue) int {
	idx, ok := q.PopTail()
	if ok {
		return is[idx]
	}
	return 0
}
