package slice

type Queue struct {
	n     int
	start int
	l     int
}

func NewQueue(n int) *Queue {
	return &Queue{n: n}
}

func (q *Queue) PushHead() (int, bool) {
	if q.l == q.n {
		return 0, false
	}
	q.start = (q.start + q.n - 1) % q.n
	q.l += 1
	return q.start, true
}

func (q *Queue) PushTail() (int, bool) {
	if q.l == q.n {
		return 0, false
	}
	idx := (q.start + q.l) % q.n
	q.l += 1
	return idx, true
}

func (q *Queue) Len() int {
	return q.l
}

func (q *Queue) PopHead() (int, bool) {
	if q.l == 0 {
		return 0, false
	}
	idx := q.start
	q.start = (q.start + 1) % q.n
	q.l -= 1
	return idx, true
}

func (q *Queue) PopTail() (int, bool) {
	if q.l == 0 {
		return 0, false
	}
	idx := (q.start + q.l - 1) % q.n
	q.l -= 1
	return idx, true
}
