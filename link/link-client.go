package link

import (
	"bufio"
	"fmt"
	"net"
)

func Ping(message string) string {

	// connect to socket
	conn, _ := net.Dial("tcp", "127.0.0.1:17000")

	// send to socket
	fmt.Fprintf(conn, "%s", message)
	// listen response
	response, _ := bufio.NewReader(conn).ReadString('\n')
	// fmt.Print(message)
	return response
}
