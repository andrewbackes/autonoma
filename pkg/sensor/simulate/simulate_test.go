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
	n := coordinates.NewPose(0, 0, 0)
	r := Reading(sensor.UltraSonic, n, occ)
	assert.InDelta(t, 10, float64(r.Value), 1)
}

func TestReadingUltraSonicEast(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 10, Y: 0})
	e := coordinates.NewPose(0, 0, 90)
	r := Reading(sensor.UltraSonic, e, occ)
	assert.InDelta(t, 10, float64(r.Value), 0)
}

func TestReadingUltraSonicSouth(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 0, Y: -10})
	s := coordinates.NewPose(0, 0, 180)
	r := Reading(sensor.UltraSonic, s, occ)
	assert.InDelta(t, 10, float64(r.Value), 0)
}

func TestReadingUltraSonicWest(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: -10, Y: 0})
	w := coordinates.NewPose(0, 0, 270)
	r := Reading(sensor.UltraSonic, w, occ)
	assert.InDelta(t, 10, float64(r.Value), 0)
}

func TestReadingOffset(t *testing.T) {
	occ := coordinates.NewCartesianSet()
	occ.Add(coordinates.Cartesian{X: 10, Y: 10})
	p := coordinates.NewPose(10, 0, 0)
	r := Reading(sensor.UltraSonic, p, occ)
	assert.InDelta(t, 10, float64(r.Value), 0)
}
