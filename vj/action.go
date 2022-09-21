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
		vj.Player.Media.Multiple = 1.0
		vj.Player.Media.Grooverize(vj.Player.Transport)
	case "-":
		vj.Player.Media.Multiple = vj.Player.Media.Multiple * 2.0
		vj.Player.Media.Grooverize(vj.Player.Transport)
	case "=":
		vj.Player.Media.Multiple = vj.Player.Media.Multiple * 0.5
		vj.Player.Media.Grooverize(vj.Player.Transport)
	case "_":
		vj.Player.Media.Multiple = vj.Player.Media.Multiple * 1.5
		vj.Player.Media.Grooverize(vj.Player.Transport)
	case "+":
		vj.Player.Media.Multiple = vj.Player.Media.Multiple * 0.75
		vj.Player.Media.Grooverize(vj.Player.Transport)
	// Direction
	case "o":
		vj.Player.Media.Swap()
		fmt.Println()
		fmt.Println("Fucking Swap!")
	// palindrome
	case "p":
		vj.Player.Media.PalindromemordnilaP(vj.Player.Transport)
	}
	// debug info
	fmt.Println()
	fmt.Println(key)
	fmt.Println(ascii)
	fmt.Println(vj.Config[ascii])

}
