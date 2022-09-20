package vj

import (
	"fmt"
)

var key string

func (vj *VJ) Action(ascii int) {
	key = string(rune(ascii))
	switch key {
	// Random file from library
	case "`":
		vj.Player.SwitchMedia(vj.Lib.FileSuperRnd())
	// Random file from folder
	case "/":
		vj.Player.SwitchMedia(vj.Lib.FileRandom())
	// next
	case ".":
		vj.Player.SwitchMedia(vj.Lib.FileNext())
	// previous
	case ",":
		vj.Player.SwitchMedia(vj.Lib.FilePrev())
	// random folder
	case "§":
		vj.Player.SwitchMedia(vj.Lib.FldRnd())
	// next
	case "]":
		vj.Player.SwitchMedia(vj.Lib.FldNext())
	// previous
	case "[":
		vj.Player.SwitchMedia(vj.Lib.FldPrev())
	}
	// debug info
	fmt.Println()
	fmt.Println(key)
	fmt.Println(ascii)
	fmt.Println(vj.Config[ascii])

}
