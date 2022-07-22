package clock

import (
	"log"
	"time"
)

const (
	fps = 26
	gap = 1 * time.Second / fps // time between two frames
)

type VClock struct {
	Idle    time.Duration
	Trigger chan struct{}
	stop    chan struct{}
}

func NewVClock() *VClock {
	vc := &VClock{
		Idle:    gap,
		Trigger: make(chan struct{}),
		stop:    make(chan struct{}),
	}
	return vc
}

func (vc *VClock) Start() {
	go func() {
		for {
			select {
			case <-vc.stop:
				log.Println("stop clock")
				close(vc.Trigger)
				return
			default:
				// pass
			}
			vc.Trigger <- struct{}{}
			time.Sleep(vc.Idle)
		}
	}()

}

func (vc *VClock) Stop() {
	vc.stop <- struct{}{}
}
