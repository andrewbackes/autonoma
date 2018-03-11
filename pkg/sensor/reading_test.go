package sensor

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadingAnalysis(t *testing.T) {
	r := Reading{
		Value:  15,
		Sensor: IRDistance,
		Pose: Pose{
			X:       0,
			Y:       0,
			Heading: 0,
		},
	}
	vac, occ := r.Analysis()
	assert.Equal(t, 1, len(occ))
	assert.True(t, occ.Contains(coordinates.Cartesian{X: 0, Y: 15}))
	assert.Equal(t, 5, len(vac))
}
