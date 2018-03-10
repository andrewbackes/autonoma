package coordinates

import (
	"math"
)

type CompassRose struct {
	Distance float64
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
		X: int64(c.Distance * math.Cos(toRadians(angle))),
		Y: int64(c.Distance * math.Sin(toRadians(angle))),
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}
