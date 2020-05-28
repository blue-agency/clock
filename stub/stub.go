package stub

import "time"

type stub struct {
	fixedTime time.Time
}

func New(fixedTime time.Time) *stub {
	return &stub{
		fixedTime: fixedTime,
	}
}

func (s *stub) Now() time.Time {
	return s.fixedTime
}
