package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate0(t *testing.T) {
	v1 := PolarLikeCoordToVector(0, 1)
	assert.InDelta(t, 0, v1.X, 0.01)
	assert.InDelta(t, 1, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestRotate180(t *testing.T) {
	v1 := PolarLikeCoordToVector(180, 1)
	assert.InDelta(t, 0, v1.X, 0.01)
	assert.InDelta(t, -1, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestRotate90(t *testing.T) {
	v1 := PolarLikeCoordToVector(90, 1)
	assert.InDelta(t, 1, v1.X, 0.01)
	assert.InDelta(t, 0, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestRotate270(t *testing.T) {
	v1 := PolarLikeCoordToVector(270, 1)
	assert.InDelta(t, -1, v1.X, 0.01)
	assert.InDelta(t, 0, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}

func TestRotate45(t *testing.T) {
	v1 := PolarLikeCoordToVector(45, 1)
	assert.InDelta(t, 0.7, v1.X, 0.01)
	assert.InDelta(t, 0.7, v1.Y, 0.01)
	assert.InDelta(t, 0, v1.Z, 0.01)
}
