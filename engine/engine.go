// Package engine is for controlling bots.
package engine

import (
	"github.com/andrewbackes/autonoma/engine/controller"
	"github.com/andrewbackes/autonoma/engine/hud"
	"github.com/andrewbackes/autonoma/engine/occgrid"
	"github.com/andrewbackes/autonoma/engine/receiver"
	"log"
	"net"
	"sync"
	"time"
)

// Engine recieves and processes sensor data in order to control a bot.
type Engine struct {
	hud        *hud.Hud
	receiver   *receiver.Receiver
	controller *controller.Controller

	botAddr string
	conn    net.Conn
	wg      *sync.WaitGroup
}

// NewEngine returns an engine with default parameters.
func NewEngine() *Engine {
	grid := occgrid.NewGrid(1000, 1000, 10)
	return &Engine{
		hud:        hud.New(grid),
		controller: controller.New(grid),
		receiver:   receiver.New(grid),
		botAddr:    "localhost:9091",
		wg:         &sync.WaitGroup{},
	}
}

// Start turns on the engine.
func (e *Engine) Start() {
	e.connect()
	e.wg.Add(1)
	go func() {
		e.hud.Serve()
		e.wg.Done()
	}()
	e.wg.Add(1)
	go func() {
		e.receiver.Listen(e.conn)
		e.wg.Done()
	}()
	e.wg.Add(1)
	go func() {
		e.controller.Start(e.conn)
		e.wg.Done()
	}()
	e.wg.Wait()
}

func (e *Engine) connect() {
	if e.conn != nil {
		e.conn.Close()
	}
	c, err := net.DialTimeout("tcp", e.botAddr, 15*time.Second)
	if err != nil {
		panic("Could't connect to bot")
	}
	e.conn = c
	log.Println("Connected to", c.RemoteAddr())
}
