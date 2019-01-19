package pointcloud

import (
	"gonum.org/v1/gonum/mat"
)

type Transformation struct {
	Translation Point
	Rotation    mat.Matrix
}

func (t *Transformation) Transform(p *PointCloud) *PointCloud {
	result := &PointCloud{Points: make([]Point, len(p.Points))}
	for i, pt := range p.Points {
		col := pt.ColMatrix()
		var mult mat.Dense
		mult.Mul(t.Rotation, col)
		afterRotation := matToPoint(&mult)
		transformed := Subtract(afterRotation, t.Translation)
		result.Points[i] = transformed
	}
	return result
}
