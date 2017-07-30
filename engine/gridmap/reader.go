package gridmap

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
	"image"
)

// Reader can display a map.
type Reader interface {
	image.Image

	Center() (x, y int)
	IsOccupied(sensor.Location) bool
	IsUnexplored(loc sensor.Location) bool
}
