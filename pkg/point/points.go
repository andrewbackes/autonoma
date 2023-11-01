package point

import (
	"gonum.org/v1/gonum/mat"
)

func RemoveOutliers(vs []Point, ptCount int, perDist float64) []Point {
	ret := make([]Point, 0, len(vs))
	for _, v := range vs {
		count := 0
		for _, w := range vs {
			if v != w && Distance(v, w) <= perDist {
				count++
			}

		}
		if count >= ptCount {
			ret = append(ret, v)
		}
	}
	return ret
}

func AveDistance(vs, ws []Point) float64 {
	if len(vs) != len(ws) {
		panic("can not find average distance of different length sets")
	}
	sum := float64(0)
	for i := range vs {
		sum += Distance(vs[i], ws[i])
	}
	return sum / float64(len(vs))
}

func Centroid(vs []Point) Point {
	if len(vs) == 0 {
		panic("can not find centroid of empty Point array")
	}
	centroid := &Point{}
	for dim := 0; dim < Dimensions; dim++ {
		sum := float64(0)
		for _, v := range vs {
			sum += v.Index(dim)
		}
		centroid.SetIndex(dim, sum/float64(len(vs)))
	}
	return *centroid
}

func Copy(source []Point) []Point {
	c := make([]Point, len(source))
	for i, v := range source {
		c[i] = v
	}
	return c
}

// Shift in place by a Point.
func Shift(Points []Point, by Point) {
	for i, v := range Points {
		Points[i] = Subtract(v, by)
	}
}

func Matrix(Points []Point) mat.Matrix {
	if len(Points) == 0 {
		panic("can not create matrix from an empty array of Points")
	}
	m := mat.NewDense(Dimensions, len(Points), nil)
	for col, v := range Points {
		for row, val := range v.Array() {
			m.Set(row, col, float64(val))
		}
	}
	return m
}
