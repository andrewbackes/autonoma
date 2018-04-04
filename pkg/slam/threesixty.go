package slam

import (
	log "github.com/sirupsen/logrus"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

func Threesixty(g *grid.Grid, bot Bot) {
	log.Info("Mapping...")
	g.Apply(bot.Scan()...)
	log.Info("Done mapping.")
}
