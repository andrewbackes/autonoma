package coordinates

import (
	"math"

	"github.com/andrewbackes/autonoma/pkg/distance"
)

type CompassRose struct {
	Distance distance.Distance
	Heading  float64
}

func (c CompassRose) Cartesian() Cartesian {
	// Compass rose coordinates go clockwise with north being on the y axis.
	// Polar coordinates start on the x axis and go counter clockwise.
	// To compensate take:
	//		compassdir = -polardir + 90
	//		polardir = -compassdir + 90
	angle := math.Mod(-c.Heading+90, 360)
	return Cartesian{
		X: int(float64(c.Distance) * math.Cos(toRadians(angle))),
		Y: int(float64(c.Distance) * math.Sin(toRadians(angle))),
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}
