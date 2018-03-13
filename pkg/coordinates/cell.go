package coordinates

import (
	"github.com/andrewbackes/autonoma/pkg/distance"
)

// Square around a coordinate.
func Square(c Cartesian, size distance.Distance) CartesianSet {
	minX := c.X - (int(size) / 2)
	minY := c.X - (int(size) / 2)
	maxX := minX + int(size)
	maxY := minY + int(size)
	s := NewCartesianSet()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			s.Add(Cartesian{X: x, Y: y})
		}
	}
	return s
}
