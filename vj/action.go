package vj

import (
	"fmt"
	"voop/config"
)

var key string

func (vj *VJ) Action(action config.Action) {
	switch action.Target {
	case "Clip":
		switch action.Command {

		// Random file from library
		case "super_random":
			vj.SwitchMedia(vj.Lib.FileSuperRnd())
		// Random file from folder
		case "random":
			vj.SwitchMedia(vj.Lib.FileRandom())
		// next
		case "next":
			vj.SwitchMedia(vj.Lib.FileNext())
		// previous
		case "previous":
			vj.SwitchMedia(vj.Lib.FilePrev())
		// RATE
		case "rate_default":
			vj.Player.Media.DefaultRate()
		case "rate_half":
			vj.Player.Media.RateX <- 2
		case "rate_double":
			vj.Player.Media.RateX <- 0.5
		case "rate_0.75":
			vj.Player.Media.RateX <- 1.5
		case "rate_1.5":
			vj.Player.Media.RateX <- 0.75
		// Sync default to link clock
		case "resync":
			vj.Player.Media.ReSync()
		// HardSync Mode
		case "hard_sync":
			vj.Player.Media.HardSyncToggle()
		// Zero (play from frame 0)
		case "zero":
			vj.Player.Media.Zero()
		// Direction
		case "dir_flip":
			vj.Player.Media.Swap()
		// Jump!
		case "jump_random":
			vj.Player.Media.Jump()
		// palindrome
		case "palindrome":
			vj.Player.Media.PalindromemordnilaP()
		}
	case "Kit":
		switch action.Command {

		// random folder
		case "random":
			vj.SwitchMedia(vj.Lib.FldRnd())
		// next
		case "next":
			vj.SwitchMedia(vj.Lib.FldNext())
		// previous
		case "previous":
			vj.SwitchMedia(vj.Lib.FldPrev())
		}
	}
	// debug info
	fmt.Println()
	fmt.Println(action.Target, action.Command)

}
