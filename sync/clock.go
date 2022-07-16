package sync

import (
	"log"
	"time"
)

type Clock struct {
	Idle    time.Duration
	Trigger chan struct{}
	stop    chan struct{}
}

func NewClock(t time.Duration) *Clock {
	c := Clock{Idle: t, Trigger: make(chan struct{})}
	return &c
}

func (c *Clock) Start() {
	go func() {
		for {
			select {
			case <-c.stop:
				log.Println("stop clock")
				close(c.Trigger)
				return
			default:
				// pass
			}
			c.Trigger <- struct{}{}
			time.Sleep(c.Idle)
			log.Println("clock")
		}
	}()

}

func (c *Clock) Stop() {
	c.stop <- struct{}{}
}
