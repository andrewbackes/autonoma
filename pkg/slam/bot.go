package slam

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

// Bot is for controlling bots.
type bot interface {
	Rotate(heading float64)
	Move(d distance.Distance)

	Pose() coordinates.Pose
	SetPose(coordinates.Pose)

	Scan() []sensor.Reading
	LidarScan(verticalPos int) []coordinates.Point
}
