package coordinates

import (
	"math"
)

type Vector struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func NewVector(horizontalPosition, verticalPosition, distance float64) Vector {
	roe := math.Mod(-verticalPosition+90, 360)
	theta := math.Mod(-horizontalPosition+90, 360)
	r := float64(distance)
	return Vector{
		X: int(r * math.Sin(toRadians(theta)) * math.Cos(toRadians(roe))),
		Y: int(r * math.Sin(toRadians(theta)) * math.Sin(toRadians(roe))),
		Z: int(r * math.Cos(toRadians(theta))),
	}
}
