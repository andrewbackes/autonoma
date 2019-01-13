package pointcloud

import (
	"math"
)

// ICP is an implementation of iterative closest point.
func ICP(source, target *PointCloud) *Transformation {
	return &Transformation{}
}

func closestPoints(source, target *PointCloud) (*PointCloud, float64) {
	if len(target.points) < len(source.points) {
		return source, math.MaxFloat64
	}
	closest := &PointCloud{points: make([]Point, len(source.points))}
	distances := float64(0)
	for _, p := range source.points {
		c, d := target.closest(p)
		distances += d
		closest.Add(c)
	}
	aveDist := distances / float64(len(source.points))
	return closest, aveDist
}

// Distance is the shortest distance from a point to the point cloud.
func (p *PointCloud) closest(to Point) (Point, float64) {
	if len(p.points) == 0 {
		return Point{}, math.MaxFloat64
	}
	min := dist(p.points[0], to)
	closest := p.points[0]
	for i := 1; i < len(p.points); i++ {
		d := dist(p.points[i], to)
		if d < min {
			min = d
			closest = p.points[i]
		}
	}
	return closest, min
}
