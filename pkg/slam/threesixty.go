package slam

import (
	log "github.com/sirupsen/logrus"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

func Threesixty(g *grid.Grid, bot Bot) {
	log.Info("Mapping...")
	done := false
	for !done {
		for i := 0.0; i <= 360; i += 15.0 {
			bot.Rotate(i)
			g.Apply(bot.Scan()...)
		}
	}
	log.Info("Done mapping.")
}
