package mapper

import (
	log "github.com/sirupsen/logrus"
	"math/rand"

	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	g.Apply(bot.Scan(360)...)
}

func RandomlyMap(g *grid.Grid, bot bot.Controller) {

	done := false
	for !done {

		startHeading := bot.Pose().Heading
		g.Apply(bot.Scan(360)...)
		bot.Heading(startHeading)

		// First try to move forward
		pt := coordinates.Add(bot.Pose().Location, coordinates.CompassRose{Heading: startHeading, Distance: g.CellSize()})
		if g.CellIsVacant(pt) {
			bot.Move(g.CellSize())
		} else {
			heading := bot.Pose().Heading + float64(rand.Intn(360)/45)*45.0
			bot.Heading(heading)
		}
	}
	log.Info("Done mapping.")
}
