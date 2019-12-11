package pool

import (
	"fmt"
	"time"

	"github.com/morikuni/slice/queue"
)

// Pool is used with any types of slices.
// It calculates an index of slice for each operation.
type Pool struct {
	queue    *queue.Queue
	timeouts []time.Time
	conf     *config
}

// New creates a new Pool from size and options.
// The size configures the maximum number of the idle elements.
// Basically, the size is the length of the slice managed by pool.
func New(size int, opts ...Option) *Pool {
	return &Pool{
		queue:    queue.New(size),
		timeouts: make([]time.Time, size),
		conf:     evaluateOptions(opts),
	}
}

// Get returns an index of the latest idle element.
// If the second value was false, it means there is
// no idle element in the pool.
func (p *Pool) Get() (int, bool) {
	return p.queue.PopTail()
}

// Put returns an index of the slice where idle element must be put.
// If the second value was false, it means there is no room
// for the idle element.
func (p *Pool) Put() (int, bool) {
	idx, ok := p.queue.PushTail()
	if !ok {
		return 0, false
	}

	if p.conf.idleTimeout != 0 {
		p.timeouts[idx] = p.timeout()
	}

	return idx, true
}

var longEnough = 24 * time.Hour

// CloseIdle returns an index of the oldest idle element that should be closed.
// If the second value was false, it means there is no idle element should be closed.
// If there was no IdleTimeout option, it always returns false.
// The third value is the time when this function may return true again.
func (p *Pool) CloseIdle() (idx int, shouldClose bool, next time.Time) {
	if p.conf.idleTimeout == 0 {
		return 0, false, p.conf.nowFunc().Add(longEnough)
	}

	idx, ok := p.queue.PeekHead()
	if !ok {
		return 0, false, p.timeout()
	}

	l := p.queue.Len()
	if l == 0 {
		return 0, false, p.timeout()
	}

	if l <= p.conf.min {
		return 0, false, p.timeouts[idx]
	}

	if p.timeouts[idx].After(p.conf.nowFunc()) {
		return 0, false, p.timeouts[idx]
	}

	idx2, ok2 := p.queue.PopHead()
	if idx != idx2 || ok != ok2 {
		panic(fmt.Errorf("race condition detected. please use *slice.Pool with mutex: dx1=%v idx2=%v ok1=%v ok2=%v", idx, idx2, ok, ok2))
	}

	return idx, true, p.timeouts[idx]
}

func (p *Pool) timeout() time.Time {
	return p.conf.nowFunc().Add(p.conf.idleTimeout)
}
