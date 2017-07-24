package gridmap

import (
	"image"
)

// Reader can display a map.
type Reader interface {
	image.Image
	Center() (x, y int)
}
