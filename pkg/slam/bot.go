package slam

import (
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

// Bot is for controlling bots.
type Bot interface {
	Rotate(heading float64)
	Move(d distance.Distance)

	Pose() coordinates.Pose
	SetPose(coordinates.Pose)
	Reset()

	// Readings() []sensor.Reading
	Scan() []sensor.Reading
}
