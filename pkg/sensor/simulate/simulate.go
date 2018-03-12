package simulate

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"
	"math"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

const (
	simulatorSensorStepAngle = 0.5
)

// Reading creates a sensor reading based on
func Reading(s sensor.Sensor, p coordinates.Pose, occupied coordinates.CartesianSet) sensor.Reading {
	r := sensor.Reading{
		TimeStamp: time.Now(),
		Value:     s.MaxDistance,
		Sensor:    s,
		Pose:      p,
	}
	startAngle := p.Heading - (s.ViewAngle / 2)
	endAngle := startAngle + s.ViewAngle
	for d := distance.Distance(0); d < s.MaxDistance; d++ {
		for a := startAngle; a <= endAngle; a += simulatorSensorStepAngle {
			angle := math.Mod(a, 360)
			coord := coordinates.CompassRose{
				Distance: d,
				Heading:  angle,
			}.Cartesian()
			coord.X += p.Location.X
			coord.Y += p.Location.Y

			if occupied.Contains(coord) && d < s.MinDistance {
				return r
			}

			if occupied.Contains(coord) && d >= s.MinDistance {
				log.Debug("Reading found occupied square ", coord)
				r.Value = d
				return r
			}
		}
	}
	return r
}

// Poses creates poses at a certain distance and in a circle around
// each point.
func Poses(maxX, maxY, spacingCm int, rotDeg float64) []coordinates.Pose {
	ps := make([]coordinates.Pose, 0)
	for x := -maxX + spacingCm; x < maxX; x += spacingCm {
		for y := -maxY + spacingCm; y < maxY; y += spacingCm {
			for h := float64(0.0); h <= 360.0; h += rotDeg {
				p := coordinates.NewPose(x, y, h)
				ps = append(ps, p)
			}
		}
	}
	return ps
}
