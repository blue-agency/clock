package clock

import (
	"errors"
	"time"
)

type Clock interface {
	Now() time.Time
}

type clock struct {
	location *time.Location
}

func New(loc *time.Location) (*clock, error) {
	if loc == nil {
		return nil, errors.New("'loc' must not be nil")
	}

	return &clock{location: loc}, nil
}

func (c *clock) Now() time.Time {
	return time.Now().In(c.location)
}
