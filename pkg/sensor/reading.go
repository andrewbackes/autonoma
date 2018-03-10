package sensor

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"math"
	"time"
)

const (
	sensorAngleStepDegrees = 0.5
)

// Reading from a sensor.
type Reading struct {
	// TimeStamp when reading was taken.
	TimeStamp time.Time
	// Value returned by sensor.
	Value float64
	// Sensor that generated the reading.
	Sensor Sensor
	// Pose is the orientation of the sensor.
	Pose Pose
}

func (r Reading) Distance() *float64 {
	if r.Value >= r.Sensor.MaxDistance {
		return nil
	}
	return &r.Value
}

func (r Reading) Analysis() (vacant, occupied coordinates.CartesianSet) {
	vacant = coordinates.NewCartesianSet()
	occupied = coordinates.NewCartesianSet()
	startAngle := r.Pose.Heading - (r.Sensor.ViewAngle / 2)
	endAngle := startAngle + r.Sensor.ViewAngle
	for d := r.Sensor.MinDistance; ; d++ {
		// reached max range
		if d >= r.Sensor.MaxDistance {
			break
		}
		// Don't xray through the obsticle.
		if d > math.Floor(r.Value) {
			break
		}
		for a := startAngle; a <= endAngle; a += sensorAngleStepDegrees {
			angle := math.Mod(a, 360)
			coord := coordinates.CompassRose{
				Distance: d,
				Heading:  angle,
			}.Cartesian()
			// adjust for position
			coord = coordinates.Cartesian{
				X: coord.X + r.Pose.X,
				Y: coord.Y + r.Pose.Y,
			}
			if d == math.Floor(r.Value) {
				// occupied
				occupied.Add(coord)
			} else {
				// vacant
				vacant.Add(coord)
			}
		}
	}
	return
}
