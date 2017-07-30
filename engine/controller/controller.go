// Package controller is for sending instructions a bot.
package controller

import (
	"github.com/andrewbackes/autonoma/engine/gridmap"
	"github.com/andrewbackes/autonoma/engine/sensor"
	"log"
	"net"
)

// Controller is for operating a bot.
type Controller struct {
	mapReader gridmap.Reader
	conn      net.Conn

	location sensor.Location
	heading  float64
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
	c.explore()
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
