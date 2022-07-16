package vj

import (
	"fmt"
	"log"
	"voop/config"
	"voop/library"
	"voop/player"
)

type VJ struct {
	Shortcuts  config.Keyboard
	Player     *player.Player
	Library    *library.Library
	UserAction chan string
}

func NewVJ(p *player.Player, l *library.Library, k config.Keyboard) *VJ {
	return &VJ{
		Shortcuts:  k,
		Player:     p,
		Library:    l,
		UserAction: make(chan string),
	}
}

func (v *VJ) Gig(f chan int) {
	var action config.Action
	// listen:
	for k := range f {
		log.Println("Gig")
		switch k {
		case -1:
			continue
		case 27:
			// stop clock
			log.Println("sending stop")
			v.Player.Transport.Stop()
			log.Println("vj stopped")
			return
		}
		action = v.Shortcuts[k]
		fmt.Println(action)
	}
}
