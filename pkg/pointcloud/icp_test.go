package pointcloud

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
