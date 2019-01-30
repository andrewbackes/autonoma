package grid

/*
import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

func (g *Grid) CellIsVacant(c coordinates.Vector) bool {
	min, max := g.cellMinMax(c)
	return max <= occupiedThreshold && min <= vacantThreshold
}

func (g *Grid) CellIsOccupied(c coordinates.Vector) bool {
	_, max := g.cellMinMax(c)
	return max >= occupiedThreshold
}

func (g *Grid) cellMinMax(c coordinates.Vector) (min, max float64) {
	coords := coordinates.Cell(c, g.cellSize)
	max = float64(-1)
	min = float64(2)
	coords.Range(func(coor coordinates.Vector) bool {
		prob := g.Get(coor).Probability()
		if prob > max {
			max = prob
		}
		if prob < min {
			min = prob
		}
		return false
	})
	return
}
*/
