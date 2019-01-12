package planning

import (
	"github.com/andrewbackes/autonoma/pkg/perception"
)

type Planner struct{}

func NewPlanner() *Planner {
	return &Planner{}
}

func (p *Planner) Plan(env *perception.Perception) *Motions {
	return &Motions{}
}

func (p *Planner) Mission() *Mission {
	return &Mission{}
}

func (p *Planner) SetMission(m *Mission) {

}
