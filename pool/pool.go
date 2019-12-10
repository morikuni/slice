package pool

import (
	"fmt"
	"time"

	"github.com/morikuni/slice/queue"
)

type Pool struct {
	queue    *queue.Queue
	timeouts []time.Time
	conf     *config
}

func New(max int, opts ...Option) *Pool {
	return &Pool{
		queue:    queue.New(max),
		timeouts: make([]time.Time, max),
		conf:     evaluateOptions(opts),
	}
}

func (p *Pool) Get() (int, bool) {
	return p.queue.PopTail()
}

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

var longEnough = time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC)

func (p *Pool) CloseIdle() (idx int, shouldClean bool, next time.Time) {
	if p.conf.idleTimeout == 0 {
		return 0, false, longEnough
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
