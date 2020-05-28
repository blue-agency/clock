package clock

import "time"

type Clock interface {
	Now() time.Time
}

type clock struct {
}

func New() *clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return time.Now()
}
