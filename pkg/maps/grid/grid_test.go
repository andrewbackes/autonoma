package grid

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCellOccupied(t *testing.T) {
	g := New(3)
	max := LogOdds{odds: maxLogOdds}
	pt := coordinates.Cartesian{X: 0, Y: 0}
	g.set(pt, &max)
	assert.True(t, g.CellIsOccupied(pt))
	assert.False(t, g.CellIsVacant(pt))
	pt2 := coordinates.Cartesian{X: 3, Y: 3}
	assert.False(t, g.CellIsOccupied(pt2))
	assert.False(t, g.CellIsVacant(pt2))
}
