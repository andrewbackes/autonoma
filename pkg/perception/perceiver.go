package perception

import (
	"github.com/andrewbackes/autonoma/pkg/sensing"
)

type Perceiver struct{}

func NewPerceiver() *Perceiver {
	return &Perceiver{}
}

func (p *Perceiver) Perceive(*sensing.SensorData) *Perception {
	return &Perception{}
}

func (p *Perceiver) Perception() *Perception {
	return &Perception{}
}
