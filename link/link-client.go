package link

import (
	"bufio"
	"fmt"
	"net"
)

func Ping(message string) string {

	// open socket
	conn, _ := net.Dial("tcp", "127.0.0.1:17000")

	// send message
	fmt.Fprintf(conn, "%s", message)
	// listen response
	response, _ := bufio.NewReader(conn).ReadString('\n')
	return response
}
