package slam

import (
	log "github.com/sirupsen/logrus"
	"math/rand"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

func RandomlyMap(g *grid.Grid, bot Bot) {

	done := false
	for !done {

		startHeading := bot.Pose().Heading
		g.Apply(bot.Scan()...)
		bot.Rotate(startHeading)

		// First try to move forward
		pt := coordinates.Add(bot.Pose().Location, coordinates.CompassRose{Heading: startHeading, Distance: g.CellSize()})
		if g.CellIsVacant(pt) {
			bot.Move(g.CellSize())
			p := g.LocalizePose(bot.Pose())
			bot.SetPose(p)
		} else {
			heading := bot.Pose().Heading + float64(rand.Intn(360)/45)*45.0
			bot.Rotate(heading)
		}
	}
	log.Info("Done mapping.")
}
