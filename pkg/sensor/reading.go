package sensor

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
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
	Pose coordinates.Pose
}

func (r Reading) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("Distance %f", r.Value)
	}
	return string(b)
}

func (r Reading) Distance() *float64 {
	if r.Value >= r.Sensor.MaxDistance {
		return nil
	}
	return &r.Value
}

func (r Reading) Analysis() (vacant, occupied coordinates.CartesianSet) {
	if r.Value < r.Sensor.MaxDistance {
		log.Debug(r)
	}
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
				X: coord.X + r.Pose.Location.X,
				Y: coord.Y + r.Pose.Location.Y,
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
