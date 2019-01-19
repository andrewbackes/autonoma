package signal

import (
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
	"github.com/andrewbackes/autonoma/pkg/vector"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"

	"github.com/andrewbackes/autonoma/pkg/perception"
)

// UpdatePerception the signal and fit it into the perception of the world.
func UpdatePerception(s *Signal, p *perception.Perception) *perception.Perception {
	switch s.Type {
	case "lidarscan":
		var l LidarScan
		err := mapstructure.Decode(s.Event, &l)
		if err != nil {
			log.Error("Could not cast signal as LidarScan")
			return p
		}
		fitLidarscan(&l, p)
	}
	return p
}

func fitLidarscan(l *LidarScan, p *perception.Perception) *perception.Perception {
	delta := vector.PolarLikeCoordToVector(l.Orientation.Yaw, l.Odometer)
	origin := vector.Add(p.VehiclePose.Location, delta)
	for _, pt := range l.Vectors {
		p.EnvironmentModel.PointCloud.Add(pointcloud.NewPoint(origin.X+pt.X, origin.Y+pt.Y, origin.Z+pt.Z))
	}
	p.VehiclePose.Location = origin
	/*
		source := pointcloud.New()
		for _, pt := range l.Vectors {
			source.Add(pointcloud.NewPoint(origin.X+pt.X, origin.Y+pt.Y, origin.Z+pt.Z))
		}
		pointcloud.Fit(source, p.EnvironmentModel.PointCloud)
	*/
	return p
}
