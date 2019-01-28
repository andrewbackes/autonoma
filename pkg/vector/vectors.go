package vector

import (
	"gonum.org/v1/gonum/mat"
)

func RemoveOutliers(vs []Vector, ptCount int, perDist float64) []Vector {
	ret := make([]Vector, 0, len(vs))
	for _, v := range vs {
		count := 0
		for _, w := range vs {
			if Distance(v, w) < perDist {
				count++
			}
		}
		if count >= ptCount {
			ret = append(ret, v)
		}
	}
	return ret
}

func AveDistance(vs, ws []Vector) float64 {
	if len(vs) != len(ws) {
		panic("can not find average distance of different length sets")
	}
	sum := float64(0)
	for i := range vs {
		sum += Distance(vs[i], ws[i])
	}
	return sum / float64(len(vs))
}

func Centroid(vs []Vector) Vector {
	if len(vs) == 0 {
		panic("can not find centroid of empty vector array")
	}
	centroid := &Vector{}
	for dim := 0; dim < Dimensions; dim++ {
		sum := float64(0)
		for _, v := range vs {
			sum += v.Index(dim)
		}
		centroid.SetIndex(dim, sum/float64(len(vs)))
	}
	return *centroid
}

func Copy(source []Vector) []Vector {
	c := make([]Vector, len(source))
	for i, v := range source {
		c[i] = v
	}
	return c
}

// Shift in place by a vector.
func Shift(vectors []Vector, by Vector) {
	for i, v := range vectors {
		vectors[i] = Subtract(v, by)
	}
}

func Matrix(vectors []Vector) mat.Matrix {
	if len(vectors) == 0 {
		panic("can not create matrix from an empty array of vectors")
	}
	m := mat.NewDense(Dimensions, len(vectors), nil)
	for col, v := range vectors {
		for row, val := range v.Array() {
			m.Set(row, col, float64(val))
		}
	}
	return m
}
