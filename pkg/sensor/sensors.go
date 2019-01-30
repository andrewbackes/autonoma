package sensor

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
)

var (
	// UltraSonicHCSR04 distance sensor.
	UltraSonicHCSR04 = Sensor{
		ViewAngle:   15.0,
		MaxDistance: 140 * distance.Centimeter,
		MinDistance: 2 * distance.Centimeter,
		Binary:      false,
	}

	// SharpGP2Y0A21YK0F IR Distance Sensor.
	SharpGP2Y0A21YK0F = Sensor{
		ViewAngle:   0.0,
		MaxDistance: 80 * distance.Centimeter,
		MinDistance: 10 * distance.Centimeter,
		Binary:      false,
	}

	// GarminLidarLiteV3 Sensor.
	GarminLidarLiteV3 = Sensor{
		ViewAngle:   0.0,
		MaxDistance: 40 * distance.Meter,
		MinDistance: 2 * distance.Centimeter,
		Binary:      false,
	}
)
