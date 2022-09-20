package vj

import (
	"log"
	"voop/clip"
	"voop/config"
	"voop/library"
	"voop/player"
)

type VJ struct {
	Player player.Player
	Lib    *library.Library
	Config config.Keyboard
}

func (vj *VJ) OpenLibrary(folder *string) {

	lib, err := library.NewLibrary(folder)
	if err != nil {
		log.Fatal("cannot preload library", err)
	}
	vj.Lib = lib
	m := vj.Lib.FileDefault()
	vj.Player.Media, _ = clip.NewMedia(m, vj.Player.Transport)
}

func (vj *VJ) WaitForAction() {
	for key := range vj.Player.HotKey {
		vj.Action(key)
	}
}
