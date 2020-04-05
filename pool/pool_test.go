package pool

import (
	"testing"
	"time"

	"github.com/morikuni/slice/internal/assert"
)

var (
	now  = time.Now()
	base = now
)

func nowFunc() time.Time {
	return now
}

func setNow(t time.Time) {
	now = t
}

func TestPool(t *testing.T) {
	const timeout = time.Second

	pool, err := New(5,
		MinIdle(2),
		IdleTimeout(timeout),
		withNowFunc(nowFunc),
	)
	assert.Equal(t, nil, err)

	checkGet(t, pool, 0, false)
	checkCloseIdle(t, pool, 0, false, base.Add(timeout))
	checkPut(t, pool, 0, true)

	// no close because min = 2
	checkCloseIdle(t, pool, 0, false, base.Add(timeout))
	checkPut(t, pool, 1, true)
	checkCloseIdle(t, pool, 0, false, base.Add(timeout))

	// no close because not timed out
	checkPut(t, pool, 2, true)
	checkCloseIdle(t, pool, 0, false, base.Add(timeout))

	// close because idle timeout.
	setNow(base.Add(2 * timeout))
	checkCloseIdle(t, pool, 0, true, base.Add(timeout))

	// no close because min = 2
	checkCloseIdle(t, pool, 0, false, base.Add(3*timeout)) // setNow(2) + idleTimeout(1)

	checkPut(t, pool, 3, true)
	checkPut(t, pool, 4, true)
	checkPut(t, pool, 0, true)
	checkPut(t, pool, 0, false)

	setNow(base.Add(4 * timeout))
	checkCloseIdle(t, pool, 1, true, base.Add(timeout))
	checkCloseIdle(t, pool, 2, true, base.Add(timeout))
	checkCloseIdle(t, pool, 3, true, base.Add(3*timeout)) // setNow(2) + idleTimeout(1)

	checkGet(t, pool, 0, true)
	checkGet(t, pool, 4, true)
	checkGet(t, pool, 0, false)
}

func checkGet(t *testing.T, pool *Pool, wantIdx int, wantOK bool) {
	t.Helper()

	idx, ok := pool.Get()
	assert.Equal(t, wantIdx, idx)
	assert.Equal(t, wantOK, ok)
}

func checkPut(t *testing.T, pool *Pool, wantIdx int, wantOK bool) {
	t.Helper()

	idx, ok := pool.Put()
	assert.Equal(t, wantIdx, idx)
	assert.Equal(t, wantOK, ok)
}

func checkCloseIdle(t *testing.T, pool *Pool, wantIdx int, wantOK bool, wantTimeout time.Time) {
	t.Helper()

	idx, ok, next := pool.CloseIdle()
	assert.Equal(t, wantIdx, idx)
	assert.Equal(t, wantOK, ok)
	assert.Equal(t, wantTimeout, next)
}
