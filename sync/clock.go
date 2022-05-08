package sync

import "time"

type Clock struct {
	Idle    time.Duration
	Trigger chan struct{}
}

func NewClock(t time.Duration) *Clock {
	c := Clock{Idle: t, Trigger: make(chan struct{})}
	go func() {
		for {
			c.Trigger <- struct{}{}
			time.Sleep(c.Idle)
		}
	}()
	return &c
}
