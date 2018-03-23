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
	// UltraSonicHCSR04 represents an UltraSonic HC-SR04 distance sensor.
	UltraSonicHCSR04 = Sensor{
		ViewAngle:   15.0,
		MaxDistance: 140 * distance.Centimeter,
		MinDistance: 2 * distance.Centimeter,
		Binary:      false,
	}

	// SharpGP2Y0A21YK0F represents a Sharp IR Distance Sensor GP2Y0A21YK0F
	SharpGP2Y0A21YK0F = Sensor{
		ViewAngle:   0.0,
		MaxDistance: 80 * distance.Centimeter,
		MinDistance: 10 * distance.Centimeter,
		Binary:      false,
	}
)
