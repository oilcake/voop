package sync

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"olympos.io/encoding/edn"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"
)

type Carabiner struct {
	tcp       *net.Conn
	message   *string
	st        *Status
	tempoRary float32 // temporary placeholder for tempo changes checking
}

// Carabiner and its methods
func NewCarabiner() *Carabiner {
	return &Carabiner{
		message:   new(string),
		tcp:       NewConnection(),
		st:        new(Status),
		tempoRary: 0.0,
	}
}

// implementation of Linker interface
func (c *Carabiner) Link(status chan Status, t TempoWatcher) {
	c.sync()
	c.checkBpm(t)
	status <- *c.st
}

func (c *Carabiner) sync() {
	c.getStatus()
}

// check if Tempo has changed
func (c *Carabiner) checkBpm(t TempoWatcher) {
	if c.tempoRary != c.st.Bpm {
		t <- struct{}{}
		c.tempoRary = c.st.Bpm
	}
}

// Get info from Carabiner
func (c *Carabiner) getStatus() {
	fmt.Fprintf(*c.tcp, "%s", "status")                       // send message
	response, err := bufio.NewReader(*c.tcp).ReadString('\n') // listen response
	if err != nil {
		log.Fatal("connection lost", err)
	}
	response = response[7:]
	c.message = &response
	c.parse()
}

func (c *Carabiner) parse() {
	err := edn.Unmarshal([]byte(*c.message), c.st)
	if err != nil {
		log.Fatal("cannot parse link info", err)
	}
}
