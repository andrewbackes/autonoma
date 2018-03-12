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

var (
	// UltraSonic represents a physical UltraSonic sensor.
	UltraSonic = Sensor{
		ViewAngle:   15.0,
		MaxDistance: 18 * distance.Centimeter,
		MinDistance: 1 * distance.Centimeter,
		Binary:      false,
	}

	IRDistance = Sensor{
		ViewAngle:   0.0,
		MaxDistance: 80 * distance.Centimeter,
		MinDistance: 10 * distance.Centimeter,
		Binary:      false,
	}
)
