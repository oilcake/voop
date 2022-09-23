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
	// RATE
	case "0":
		vj.Player.Media.UpdateRate(1, vj.Player.Transport)
	case "-":
		vj.Player.Media.UpdateRate(2, vj.Player.Transport)
	case "=":
		vj.Player.Media.UpdateRate(0.5, vj.Player.Transport)
	case "_":
		vj.Player.Media.UpdateRate(1.5, vj.Player.Transport)
	case "+":
		vj.Player.Media.UpdateRate(0.75, vj.Player.Transport)
	// Sync default to link clock
	case "r":
		vj.Player.Media.ReSync()
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
		vj.Player.Media.PalindromemordnilaP(vj.Player.Transport)
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
