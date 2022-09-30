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
		vj.SwitchMedia(vj.Lib.FileSuperRnd())
	// Random file from folder
	case "/":
		vj.SwitchMedia(vj.Lib.FileRandom())
	// next
	case ".":
		vj.SwitchMedia(vj.Lib.FileNext())
	// previous
	case ",":
		vj.SwitchMedia(vj.Lib.FilePrev())
	// random folder
	case "§":
		vj.SwitchMedia(vj.Lib.FldRnd())
	// next
	case "]":
		vj.SwitchMedia(vj.Lib.FldNext())
	// previous
	case "[":
		vj.SwitchMedia(vj.Lib.FldPrev())
	// RATE
	case "0":
		vj.Player.Media.DefaultRate()
	case "-":
		vj.Player.Media.RateX <- 2
	case "=":
		vj.Player.Media.RateX <- 0.5
	case "_":
		vj.Player.Media.RateX <- 1.5
	case "+":
		vj.Player.Media.RateX <- 0.75
	// Sync default to link clock
	case "r":
		vj.Player.Media.ReSync()
	// HardSync Mode
	case "H":
		vj.Player.Media.HardSyncToggle()
	// Zero (play from frame 0)
	case "z":
		vj.Player.Media.Zero()
	// Direction
	case "o":
		vj.Player.Media.Swap()
	// Jump!
	case "j":
		vj.Player.Media.Jump()
	// palindrome
	case "p":
		vj.Player.Media.PalindromemordnilaP()
	// Fullscreen
	case "f":
		vj.Player.Window.Fullscreen()
	}
	// debug info
	fmt.Println()
	fmt.Println(key)
	fmt.Println(ascii)
	fmt.Println(vj.Config[ascii])

}
