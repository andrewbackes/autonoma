package grid

import (
	log "github.com/sirupsen/logrus"
	"math"
	"sync"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Grid struct {
	grid                   sync.Map
	minX, minY, maxX, maxY int
	path                   coordinates.CartesianSet
	newOdds                func() Odds
}

func New() *Grid {
	return &Grid{
		grid:    sync.Map{},
		minX:    math.MaxInt64,
		maxX:    -math.MaxInt64,
		minY:    math.MaxInt64,
		maxY:    -math.MaxInt64,
		path:    coordinates.NewCartesianSet(),
		newOdds: NewLogOdds,
	}
}

func (g *Grid) Get(c coordinates.Cartesian) Odds {
	val, exists := g.grid.Load(c)
	if exists {
		odds, ok := val.(Odds)
		if !ok {
			panic("could not cast Odds")
		}
		return odds
	}
	o := g.newOdds()
	g.grid.Store(c, o)
	return o
}

func (g *Grid) Apply(rs ...sensor.Reading) {
	for _, r := range rs {
		g.apply(r)
	}
}

func (g *Grid) set(c coordinates.Cartesian, odds Odds) {
	if c.X < g.minX {
		g.minX = c.X
	}
	if c.X > g.maxX {
		g.maxX = c.X
	}
	if c.Y < g.minY {
		g.minY = c.Y
	}
	if c.Y > g.maxY {
		g.maxY = c.Y
	}
	g.grid.Store(c, odds)
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
	return g.minX, -g.maxY, g.maxX, -g.minY
}
