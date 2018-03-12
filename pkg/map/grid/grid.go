package grid

import (
	log "github.com/sirupsen/logrus"
	"math"
	"sync"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Grid struct {
	grid                   sync.Map
	minX, minY, maxX, maxY int
	cellSize               distance.Distance
	path                   coordinates.CartesianSet
	newOdds                func() Odds
}

func New(cellSize distance.Distance) *Grid {
	return &Grid{
		grid:     sync.Map{},
		minX:     math.MaxInt64,
		maxX:     -math.MaxInt64,
		minY:     math.MaxInt64,
		maxY:     -math.MaxInt64,
		cellSize: cellSize,
		path:     coordinates.NewCartesianSet(),
		newOdds:  NewLogOdds,
	}
}

func (g *Grid) cell(c coordinates.Cartesian) coordinates.Cartesian {
	size := int(g.cellSize)
	return coordinates.Cartesian{
		X: (c.X / size) * size,
		Y: (c.Y / size) * size,
	}
}

func (g *Grid) Get(c coordinates.Cartesian) Odds {
	cell := g.cell(c)
	val, exists := g.grid.Load(cell)
	if exists {
		odds, ok := val.(Odds)
		if !ok {
			panic("could not cast Odds")
		}
		return odds
	}
	o := g.newOdds()
	g.grid.Store(cell, o)
	return o
}

func (g *Grid) Apply(rs ...sensor.Reading) {
	for _, r := range rs {
		g.apply(r)
	}
}

func (g *Grid) set(c coordinates.Cartesian, odds Odds) {
	cell := g.cell(c)
	if cell.X < g.minX {
		g.minX = cell.X
	}
	if cell.X > g.maxX {
		g.maxX = cell.X
	}
	if cell.Y < g.minY {
		g.minY = cell.Y
	}
	if cell.Y > g.maxY {
		g.maxY = cell.Y
	}
	g.grid.Store(cell, odds)
}

func (g *Grid) update(c coordinates.Cartesian, prob float64) {
	odds := g.Get(c)
	odds.Adjust(prob)
	g.set(c, odds)
}

func (g *Grid) apply(r sensor.Reading) {
	log.Debug("Applying reading ", r)
	g.path.Add(r.Pose.Location)
	v, o := r.Analysis()
	v.Range(func(coord coordinates.Cartesian) bool {
		prob := 0.1
		g.update(coord, prob)
		return false
	})
	o.Range(func(coord coordinates.Cartesian) bool {
		prob := 0.9
		g.update(coord, prob)
		return false
	})
	log.Debug("Reading vacant/occupied: ", v, o)

}

func (g *Grid) bounds() (minX, minY, maxX, maxY int) {
	log.Debugf("Grid boundaries: %d %d %d %d", g.minX, g.minY, g.maxX, g.maxY)
	// flip minY and maxY because of how the golang image library works.
	cellSize := int(g.cellSize)
	return g.minX - cellSize, -g.maxY - cellSize, g.maxX + cellSize, -g.minY - cellSize
}
