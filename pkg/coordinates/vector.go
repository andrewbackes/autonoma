package coordinates

import (
	"math"
)

type Vector struct {
	X, Y, Z int
}

func NewVector(horizontalPosition, verticalPosition, distance float64) Vector {
	theta := math.Mod(-verticalPosition+90, 360)
	roe := math.Mod(-horizontalPosition+90, 360)
	r := float64(distance)
	return Vector{
		X: int(r * math.Sin(toRadians(theta)) * math.Cos(toRadians(roe))),
		Y: int(r * math.Sin(toRadians(theta)) * math.Sin(toRadians(roe))),
		Z: int(r * math.Cos(toRadians(theta))),
	}
}
