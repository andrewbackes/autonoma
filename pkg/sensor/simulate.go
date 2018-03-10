package sensor

import (
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/set"
)

const (
	simulatorSensorStepAngle = 0.25
)

// SimulateReading creates a sensor reading based on
func SimulateReading(s Sensor, p Pose, occupied set.Set) Reading {
	r := Reading{
		Value:  s.MaxDistance,
		Sensor: s,
		Pose:   p,
	}
	startAngle := p.Heading - (s.ViewAngle / 2)
	endAngle := startAngle + s.ViewAngle
	for d := s.MinDistance; d < s.MaxDistance; d++ {
		for a := startAngle; a <= endAngle; a += simulatorSensorStepAngle {
			angle := math.Mod(a, 360)
			key := coordinates.CompassRose{
				Distance: d,
				Heading:  angle,
			}.Cartesian().String()
			if occupied.Contains(key) {
				r.Value = d
				return r
			}
		}
	}
	return r
}

// SimulatePoses creates poses at a certain distance and in a circle around
// each point.
func SimulatePoses(maxX, maxY, spacingCm int64, rotDeg float64) []Pose {
	ps := make([]Pose, 0)
	for x := -maxX + spacingCm; x < maxX; x += spacingCm {
		for y := -maxY + spacingCm; y < maxY; y += spacingCm {
			for h := float64(0.0); h <= 360.0; h += rotDeg {
				p := Pose{X: x, Y: y, Heading: h}
				ps = append(ps, p)
			}
		}
	}
	return ps
}
