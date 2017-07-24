// Package engine is for controlling bots.
package engine

import (
	"github.com/andrewbackes/autonoma/engine/controller"
	"github.com/andrewbackes/autonoma/engine/hud"
	"github.com/andrewbackes/autonoma/engine/occgrid"
	"github.com/andrewbackes/autonoma/engine/receiver"
	"sync"
)

// Engine recieves and processes sensor data in order to control a bot.
type Engine struct {
	hud        *hud.Hud
	receiver   *receiver.Receiver
	controller *controller.Controller

	wg *sync.WaitGroup
}

// Starter can begin a service.
type Starter interface {
	Start()
}

// NewEngine returns an engine with default parameters.
func NewEngine() *Engine {
	grid := occgrid.NewGrid(1000, 1000, 10, 5)
	return &Engine{
		hud:        hud.New(grid),
		controller: controller.New(grid),
		receiver:   receiver.New(grid),
		wg:         &sync.WaitGroup{},
	}
}

func (e *Engine) start(s Starter) {
	if s != nil {
		e.wg.Add(1)
		go func() {
			s.Start()
			e.wg.Done()
		}()
	}
}

// Start turns on the engine.
func (e *Engine) Start() {
	e.start(e.hud)
	e.start(e.controller)
	e.start(e.receiver)
	e.wg.Wait()
}
