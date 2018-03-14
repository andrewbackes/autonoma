package sensor

import (
	"encoding/json"
	"fmt"
	// log "github.com/sirupsen/logrus"
	"math"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
)

const (
	sensorAngleStepDegrees = 0.5
)

// Reading from a sensor.
type Reading struct {
	// TimeStamp when reading was taken.
	TimeStamp time.Time
	// Value returned by sensor.
	Value distance.Distance
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

func (r Reading) Analysis() (vacant, occupied coordinates.CartesianSet) {
	vacant = coordinates.NewCartesianSet()
	occupied = coordinates.NewCartesianSet()
	startAngle := r.Pose.Heading - (r.Sensor.ViewAngle / 2)
	endAngle := startAngle + r.Sensor.ViewAngle
	val := r.Value.Floor(distance.Centimeter)
	for d := r.Sensor.MinDistance.Floor(distance.Centimeter); d <= val; d++ {
		// reached max range
		if d >= r.Sensor.MaxDistance {
			return
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
			if d == val {
				// occupied
				occupied.Add(coord)
			} else if !occupied.Contains(coord) {
				// vacant
				vacant.Add(coord)
			}
		}
	}
	return
}
