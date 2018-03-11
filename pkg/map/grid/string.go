package grid

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

func (g *Grid) String() string {
	s := ""
	for x := g.minX; x <= g.maxX; x++ {
		for y := g.minY; y <= g.maxY; y++ {
			odds := g.Get(coordinates.Cartesian{X: x, Y: y})
			p := fmt.Sprintf("%.2f ", odds.Probability())
			s += p
		}
		s += "\n"
	}
	return s
}
