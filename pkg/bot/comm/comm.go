// Package comm is used to communicate with the bot.
package comm

import (
	log "github.com/sirupsen/logrus"
	"net"
)

type Comm struct {
	conn    net.Conn
	address string
}

func New(address string) *Comm {
	return &Comm{address: address}
}

func (c *Comm) connect() {
	log.Info("Connecting to ", c.address)
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		panic(err)
	}
	c.conn = conn
}
