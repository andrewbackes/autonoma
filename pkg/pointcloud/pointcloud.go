package pointcloud

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

// Point in space.
type Point struct {
	x []float64
}

func NewPoint(x ...float64) Point {
	p := Point{x: make([]float64, len(x))}
	for i, xv := range x {
		p.x[i] = xv
	}
	return p
}

func (p Point) ColMatrix() mat.Matrix {
	m := mat.NewDense(len(p.x), 1, nil)
	for i, x := range p.x {
		m.Set(i, 0, float64(x))
	}
	return m
}

func Subtract(a, b Point) Point {
	p := Point{x: make([]float64, len(a.x))}
	for i := range a.x {
		p.x[i] = a.x[i] - b.x[i]
	}
	return p
}

// PointCloud is a collection of points.
type PointCloud struct {
	points []Point
}

// Add a point to the cloud.
func (p *PointCloud) Add(pt Point) {
	p.points = append(p.points, pt)
}
func (p *PointCloud) Copy() *PointCloud {
	c := &PointCloud{points: make([]Point, len(p.points))}
	for i, v := range p.points {
		c.points[i] = v
	}
	return c
}

func (p *PointCloud) Centroid() Point {
	if len(p.points) == 0 {
		return Point{}
	}
	dimensions := p.Dimensions()
	centroid := Point{x: make([]float64, dimensions)}
	for dim := 0; dim < dimensions; dim++ {
		sum := float64(0)
		for _, pt := range p.points {
			sum += pt.x[dim]
		}
		centroid.x[dim] = sum / float64(len(p.points))
	}
	return centroid
}

// Subtract a point from every point in the point cloud. Returns a copy.
func (p *PointCloud) Subtract(pt Point) *PointCloud {
	c := p.Copy()
	for i := range c.points {
		c.points[i] = Subtract(c.points[i], pt)
	}
	return c
}

/*
// Distance is the shortest distance from a point to the point cloud.
func (p *PointCloud) Distance(to Point) float64 {
	if len(p.points) == 0 {
		return math.MaxFloat64
	}
	min := dist(p.points[0], to)
	for i := 1; i < len(p.points); i++ {
		d := dist(p.points[i], to)
		if d < min {
			min = d
		}
	}
	return min
}
*/

func dist(p, q Point) float64 {
	sum := float64(0)
	for i := 0; i < len(p.x); i++ {
		sum += math.Pow(float64(p.x[i]-q.x[i]), 2)
	}
	return math.Sqrt(sum)
}

func (p *PointCloud) Dimensions() int {
	if len(p.points) == 0 {
		return 0
	}
	return len(p.points[0].x)
}

func (p *PointCloud) Len() int {
	return len(p.points)
}

func (p *PointCloud) Matrix() mat.Matrix {
	if len(p.points) == 0 {
		return &mat.Dense{}
	}
	m := mat.NewDense(p.Dimensions(), p.Len(), nil)
	for col, pt := range p.points {
		for row, val := range pt.x {
			m.Set(row, col, float64(val))
		}
	}
	return m
}
