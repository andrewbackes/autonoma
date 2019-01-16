// Package os is the operating system of the vehicle.
package os

import (
	"github.com/andrewbackes/autonoma/pkg/bot/v3"
	"github.com/andrewbackes/autonoma/pkg/control"
	"github.com/andrewbackes/autonoma/pkg/perception"
	"github.com/andrewbackes/autonoma/pkg/perception/signal"
	"github.com/andrewbackes/autonoma/pkg/planning"
)

// OperatingSystem is the stack used in the operation of an autonomous robot.
type OperatingSystem struct {
	p   *perception.Perception
	m   *planning.Mission
	bot *v3.Bot
}

func (os *OperatingSystem) Perception() *perception.Perception {
	return os.p
}

func (os *OperatingSystem) signalHandler(s *signal.Signal) {
	os.p = signal.UpdatePerception(s, os.p)
	actions := planning.Plan(os.p, os.m)
	control.Execute(actions, os.bot)
}

func New(bot *v3.Bot) *OperatingSystem {
	return &OperatingSystem{
		p:   perception.New(),
		m:   planning.DefaultMission,
		bot: bot,
	}
}

func (os *OperatingSystem) Start() {
	os.bot.Listen(os.signalHandler)
}
