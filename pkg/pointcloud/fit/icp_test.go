package fit

import (
	"fmt"
	"testing"
)

func TestICPTranslation(t *testing.T) {
	env := &PointCloud{Points: make([]Point, 0)}
	env.Add(NewPoint(1, 1))
	env.Add(NewPoint(1, 3))
	env.Add(NewPoint(1, 6))

	src := &PointCloud{Points: make([]Point, 0)}
	src.Add(NewPoint(3, 3))
	src.Add(NewPoint(3, 5))
	src.Add(NewPoint(3, 8))
	result, _, e := ICP(src, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
}

func TestICPRotation(t *testing.T) {
	env := &PointCloud{Points: make([]Point, 0)}
	env.Add(NewPoint(1, 1))
	env.Add(NewPoint(1, 3))
	env.Add(NewPoint(1, 6))

	src := &PointCloud{Points: make([]Point, 0)}
	src.Add(NewPoint(-0.32, 1.38))
	src.Add(NewPoint(-2.02, 2.44))
	src.Add(NewPoint(-4.57, 4.03))
	result, _, e := ICP(src, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
}

func TestICPBox(t *testing.T) {
	env := &PointCloud{Points: make([]Point, 0)}
	env.Add(NewPoint(0, 0))
	env.Add(NewPoint(10, 10))
	env.Add(NewPoint(0, 10))
	env.Add(NewPoint(10, 0))

	src := &PointCloud{Points: make([]Point, 0)}
	src.Add(NewPoint(2, 2))
	src.Add(NewPoint(12, 12))
	src.Add(NewPoint(2, 12))
	src.Add(NewPoint(12, 2))
	result, _, e := ICP(src, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
}

func TestICPBoxRotate30(t *testing.T) {
	env := &PointCloud{Points: make([]Point, 0)}
	env.Add(NewPoint(1, 1))
	env.Add(NewPoint(11, 11))
	env.Add(NewPoint(1, 11))
	env.Add(NewPoint(11, 1))

	src := &PointCloud{Points: make([]Point, 0)}
	src.Add(NewPoint(1.3660254037844, 0.36602540378444))
	src.Add(NewPoint(15.026279441629, 4.0262794416288))
	src.Add(NewPoint(6.3660254037844, 9.0262794416288))
	src.Add(NewPoint(10.026279441629, -4.6339745962156))
	result, _, e := ICP(src, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
}

func TestICPBoxRotate10(t *testing.T) {
	env := &PointCloud{Points: make([]Point, 0)}
	env.Add(NewPoint(1, 1))
	env.Add(NewPoint(11, 11))
	env.Add(NewPoint(1, 11))
	env.Add(NewPoint(11, 1))

	src := &PointCloud{Points: make([]Point, 0)}
	src.Add(NewPoint(1.1584559306791, 0.81115957534528))
	src.Add(NewPoint(12.743015237471, 8.9227553287981))
	src.Add(NewPoint(2.8949377073484, 10.659237105467))
	src.Add(NewPoint(11.006533460801, -0.92532220132403))
	result, _, e := ICP(src, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
}
