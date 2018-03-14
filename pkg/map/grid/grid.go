package grid

import (
	log "github.com/sirupsen/logrus"
	"math"
	"sync"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

const (
	vacantThreshold   = 0.15
	occupiedThreshold = 0.85
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

func (g *Grid) CellSize() distance.Distance {
	return g.cellSize
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
	vac, occ := r.Analysis()
	// mark the probabilities:
	occ.Range(func(coord coordinates.Cartesian) bool {
		prob := 0.99
		g.update(coord, prob)
		return false
	})
	vac.Range(func(coord coordinates.Cartesian) bool {
		prob := 0.1
		g.update(coord, prob)
		return false
	})
	log.Debug("Reading vacant/occupied: ", vac, occ)
}

func (g *Grid) bounds() (minX, minY, maxX, maxY int) {
	log.Debugf("Grid boundaries: %d %d %d %d", g.minX, g.minY, g.maxX, g.maxY)
	// flip minY and maxY because of how the golang image library works.
	cellSize := int(g.cellSize)
	return g.minX - cellSize, -(g.maxY + cellSize), g.maxX + cellSize, -(g.minY - cellSize)
}

func (g *Grid) Line(origin coordinates.Cartesian, bearing coordinates.CompassRose) coordinates.CartesianSet {
	coors := coordinates.NewCartesianSet()
	for d := distance.Distance(0); d <= bearing.Distance; d += 0.5 {
		interim := coordinates.CompassRose{
			Distance: d,
			Heading:  bearing.Heading,
		}
		dest := coordinates.Add(origin, interim)
		coors.Add(dest)
	}
	return coors
}

func (g *Grid) Vacant(coords coordinates.CartesianSet) bool {
	vacant := true
	coords.Range(func(coor coordinates.Cartesian) bool {
		occupied := g.Get(coor).Probability() > vacantThreshold
		if occupied {
			vacant = false
			return false
		}
		return true
	})
	return vacant
}

func (g *Grid) CellIsVacant(c coordinates.Cartesian) bool {
	max := g.cellMax(c)
	return max <= vacantThreshold
}

func (g *Grid) CellIsOccupied(c coordinates.Cartesian) bool {
	max := g.cellMax(c)
	return max >= occupiedThreshold
}

func (g *Grid) cellMax(c coordinates.Cartesian) float64 {
	coords := coordinates.Cell(c, g.cellSize)
	max := float64(-1)
	coords.Range(func(coor coordinates.Cartesian) bool {
		prob := g.Get(coor).Probability()
		if prob > max {
			max = prob
		}
		return false
	})
	return max
}
