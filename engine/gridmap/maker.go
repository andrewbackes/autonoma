package gridmap

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
)

// Maker is an interface for creating maps.
type Maker interface {
	Mark(x, y int, s sensor.Measurement) error
}
