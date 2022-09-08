package vj

import "fmt"

func (vj *VJ) Action(key int) {
	switch key {
	case getKey('/'):
		vj.ChooseMedia(vj.Set.Read.Random)
	case getKey(']'):
		go func() {
			vj.Set.Close()
			vj.LoadSet(vj.Lib.Read.Next)
			vj.ChooseMedia(vj.Set.Read.Default)
		}()
	case getKey('`'):
		vj.Lib.Read.Random()
		vj.OpenRndMediaParallel()
	}
	fmt.Println()
	fmt.Println(key)
	fmt.Println(vj.Config[key])

}

func getKey(r rune) int {
	return int(r)
}
