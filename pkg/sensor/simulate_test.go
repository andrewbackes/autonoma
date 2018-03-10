package sensor

import (
	"github.com/andrewbackes/autonoma/pkg/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateReadingUltraSonicNorth(t *testing.T) {
	occ := set.New()
	occ.Add("0,10")
	n := Pose{X: 0, Y: 0, Heading: 0}
	r := SimulateReading(UltraSonic, n, occ)
	assert.InDelta(t, 10, r.Value, 1)
}

func TestSimulateReadingUltraSonicEast(t *testing.T) {
	occ := set.New()
	occ.Add("10,0")
	e := Pose{X: 0, Y: 0, Heading: 90}
	r := SimulateReading(UltraSonic, e, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func TestSimulateReadingUltraSonicSouth(t *testing.T) {
	occ := set.New()
	occ.Add("0,-10")
	s := Pose{X: 0, Y: 0, Heading: 180}
	r := SimulateReading(UltraSonic, s, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func TestSimulateReadingUltraSonicWest(t *testing.T) {
	occ := set.New()
	occ.Add("-10,0")
	w := Pose{X: 0, Y: 0, Heading: 270}
	r := SimulateReading(UltraSonic, w, occ)
	assert.InDelta(t, 10, r.Value, 0)
}

func occSet() set.Set {
	occ := set.New()
	occ.Add("15,0")
	occ.Add("0,13")
	occ.Add("-1,10")
	occ.Add("1,8")
	return occ
}
