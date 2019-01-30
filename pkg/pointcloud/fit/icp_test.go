package fit

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
	"github.com/andrewbackes/autonoma/pkg/vector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestICPTranslation(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{1, 3, 0})
	env.Add(vector.Vector{1, 5, 0})
	env.Add(vector.Vector{1, 8, 0})
	src := make([]vector.Vector, 3)
	src[0] = vector.Vector{3, 1, 0}
	src[1] = vector.Vector{3, 3, 0}
	src[2] = vector.Vector{3, 6, 0}
	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}

func TestICPRotation90(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{0, 3, 0})
	env.Add(vector.Vector{0, 5, 0})
	env.Add(vector.Vector{0, 8, 0})
	src := make([]vector.Vector, 3)
	src[0] = vector.Vector{3, 0, 0}
	src[1] = vector.Vector{5, 0, 0}
	src[2] = vector.Vector{8, 0, 0}
	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}

func TestICPRotation(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{1, 1, 0})
	env.Add(vector.Vector{1, 3, 0})
	env.Add(vector.Vector{1, 6, 0})
	src := make([]vector.Vector, 3)
	src[0] = vector.Vector{-0.32, 1.38, 0}
	src[1] = vector.Vector{-2.02, 2.44, 0}
	src[2] = vector.Vector{-4.57, 4.03, 0}
	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}

func TestICPBox(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{0, 0, 0})
	env.Add(vector.Vector{10, 10, 0})
	env.Add(vector.Vector{0, 10, 0})
	env.Add(vector.Vector{10, 0, 0})

	src := make([]vector.Vector, 4)
	src[0] = vector.Vector{2, 2, 0}
	src[1] = vector.Vector{12, 12, 0}
	src[2] = vector.Vector{2, 12, 0}
	src[3] = vector.Vector{12, 2, 0}

	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}

func TestICPBoxRotate30(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{1, 1, 0})
	env.Add(vector.Vector{11, 11, 0})
	env.Add(vector.Vector{1, 11, 0})
	env.Add(vector.Vector{11, 1, 0})

	src := make([]vector.Vector, 4)
	src[0] = vector.Vector{1.3660254037844, 0.36602540378444, 0}
	src[1] = vector.Vector{15.026279441629, 4.0262794416288, 0}
	src[2] = vector.Vector{6.3660254037844, 9.0262794416288, 0}
	src[3] = vector.Vector{10.026279441629, -4.6339745962156, 0}

	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}

func TestICPBoxRotate10(t *testing.T) {
	env := pointcloud.New()
	env.Add(vector.Vector{1, 1, 0})
	env.Add(vector.Vector{11, 11, 0})
	env.Add(vector.Vector{1, 11, 0})
	env.Add(vector.Vector{11, 1, 0})

	src := make([]vector.Vector, 4)
	src[0] = vector.Vector{1.1584559306791, 0.81115957534528, 0}
	src[1] = vector.Vector{12.743015237471, 8.9227553287981, 0}
	src[2] = vector.Vector{2.8949377073484, 10.659237105467, 0}
	src[3] = vector.Vector{11.006533460801, -0.92532220132403, 0}

	result, _, e := ICP(src, vector.Vector{}, env, 0.5, 10)
	fmt.Println("result: ", result, "error: ", e)
	assert.InDelta(t, 0, e, 0.01)
}
