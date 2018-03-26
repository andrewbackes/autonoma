package sensor

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
)

// Sensor is all of the meta data for a physical sensor.
type Sensor struct {
	// ViewAngle in degrees.
	ViewAngle float64
	// MinDistance
	MinDistance distance.Distance
	// MaxDistance is the furthest the sensor can read in cm.
	MaxDistance distance.Distance
	// A binary sensor is on when an object is within range and
	// off otherwise.
	Binary bool
}
