package transformation

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Transformation struct {
	Translation Point
	Rotation    mat.Matrix
}

func (t Transformation) String() string {
	s := fmt.Sprint(t.Translation, "\n")
	rows, cols := t.Rotation.Dims()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			v := t.Rotation.At(r, c)
			s = fmt.Sprintf("%v %v", s, v)
		}
		s = s + "\n"
	}
	return s
}

func NewTransformation() *Transformation {
	return &Transformation{
		Translation: NewPoint(0, 0, 0),
		Rotation: mat.NewDense(3, 3, []float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		}),
	}
}

func (t *Transformation) Transform(p *PointCloud) *PointCloud {
	result := &PointCloud{Points: make([]Point, len(p.Points))}
	for i, pt := range p.Points {
		result.Points[i] = t.TransformPoint(pt)
		//fmt.Println(pt, "->", result.Points[i])
	}
	return result
}

func (t *Transformation) TransformPoint(pt Point) Point {
	//fmt.Println("---> transformation", t)
	//fmt.Println("---> pt", pt)
	col := pt.ColMatrix()
	var mult mat.Dense
	mult.Mul(t.Rotation, col)
	afterRotation := matToPoint(&mult)
	transformed := Subtract(afterRotation, t.Translation)
	return transformed
}
