package autonoma

import (
	"github.com/andrewbackes/autonoma/pkg/world"
)

type Autonoma struct {
	world *world.World
}

type mission interface {
	Start()
	Stop()
}

type Option func(*Autonoma)

func New(options ...Option) *Autonoma {
	a := &Autonoma{
		world: world.New(),
	}
	for _, option := range options {
		option(a)
	}
	return a
}

func (a *Autonoma) SetWorld(w *world.World) {
	a.world = w
}

func (a *Autonoma) World() *world.World {
	return a.world
}

func (a *Autonoma) NewWorld(w *world.World) {
	a.world = world.New()
}

func (a *Autonoma) Start() {}
