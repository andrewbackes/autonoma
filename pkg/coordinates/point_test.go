package coordinates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToVectorY(t *testing.T) {
	p := Point{
		Orientation: Euler{
			Yaw: 90,
		},
		Vector: Vector{
			X: 0,
			Y: 1,
		},
	}
	v := p.ToVector()
	assert.Equal(t, Vector{X: 1, Y: 0}, v)
}

func TestToVector45(t *testing.T) {
	p := Point{
		Orientation: Euler{
			Yaw: 45,
		},
		Vector: Vector{
			X: 1,
			Y: 1,
		},
	}
	v := p.ToVector()
	assert.Equal(t, Vector{X: 1, Y: 0}, v)
}

func TestToVectorX(t *testing.T) {
	p := Point{
		Orientation: Euler{
			Yaw: 90,
		},
		Vector: Vector{
			X: 1,
			Y: 0,
		},
	}
	v := p.ToVector()
	assert.Equal(t, Vector{X: 0, Y: -1}, v)
}

func TestToVector90(t *testing.T) {
	p := Point{
		Orientation: Euler{
			Yaw: 90,
		},
		Vector: Vector{
			X: 1,
			Y: 1,
		},
	}
	v := p.ToVector()
	assert.Equal(t, Vector{X: 1, Y: -1}, v)
}
