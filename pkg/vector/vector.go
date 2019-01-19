package vector

import (
	"math"
)

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Add two vectors together.
func Add(v1, v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

// PolarLikeCoordToVector takes a heading from a compass and distance travelled,
// then returns a vector.
func PolarLikeCoordToVector(compassAngle float64, distance float64) Vector {
	// Compass rose coordinates go clockwise with north being on the y axis.
	// Polar coordinates start on the x axis and go counter clockwise.
	// To compensate take:
	//		compassdir = -polardir + 90
	//		polardir = -compassdir + 90
	angle := math.Mod(-compassAngle+90, 360)
	return Vector{
		X: float64(distance) * math.Cos(toRadians(angle)),
		Y: float64(distance) * math.Sin(toRadians(angle)),
		Z: 0,
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}
