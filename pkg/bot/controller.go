package bot

import (
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Controller interface {
	Heading(degrees float64)
	Move(distance float64)

	Readings() []sensor.Reading
	Scan(degrees float64) []sensor.Reading
}
