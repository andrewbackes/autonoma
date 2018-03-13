package mapper

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"

	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	g.Apply(bot.Scan(360)...)
}

func RandomlyMap(g *grid.Grid, bot bot.Controller) {
	lookAhead := 10 * distance.Centimeter
	moveAhead := 5 * distance.Centimeter

	startHeading := bot.Pose().Heading
	g.Apply(bot.Scan(360)...)
	bot.Heading(startHeading)

	done := false
	for !done {
		// First try to move forward
		line := g.Line(bot.Pose().Location, coordinates.CompassRose{Heading: startHeading, Distance: lookAhead})
		if g.Vacant(line) {
			log.Debug("Moving Forward")
			bot.Move(moveAhead)
		} else {
			heading := bot.Pose().Heading + 15 // float64(rand.Intn(360)/45) * 15.0
			bot.Heading(heading)
			g.Apply(bot.Scan(360)...)
			bot.Heading(heading)
		}
	}
	log.Info("Done mapping.")
}
