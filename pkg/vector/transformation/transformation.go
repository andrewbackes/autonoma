// Package transformation handles linear transformations on vectors. It allows you to translate and rotate a vector.
package transformation

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/vector"
	"gonum.org/v1/gonum/mat"
)

type Transformation struct {
	Translation vector.Vector
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
		Translation: vector.Vector{},
		Rotation: mat.NewDense(3, 3, []float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		}),
	}
}

func (t *Transformation) Apply(v vector.Vector) vector.Vector {
	col := v.Matrix()
	var mult mat.Dense
	mult.Mul(t.Rotation, col)
	afterRotation := vector.FromMatrix(&mult)
	transformed := vector.Subtract(afterRotation, t.Translation)
	return transformed
}
