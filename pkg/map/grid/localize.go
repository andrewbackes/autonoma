package grid

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

func (g *Grid) LocalizePose(p coordinates.Pose) coordinates.Pose {
	c := p
	g.correctedPositions.Add(c.Location)
	return p
}
