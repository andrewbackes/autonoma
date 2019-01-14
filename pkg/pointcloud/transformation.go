package pointcloud

import (
	"gonum.org/v1/gonum/mat"
)

type Transformation struct {
	Translation Point
	Rotation    mat.Matrix
}

func (t *Transformation) Transform(p *PointCloud) *PointCloud {
	return p
}
