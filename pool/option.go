package pool

import (
	"errors"
	"fmt"
	"time"
)

// Option is optional parameter for New function.
type Option func(conf *config)

// MinIdle configures the number of the idle objects in the pool.
// The idle objects are closed by Pool.CloseIdle function considering
// IdleTimeout option.
func MinIdle(n int) Option {
	return func(conf *config) {
		conf.min = n
	}
}

// IdleTimeout configures the duration which is used to detect
// idle objects no longer used.
func IdleTimeout(d time.Duration) Option {
	return func(conf *config) {
		conf.idleTimeout = d
	}
}

// for test.
func withNowFunc(now func() time.Time) Option {
	return func(conf *config) {
		conf.nowFunc = now
	}
}

type config struct {
	nowFunc     func() time.Time
	idleTimeout time.Duration
	min         int
}

func evaluateOptions(opts []Option) (*config, error) {
	conf := &config{
		nowFunc:     time.Now,
		idleTimeout: 0,
		min:         0,
	}

	for _, o := range opts {
		o(conf)
	}

	if conf.nowFunc == nil {
		return nil, errors.New("nowFunc is nil")
	}

	if conf.idleTimeout < 0 {
		return nil, fmt.Errorf("idle timeout must not be less than 0 but got %d", conf.idleTimeout)
	}

	if conf.min < 0 {
		return nil, fmt.Errorf("min must not be less than 0 but got %d", conf.min)
	}

	return conf, nil
}
