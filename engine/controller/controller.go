package controller

import (
	"fmt"
	"github.com/andrewbackes/autonoma/engine/gridmap"
)

type Controller struct {
	mapReader gridmap.Reader
}

func New(r gridmap.Reader) *Controller {
	return &Controller{
		mapReader: r,
	}
}

func (c *Controller) Start() {
	fmt.Println("Starting Controller.")
	fmt.Println("Stopped Controller.")
}
