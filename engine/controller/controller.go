// Package controller is for sending instructions a bot.
package controller

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/gridmap"
)

// Controller is for operating a bot.
type Controller struct {
	mapReader gridmap.Reader
}

// New creates a Controller.
func New(r gridmap.Reader) *Controller {
	return &Controller{
		mapReader: r,
	}
}

// Start begins controlling the bot.
func (c *Controller) Start() {
	fmt.Println("Starting Controller.")
	fmt.Println("Stopped Controller.")
}
