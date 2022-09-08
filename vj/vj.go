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
	Set    *library.Set
	Config config.Keyboard
}

func (vj *VJ) OpenLibrary(folder *string) {

	lib, err := library.NewLibrary(folder, vj.Player.Transport)
	if err != nil {
		log.Fatal("cannot preload library", err)
	}
	vj.Lib = lib
}

func (vj *VJ) LoadSet(f func()) {
	f()
	path, ok := vj.Lib.What(vj.Lib.Now()).(*string)
	if !ok {
		log.Fatal("type conversion failed")
	}
	// preload set from folder
	set, err := library.NewSet(path, vj.Player.Transport)
	if err != nil {
		log.Fatal("cannot preload folder", err)
	}
	vj.Set = set
}

func (vj *VJ) ChooseMedia(f func()) {
	f()
	media, ok := vj.Set.What(vj.Set.Now()).(*clip.Media)
	if !ok {
		log.Fatal("type conversion failed")
	}
	vj.Player.Media = media
}

func (vj *VJ) OpenRndMediaParallel() {
	go func() {
		path, ok := vj.Lib.What(vj.Lib.Now()).(*string)
		if !ok {
			log.Fatal("type conversion failed")
		}
		file, _ := player.ChooseRandomFile(path)
		vj.Player.Media, _ = clip.NewMedia(file, vj.Player.Transport)
	}()
}

func (vj *VJ) WaitForAction() {
	for key := range vj.Player.HotKey {
		vj.Action(key)
	}
}
