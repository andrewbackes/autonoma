package point

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

const Dimensions = 3

func (v Point) Matrix() mat.Matrix {
	m := mat.NewDense(3, 1, nil)
	m.Set(0, 0, float64(v.X))
	m.Set(1, 0, float64(v.Y))
	m.Set(2, 0, float64(v.Z))
	return m
}

// Add two Points together.
func Add(v1, v2 Point) Point {
	return Point{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

// Add two Points together.
func Subtract(v1, minusV2 Point) Point {
	return Point{
		X: v1.X - minusV2.X,
		Y: v1.Y - minusV2.Y,
		Z: v1.Z - minusV2.Z,
	}
}

func Equal(v1, v2 Point) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

func Distance(v, w Point) float64 {
	sum := float64(0)
	for i := 0; i < Dimensions; i++ {
		sum += math.Pow(v.Index(i)-w.Index(i), 2)
	}
	return math.Sqrt(sum)
}

func (v Point) Array() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

func (v *Point) SetIndex(index int, value float64) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	}
}

func (v Point) Index(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	}
	panic("index out of range")
}

// Rotate by angle in 2d space.
func Rotate(v Point, compassAngle float64) Point {
	rad := toRadians(compassAngle)
	return Point{
		X: v.X*math.Cos(rad) + v.Y*math.Sin(rad),
		Y: v.Y*math.Cos(rad) - v.X*math.Sin(rad),
		Z: v.Z,
	}
}

// PolarLikeCoordToPoint takes a heading from a compass and distance travelled,
// then returns a Point.
func PolarLikeCoordToPoint(compassAngle float64, distance float64) Point {
	// Compass rose coordinates go clockwise with north being on the y axis.
	// Polar coordinates start on the x axis and go counter clockwise.
	// To compensate take:
	//		compassdir = -polardir + 90
	//		polardir = -compassdir + 90
	angle := math.Mod(-compassAngle+90, 360)
	return Point{
		X: distance * math.Cos(toRadians(angle)),
		Y: distance * math.Sin(toRadians(angle)),
		Z: 0,
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func FromMatrix(m mat.Matrix) Point {
	r, _ := m.Dims()
	if r != Dimensions {
		panic("matrix is the wrong dimension")
	}
	v := Point{}
	for i := 0; i < Dimensions; i++ {
		v.SetIndex(i, m.At(i, 0))
	}
	return v
}
