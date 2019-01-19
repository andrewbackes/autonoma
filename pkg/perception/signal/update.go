package signal

import (
	"fmt"
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
	fmt.Println("--> vehicle location", p.Vehicle.Location)
	dist := l.Odometer - p.Vehicle.Odometer
	fmt.Println("--> dist", dist)
	fmt.Println("--> yaw", l.Orientation.Yaw)
	// I don't understand why 360-yaw works =/
	delta := vector.PolarLikeCoordToVector(360-l.Orientation.Yaw, dist)
	fmt.Println("--> delta", delta)
	origin := vector.Add(p.Vehicle.Location, delta)
	fmt.Println("--> origin", origin)
	source := pointcloud.New()
	for _, pt := range l.Vectors {
		rotated := pt.Rotate(l.Orientation.Yaw)
		source.Add(pointcloud.NewPoint(origin.X+rotated.X, origin.Y+rotated.Y, origin.Z+rotated.Z))
	}
	fitted, transformation, e := pointcloud.ICP(source, p.EnvironmentModel.PointCloud, 1.0, 5)
	//fitted, transformation := source, pointcloud.NewTransformation()
	fmt.Println("--> error", e)
	fmt.Println("--> transformation", transformation)
	for _, pt := range fitted.Points {
		p.EnvironmentModel.PointCloud.Add(pt)
	}
	newOriginPt := transformation.TransformPoint(pointcloud.NewPoint(origin.X, origin.Y, origin.Z))
	p.Vehicle.Location = vector.Vector{X: newOriginPt.X[0], Y: newOriginPt.X[1], Z: newOriginPt.X[2]}
	fmt.Println("--> new origin", p.Vehicle.Location)
	p.Vehicle.Odometer = l.Odometer

	return p
}
