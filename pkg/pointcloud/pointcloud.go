package pointcloud

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// Point in space.
type Point struct {
	X []float64
}

func NewPoint(x ...float64) Point {
	p := Point{X: make([]float64, len(x))}
	for i, xv := range x {
		p.X[i] = xv
	}
	return p
}

func (p Point) ColMatrix() mat.Matrix {
	m := mat.NewDense(len(p.X), 1, nil)
	for i, x := range p.X {
		m.Set(i, 0, float64(x))
	}
	return m
}

func Subtract(a, b Point) Point {
	p := Point{X: make([]float64, len(a.X))}
	for i := range a.X {
		p.X[i] = a.X[i] - b.X[i]
	}
	return p
}

// PointCloud is a collection of Points.
type PointCloud struct {
	Points []Point
}

func New() *PointCloud {
	return &PointCloud{
		Points: make([]Point, 0),
	}
}

// Add a point to the cloud.
func (p *PointCloud) Add(pt Point) {
	p.Points = append(p.Points, pt)
}
func (p *PointCloud) Copy() *PointCloud {
	c := &PointCloud{Points: make([]Point, len(p.Points))}
	for i, v := range p.Points {
		c.Points[i] = v
	}
	return c
}

func (p *PointCloud) Centroid() Point {
	if len(p.Points) == 0 {
		return Point{}
	}
	dimensions := p.Dimensions()
	centroid := Point{X: make([]float64, dimensions)}
	for dim := 0; dim < dimensions; dim++ {
		sum := float64(0)
		for _, pt := range p.Points {
			sum += pt.X[dim]
		}
		centroid.X[dim] = sum / float64(len(p.Points))
	}
	return centroid
}

// Subtract a point from every point in the point cloud. Returns a copy.
func (p *PointCloud) Subtract(pt Point) *PointCloud {
	c := p.Copy()
	for i := range c.Points {
		c.Points[i] = Subtract(c.Points[i], pt)
	}
	return c
}

/*
// Distance is the shortest distance from a point to the point cloud.
func (p *PointCloud) Distance(to Point) float64 {
	if len(p.Points) == 0 {
		return math.MaxFloat64
	}
	min := dist(p.Points[0], to)
	for i := 1; i < len(p.Points); i++ {
		d := dist(p.Points[i], to)
		if d < min {
			min = d
		}
	}
	return min
}
*/

func dist(p, q Point) float64 {
	sum := float64(0)
	for i := 0; i < len(p.X); i++ {
		sum += math.Pow(float64(p.X[i]-q.X[i]), 2)
	}
	return math.Sqrt(sum)
}

func (p *PointCloud) Dimensions() int {
	if len(p.Points) == 0 {
		return 0
	}
	return len(p.Points[0].X)
}

func (p *PointCloud) Len() int {
	return len(p.Points)
}

func (p *PointCloud) Matrix() mat.Matrix {
	if len(p.Points) == 0 {
		return &mat.Dense{}
	}
	m := mat.NewDense(p.Dimensions(), p.Len(), nil)
	for col, pt := range p.Points {
		for row, val := range pt.X {
			m.Set(row, col, float64(val))
		}
	}
	return m
}

func printMatrix(m mat.Matrix) {
	r, c := m.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(m.At(i, j), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
