package main

import (
	"fmt"

	"voop/link"
)

func main() {
	message := link.Ping("status")
	data := message[7:]
	status, _ := link.Parse(data)
	fmt.Println("from socket", message)
	fmt.Println(status.Peers)
	fmt.Println(status.Bpm)
	fmt.Println(status.Start)
	fmt.Println(status.Beat)
}
