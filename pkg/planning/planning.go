// Package planning uses the model of the environment and the robot position to create a set of motions the robot should do in order to accomplish a mission.
package planning

import (
	"github.com/andrewbackes/autonoma/pkg/perception"
)

// Plan a set of actions to take given the view of the world and a mission.
func Plan(p *perception.Perception, m *Mission) *Motions {
	return &Motions{}
}
