package pointcloud

import (
	"fmt"
	"testing"
)

func TestICP(t *testing.T) {
	env := &PointCloud{points: make([]Point, 0)}
	env.Add(NewPoint(1, 1))
	env.Add(NewPoint(1, 3))
	env.Add(NewPoint(1, 6))

	src := &PointCloud{points: make([]Point, 0)}
	src.Add(NewPoint(3, 3))
	src.Add(NewPoint(3, 5))
	src.Add(NewPoint(3, 8))
	result := ICP(src, env, 0.1, 10)
	fmt.Println(result)
}
