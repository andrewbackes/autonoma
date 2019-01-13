package pointcloud

import (
	"math"
)

// Point in space.
type Point struct {
	x []int
}

// PointCloud is a collection of points.
type PointCloud struct {
	points []Point
}

// Add a point to the cloud.
func (p *PointCloud) Add(pt Point) {
	p.points = append(p.points, pt)
}

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

func dist(p, q Point) float64 {
	sum := float64(0)
	for i := 0; i < len(p.x); i++ {
		sum += math.Pow(float64(p.x[i]-q.x[i]), 2)
	}
	return math.Sqrt(sum)
}
