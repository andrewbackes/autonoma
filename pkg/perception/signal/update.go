package signal

import (
	"fmt"
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
	angle := l.Orientation.Yaw
	fmt.Println("--> angle", angle)
	delta := vector.PolarLikeCoordToVector(angle, dist)
	fmt.Println("--> delta", delta)
	origin := vector.Add(p.Vehicle.Location, delta)
	deadReckoning := make([]vector.Vector, len(l.Vectors))
	for i, v := range l.Vectors {
		deadReckoning[i] = vector.Add(origin, vector.Rotate(v, angle))
	}
	p.Vehicle.Location = origin
	for _, v := range deadReckoning {
		p.EnvironmentModel.PointCloud.Add(v)
	}

	/*
		fitted, newOriginVector, e := fit.ICP(deadReckoning, origin, p.EnvironmentModel.PointCloud, 1.0, 10)

		fmt.Println("--> error", e)

		for _, v := range fitted {
			p.EnvironmentModel.PointCloud.Add(v)
		}
		p.Vehicle.Location = newOriginVector
	*/
	fmt.Println("--> origin", origin, "to", p.Vehicle.Location)
	p.Vehicle.Odometer = l.Odometer

	return p
}
