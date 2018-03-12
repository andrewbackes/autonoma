package mapper

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"
	"math/rand"

	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

const (
	vacantThreshold = 0.1
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	g.Apply(bot.Scan(360)...)
}

func RandomlyMap(g *grid.Grid, bot bot.Controller) {
	lookAhead := 15 * distance.Centimeter
	moveAhead := 5 * distance.Centimeter

	startHeading := bot.Pose().Heading
	g.Apply(bot.Scan(360)...)
	bot.Heading(startHeading)

	done := false
	for !done {
		// First try to move forward
		line := g.Line(bot.Pose().Location, coordinates.CompassRose{Heading: startHeading, Distance: lookAhead})
		blocked := true
		line.Range(func(coor coordinates.Cartesian) bool {
			if g.Get(coor).Probability() < vacantThreshold {
				blocked = false
				return false
			}
			return true
		})
		if blocked {
			// rotate randomly
			heading := float64(rand.Intn(360)/45) * 15.0
			bot.Heading(heading)
			g.Apply(bot.Scan(360)...)
			bot.Heading(heading)
		} else {
			log.Debug("Moving Forward")
			bot.Move(moveAhead)
		}

		/*
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
		*/
	}
	log.Info("Done mapping.")
}
