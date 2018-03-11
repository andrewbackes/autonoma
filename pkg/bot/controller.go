package bot

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Controller interface {
	Heading(degrees float64)
	Move(distance float64)

	Pose() coordinates.Pose

	Readings() []sensor.Reading
	Scan(degrees float64) []sensor.Reading
}
