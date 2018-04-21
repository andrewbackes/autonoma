// Package maps is for creating a 2d map of the world. It takes sensor data and fits it all together.
package maps

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/maps/grid"
)

// Map is a representation of the world in 2 dimensions.
type Map struct {
	grid   grid.Grid
	path   []coordinates.Vector
	course []coordinates.Vector
}

func (m *Map) MarkOccupied(v coordinates.Vector) {
	m.grid.Increase(v)
}

func (m *Map) MarkVacant(v coordinates.Vector) {
	m.grid.Decrease(v)
}

func (m *Map) SetCourse(c []coordinates.Vector) {
	m.course = c
}

func (m *Map) AppendPath(v coordinates.Vector) {
	m.path = append(m.path, v)
}

// Fit uses points to determine where the origin must be.
func (m *Map) Fit(p []coordinates.Point) coordinates.Vector {
	prediction := p[0].Origin
	bestFit := prediction
	bestScore := 0
	for x := prediction.X - 10; x <= prediction.X+10; x++ {
		for y := prediction.Y - 10; y <= prediction.Y+10; y++ {
			guess := coordinates.Vector{X: x, Y: y}
			score := 0
			for _, pp := range p {
				pp.Origin = guess
				v := pp.ToVector()

				if odds := m.grid.Get(v); odds != nil {
					if odds.Probability() > 0.99 {
						score++
					}
				}
			}
			if score > bestScore {
				bestScore = score
				bestFit = guess
			}
		}
	}
	return bestFit
}
