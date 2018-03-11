package simulate

import (
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

const (
	simulatorSensorStepAngle = 0.5
)

// Reading creates a sensor reading based on
func Reading(s sensor.Sensor, p sensor.Pose, occupied coordinates.CartesianSet) sensor.Reading {
	r := sensor.Reading{
		Value:  s.MaxDistance,
		Sensor: s,
		Pose:   p,
	}
	startAngle := p.Heading - (s.ViewAngle / 2)
	endAngle := startAngle + s.ViewAngle
	for d := s.MinDistance; d < s.MaxDistance; d++ {
		for a := startAngle; a <= endAngle; a += simulatorSensorStepAngle {
			angle := math.Mod(a, 360)
			coord := coordinates.CompassRose{
				Distance: d,
				Heading:  angle,
			}.Cartesian()
			coord.X += p.X
			coord.Y += p.Y
			if occupied.Contains(coord) {
				r.Value = d
				return r
			}
		}
	}
	return r
}

// Poses creates poses at a certain distance and in a circle around
// each point.
func Poses(maxX, maxY, spacingCm int, rotDeg float64) []sensor.Pose {
	ps := make([]sensor.Pose, 0)
	for x := -maxX + spacingCm; x < maxX; x += spacingCm {
		for y := -maxY + spacingCm; y < maxY; y += spacingCm {
			for h := float64(0.0); h <= 360.0; h += rotDeg {
				p := sensor.Pose{X: x, Y: y, Heading: h}
				ps = append(ps, p)
			}
		}
	}
	return ps
}
