package pointcloud

import (
	"gonum.org/v1/gonum/mat"
)

type Transformation struct {
	Translation Point
	Rotation    mat.Matrix
}

func (t *Transformation) Transform(p *PointCloud) *PointCloud {
	result := &PointCloud{points: make([]Point, len(p.points))}
	for i, pt := range p.points {
		col := pt.ColMatrix()
		var mult mat.Dense
		mult.Mul(t.Rotation, col)
		afterRotation := matToPoint(&mult)
		transformed := Subtract(afterRotation, t.Translation)
		result.points[i] = transformed
	}
	return result
}
