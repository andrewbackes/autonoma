// Package controller is for sending instructions a bot.
package controller

import (
	"github.com/andrewbackes/autonoma/engine/controller/actions"
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"log"
	"net"
	"time"
)

// Controller is for operating a bot.
type Controller struct {
	mapReader gridmap.Reader
	conn      net.Conn
}

// New creates a Controller.
func New(r gridmap.Reader) *Controller {
	return &Controller{
		mapReader: r,
	}
}

// Start begins controlling the bot.
func (c *Controller) Start(conn net.Conn) {
	log.Println("Starting Controller.")
	c.conn = conn
	log.Println(c.conn)
	c.think()
	log.Println("Stopped Controller.")
}

func (c *Controller) send(payload string) {
	log.Println("Sending", payload)
	if c.conn == nil {
		log.Println("Controller not connected.")
	} else {
		c.conn.Write([]byte(payload + "\n"))
	}
}

func (c *Controller) think() {
	time.Sleep(1 * time.Second)
	c.ScanArea()
}

func (c *Controller) ScanArea() {
	for i := float64(0); i <= 360; i += 15 {
		c.send(actions.Rotate(i))
		c.send(actions.Read("all"))
	}
}
