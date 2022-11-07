package vj

import (
	"log"
	"voop/clip"
	"voop/config"
	"voop/library"
	"voop/player"
	"voop/sync"
)

var (
	err             error
	oldMedia, media *clip.Media
	mNext           chan *clip.Media
)

type VJ struct {
	Transport *sync.Transport
	Player    player.Player
	Lib       *library.Library
	Media     chan *clip.Media
	Shortcuts config.Keyboard
	Config    config.Config
}

func (vj *VJ) OpenLibrary(folder *string) {

	lib, err := library.NewLibrary(folder, vj.Config.Supported)
	if err != nil {
		log.Fatal("cannot preload library", err)
	}
	vj.Lib = lib
	m, err := clip.NewMedia(vj.Lib.FileDefault(), vj.Transport)
	if err != nil {
		log.Fatal("error while opening media", err)
	}
	vj.Player.Media = m
	vj.Player.Resizer.ResizeFrom(*m.Shape)
}

func (vj *VJ) WaitForAction() {
	for key := range vj.Player.HotKey {
		action := vj.Shortcuts[key]
		vj.Action(action)
	}
}

func (vj *VJ) SwitchMedia(path string) {
	var (
		err   error
		media *clip.Media
	)
	go func() {
		media, err = clip.NewMedia(path, vj.Transport)
		if err != nil {
			log.Fatal("error while opening media", err)
		}
		vj.Media <- media
	}()

}
