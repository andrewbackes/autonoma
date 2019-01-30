package grid

import (
	log "github.com/sirupsen/logrus"
	"math"
	"sync"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
)

const (
	vacantThreshold           = 0.15
	occupiedThreshold         = 0.85
	increaseProbabilityAmount = 0.99
	decreaseProbabilityAmount = 0.01
)

type Grid struct {
	grid                   sync.Map
	minX, minY, maxX, maxY int
	cellSize               distance.Distance
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
		newOdds:  NewLogOdds,
	}
}

func (g *Grid) CellSize() distance.Distance {
	return g.cellSize
}

// Get the odds of an obsticle being at a vector.
func (g *Grid) Get(c coordinates.Vector) Odds {
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

// Increase the probability of an obsticle being at the vector.
func (g *Grid) Increase(v coordinates.Vector) {
	g.update(v, increaseProbabilityAmount)
}

// Decrease the probability of an obsticle being at the vector.
func (g *Grid) Decrease(v coordinates.Vector) {
	g.update(v, decreaseProbabilityAmount)
}

func (g *Grid) set(c coordinates.Vector, odds Odds) {
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

func (g *Grid) update(c coordinates.Vector, prob float64) {
	odds := g.Get(c)
	odds.Adjust(prob)
	g.set(c, odds)
}

// Bounds of the grid.
func (g *Grid) Bounds() (minX, minY, maxX, maxY int) {
	log.Debugf("Grid boundaries: %d %d %d %d", g.minX, g.minY, g.maxX, g.maxY)
	// flip minY and maxY because of how the golang image library works.
	cellSize := int(g.cellSize)
	return g.minX - cellSize, -(g.maxY + cellSize), g.maxX + cellSize, -(g.minY - cellSize)
}

/*
func (g *Grid) Apply(rs ...sensor.Reading) {
	for _, r := range rs {
		g.apply(r)
	}
}

func (g *Grid) apply(r sensor.Reading) {
	log.Debug("Applying reading ", r)
	vac, occ := r.Analysis()
	// mark the probabilities:
	occ.Range(func(coord coordinates.Vector) bool {
		prob := 0.99
		g.update(coord, prob)
		return false
	})
	vac.Range(func(coord coordinates.Vector) bool {
		prob := 0.1
		g.update(coord, prob)
		return false
	})
	log.Debug("Reading vacant/occupied: ", vac, occ)
}

func (g *Grid) Line(origin coordinates.Vector, bearing coordinates.CompassRose) coordinates.VectorSet {
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
*/
/*
func (g *Grid) Vacant(coords coordinates.VectorSet) bool {
	vacant := true
	coords.Range(func(coor coordinates.Vector) bool {
		occupied := g.Get(coor).Probability() > vacantThreshold
		if occupied {
			vacant = false
			return false
		}
		return true
	})
	return vacant
}
*/
