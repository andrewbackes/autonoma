package mapper

import (
	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"math"
	"math/rand"
)

const (
	vacantThreshold = 0.15
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	g.Apply(bot.Scan(360)...)

}

func RandomlyMap(g *grid.Grid, bot bot.Controller) {
	done := false
	for !done {
		g.Apply(bot.Scan(360)...)
		var bearing float64
		// find a non-occupied location
		startBearing := float64(rand.Intn(360))
		for b := startBearing; ; b++ {
			bearing = math.Mod(b, 360)
			dest := coordinates.Add(bot.Pose().Location, coordinates.CompassRose{Heading: bearing, Distance: 10})
			if g.Get(dest).Probability() < vacantThreshold {
				break
			}
			if b >= startBearing+360 {
				done = true
				break
			}
		}
		bot.Heading(bearing)
		bot.Move(5)
	}
}
