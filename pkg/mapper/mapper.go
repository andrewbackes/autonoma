package mapper

import (
	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
)

// Map the area with the provided bot.
func Map(g *grid.Grid, bot bot.Controller) {
	//bot.Heading(45)
	g.Apply(bot.Scan(90)...)
}
