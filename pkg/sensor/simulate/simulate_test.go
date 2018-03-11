package simulate

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadingUltraSonicNorth(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 0, Y: 10})
	n := sensor.Pose{X: 0, Y: 0, Heading: 0}
	r := Reading(sensor.UltraSonic, n, occ)
	assert.InDelta(t, 10, r.Value, 1)
}

func TestReadingUltraSonicEast(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 10, Y: 0})
	e := sensor.Pose{X: 0, Y: 0, Heading: 90}
	r := Reading(sensor.UltraSonic, e, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func TestReadingUltraSonicSouth(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 0, Y: -10})
	s := sensor.Pose{X: 0, Y: 0, Heading: 180}
	r := Reading(sensor.UltraSonic, s, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func TestReadingUltraSonicWest(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: -10, Y: 0})
	w := sensor.Pose{X: 0, Y: 0, Heading: 270}
	r := Reading(sensor.UltraSonic, w, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func TestReadingOffset(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 10, Y: 10})
	p := sensor.Pose{X: 10, Y: 0, Heading: 0}
	r := Reading(sensor.UltraSonic, p, occ)
	assert.InDelta(t, 10, r.Value, 0)
}
