package bot

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Controller interface {
	Heading(degrees float64)
	Move(d distance.Distance)

	Pose() coordinates.Pose
	SetPose(coordinates.Pose)

	Readings() []sensor.Reading
	Scan(degrees float64) []sensor.Reading
}
