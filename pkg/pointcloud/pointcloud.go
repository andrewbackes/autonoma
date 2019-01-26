package pointcloud

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/vector"
	"gonum.org/v1/gonum/mat"
)

// PointCloud is a collection of Points.
type PointCloud struct {
	Points map[vector.Vector]int
}

const Dimensions = 3

func New() *PointCloud {
	return &PointCloud{
		Points: map[vector.Vector]int,
	}
}

// Add a point to the cloud.
func (p *PointCloud) Add(v vector.Vector) {
	val := p.Points[v]
	p.Points[v] = val + 1
}

func Copy(p *PointCloud) *PointCloud {
	n := New()
	for k, v := range p.Points {
		n[k] = v
	}
}

func (p *PointCloud) Centroid() vector.Vector {
	if len(p.Points) == 0 {
		return vector.Vector{}
	}
	centroid := &vector.Vector{}
	for dim := 0; dim < 3; dim++ {
		sum := float64(0)
		for v := range p.Points {
			sum += v.Array()[dim]
		}
		centroid.SetIndex(dim, sum / float64(len(p.Points)))
	}
	return *centroid
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