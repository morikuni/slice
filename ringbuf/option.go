package ringbuf

// Option is optional parameter for New function.
// This was created for the future extension.
type Option func(conf *config)

type config struct{}
