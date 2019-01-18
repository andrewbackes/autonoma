// Package os is the operating system of the vehicle.
package os

import (
	"github.com/andrewbackes/autonoma/pkg/bot/comm"
	"github.com/andrewbackes/autonoma/pkg/bot/specs"
	"github.com/andrewbackes/autonoma/pkg/control"
	"github.com/andrewbackes/autonoma/pkg/perception"
	"github.com/andrewbackes/autonoma/pkg/perception/signal"
	"github.com/andrewbackes/autonoma/pkg/planning"
)

// OperatingSystem is the stack used in the operation of an autonomous robot.
type OperatingSystem struct {
	p *perception.Perception
	m *planning.Mission
	c *comm.Comm
	s specs.Spec
}

func (os *OperatingSystem) Perception() *perception.Perception {
	return os.p
}

func (os *OperatingSystem) signalHandler(s *signal.Signal) {
	os.p = signal.UpdatePerception(s, os.p)
	actions := planning.Plan(os.p, os.m)
	control.Execute(actions, os.c)
}

func New(c *comm.Comm, s specs.Spec) *OperatingSystem {
	return &OperatingSystem{
		p: perception.New(),
		m: planning.DefaultMission,
		c: c,
		s: s,
	}
}

func (os *OperatingSystem) Start() {
	os.c.Listen(os.signalHandler)
}
