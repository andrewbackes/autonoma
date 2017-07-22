// Package engine is for controlling bots.
package engine

import (
	"github.com/andrewbackes/autonoma/engine/controller"
	"github.com/andrewbackes/autonoma/engine/hud"
	"github.com/andrewbackes/autonoma/engine/occgrid"
	"github.com/andrewbackes/autonoma/engine/receiver"
)

// Engine recieves and processes sensor data in order to control a bot.
type Engine struct {
	hud        *hud.Hud
	receiver   *receiver.Receiver
	controller *controller.Controller
	gridmap    *occgrid.Grid
}

// NewEngine returns an engine with default parameters.
func NewEngine() *Engine {
	return &Engine{
		gridmap:    occgrid.NewDefaultGrid(1000, 1000),
		hud:        hud.New(),
		controller: controller.New(),
		receiver:   receiver.New(),
	}
}

func (e *Engine) Start() {
	if e.hud != nil {
		e.hud.Start()
	}
}
