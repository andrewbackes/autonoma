package coordinates

import (
	"math"
)

type Point struct {
	Origin      Vector `json:"origin"`
	Orientation Euler  `json:"orientation"`
	Vector      Vector `json:"vector"`
}

// ToVector combines origin, orientation and vector into a single vector.
func (p *Point) ToVector() Vector {
	// rotate vector by orientation and add it to origin.
	rad := toRadians(p.Orientation.Yaw)
	y := float64(p.Vector.Y)
	x := float64(p.Vector.X)
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	return Vector{
		X: p.Origin.X + round(x*cos+y*sin),
		Y: p.Origin.Y + round(-x*sin+y*cos),
		Z: p.Origin.Z + p.Vector.Z,
	}
}

func round(a float64) int {
	if a < 0 {
		return int(math.Ceil(a - 0.5))
	}
	return int(math.Floor(a + 0.5))
}
