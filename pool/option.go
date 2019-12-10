package pool

import (
	"time"
)

type Option func(conf *config)

func MinIdle(n int) Option {
	return func(conf *config) {
		conf.min = n
	}
}

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

func evaluateOptions(opts []Option) *config {
	conf := &config{
		nowFunc:     time.Now,
		idleTimeout: 0,
		min:         0,
	}

	for _, o := range opts {
		o(conf)
	}

	return conf
}
