package slam

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
)

// Bot is for controlling bots.
type Bot interface {
	Rotate(heading float64)
	Move(d distance.Distance)

	Pose() coordinates.Pose
	SetPose(coordinates.Pose)

	LidarScan(verticalPos int) []coordinates.Point
}
