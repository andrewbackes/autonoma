package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolarRotate0(t *testing.T) {
	v1 := PolarLikeCoordToPoint(0, 1)
	assert.InDelta(t, 0, v1.X, 0.01)
	assert.InDelta(t, 1, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestPolarRotate180(t *testing.T) {
	v1 := PolarLikeCoordToPoint(180, 1)
	assert.InDelta(t, 0, v1.X, 0.01)
	assert.InDelta(t, -1, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestPolarRotate90(t *testing.T) {
	v1 := PolarLikeCoordToPoint(90, 1)
	assert.InDelta(t, 1, v1.X, 0.01)
	assert.InDelta(t, 0, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestPolarRotate270(t *testing.T) {
	v1 := PolarLikeCoordToPoint(270, 1)
	assert.InDelta(t, -1, v1.X, 0.01)
	assert.InDelta(t, 0, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestPolarRotate45(t *testing.T) {
	v1 := PolarLikeCoordToPoint(45, 1)
	assert.InDelta(t, 0.7, v1.X, 0.01)
	assert.InDelta(t, 0.7, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestAdd(t *testing.T) {
	v1 := Point{1, 2, 3}
	v2 := Point{10, 20, 30}
	assert.Equal(t, Point{11, 22, 33}, Add(v1, v2))
}

func TestRotate90(t *testing.T) {
	v1 := Point{1, 0, 0}
	v2 := Rotate(v1, 90)
	t.Log(v2)
	assert.InDelta(t, 0, v2.X, 0.01)
	assert.InDelta(t, -1, v2.Y, 0.01)
	assert.InDelta(t, 0, v2.Z, 0.01)
}

func TestRotate180(t *testing.T) {
	v1 := Point{1, 0, 0}
	v2 := Rotate(v1, 180)
	t.Log(v2)
	assert.InDelta(t, -1, v2.X, 0.01)
	assert.InDelta(t, 0, v2.Y, 0.01)
	assert.InDelta(t, 0, v2.Z, 0.01)
}

func TestDistance(t *testing.T) {
	x := Point{-4, 2, 2}
	y := Point{1, -4, 1}
	expected := 7.874007874011811
	actual := Distance(x, y)
	assert.InDelta(t, expected, actual, 0.001)
}

func TestRemoveOutliers(t *testing.T) {
	vs := []Point{
		{1, 0, 0},
		{2, 0, 0},
		{4, 0, 0},
	}
	actual := RemoveOutliers(vs, 1, 1)
	assert.Equal(t, 2, len(actual))
}
