package domain

import "time"

type Clock interface {
	Now() *time.Time
}

type SystemClock struct{}

func (c *SystemClock) Now() *time.Time {
	now := time.Now()
	return &now
}
