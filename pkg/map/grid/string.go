package grid

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

func (g Grid) String() string {
	minX, minY, maxX, maxY := g.bounds()
	s := ""
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			odds := g.Get(coordinates.Cartesian{X: x, Y: y})
			p := fmt.Sprintf("%.2f ", odds.Probability())
			s += p
		}
		s += "\n"
	}
	return s
}
