package grid

import (
	log "github.com/sirupsen/logrus"
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Grid struct {
	grid    map[coordinates.Cartesian]Odds
	path    coordinates.CartesianSet
	newOdds func() Odds
}

func New() Grid {
	return Grid{
		grid:    make(map[coordinates.Cartesian]Odds),
		path:    coordinates.NewCartesianSet(),
		newOdds: NewLogOdds,
	}
}

func (g Grid) set(c coordinates.Cartesian, odds Odds) {
	g.grid[c] = odds
}

func (g Grid) Get(c coordinates.Cartesian) Odds {
	odds, exists := g.grid[c]
	if exists {
		return odds
	}
	o := g.newOdds()
	g.grid[c] = o
	return o
}

func (g Grid) Adjust(c coordinates.Cartesian, prob float64) {
	odds := g.Get(c)
	odds.Adjust(prob)
	g.set(c, odds)
}

func (g Grid) Add(rs ...sensor.Reading) {
	for _, r := range rs {
		g.add(r)
	}
}

func (g Grid) add(r sensor.Reading) {
	g.path.Add(coordinates.Cartesian{X: r.Pose.X, Y: r.Pose.Y})
	v, o := r.Analysis()
	for coord := range v {
		prob := 1 - (1 * r.Sensor.Reliability) // close to 0
		g.Adjust(coord, prob)
	}
	for coord := range o {
		prob := (1 * r.Sensor.Reliability) // close to 1
		g.Adjust(coord, prob)
	}
}

func (g Grid) bounds() (minX, minY, maxX, maxY int) {
	minX, maxX = math.MaxInt64, -math.MaxInt64
	minY, maxY = math.MaxInt64, -math.MaxInt64
	for k := range g.grid {
		if k.X < minX {
			minX = k.X
		}
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}
	log.Debugf("Grid boundaries: %d %d %d %d", minX, minY, maxX, maxY)
	return minX, minY, maxX, maxY
}
