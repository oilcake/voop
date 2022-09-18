package vj

import "fmt"

var key string

func (vj *VJ) Action(ascii int) {
	key = string(rune(ascii))
	switch key {
	case "/":
		vj.ChooseMedia(vj.Set.Read.Random)
	case "]":
		go func() {
			vj.Set.Close()
			vj.LoadSet(vj.Lib.Read.Next)
			vj.ChooseMedia(vj.Set.Read.Default)
		}()
	case "`":
		vj.Lib.Read.Random()
		vj.OpenRndMediaParallel()
	}
	fmt.Println()
	fmt.Println(key)
	fmt.Println(vj.Config[ascii])

}
