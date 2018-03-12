package mapper

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"

	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

const (
	vacantThreshold = 0.15
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	g.Apply(bot.Scan(360)...)

}

func RandomlyMap(g *grid.Grid, bot bot.Controller) {
	lookAhead := 10 * distance.Centimeter
	moveAhead := 5 * distance.Centimeter
	done := false
	for !done {
		startHeading := bot.Pose().Heading
		g.Apply(bot.Scan(360)...)
		bot.Heading(startHeading)

		// First try to move forward
		dest := coordinates.Add(bot.Pose().Location, coordinates.CompassRose{Heading: startHeading, Distance: lookAhead})
		log.Debug("Destination ", dest)
		if g.Get(dest).Probability() < vacantThreshold {
			log.Debug("Moving Forward")
			bot.Move(moveAhead)
			continue
		}

		// find a non-occupied location
		var bearing float64
		startBearing := float64(rand.Intn(360))
		for b := startBearing; ; b += 15 {
			bearing = math.Mod(b, 360)
			dest = coordinates.Add(bot.Pose().Location, coordinates.CompassRose{Heading: bearing, Distance: lookAhead})
			if g.Get(dest).Probability() < vacantThreshold {
				log.Debug("Moving Randomly")
				bot.Heading(bearing)
				bot.Move(moveAhead)
				break
			}
			if b >= startBearing+360 {
				done = true
				break
			}
		}
	}
	log.Info("Done mapping.")
}
