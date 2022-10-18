package sync

import (
	"log"

	"olympos.io/encoding/edn"
)

type Carabiner struct {
	St  *Status
	cnn connector
	err error
}

type connector interface {
	getStatus() string
}

// Carabiner and its methods
func NewCarabiner(cnn connector) *Carabiner {
	return &Carabiner{
		St:  new(Status),
		cnn: cnn,
		err: nil,
	}
}

func (c *Carabiner) grab() string { return c.cnn.getStatus() }

func (c *Carabiner) parse(message string) {
	c.err = edn.Unmarshal([]byte(message), c.St)
	if c.err != nil {
		log.Fatal("cannot parse link info", c.err)
	}
}

// implementation of Linker interface
func (c *Carabiner) Sync() {
	c.parse(c.grab())
}

func (c *Carabiner) Provide() Status {
	c.Sync()
	return *c.St
}
