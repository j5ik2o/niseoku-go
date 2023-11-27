package test

import "time"

type MockClock struct {
	now *time.Time
}

func NewMockClock(now *time.Time) *MockClock {
	return &MockClock{now: now}
}

func (c *MockClock) Now() *time.Time {
	return c.now
}
