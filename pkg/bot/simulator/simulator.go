package simulator

import (
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/sensor/simulate"
)

const (
	scanAngleIncrement = 15.0
	rotationError      = 0.0 // degrees
	movementError      = 0.0 // percent
)

type Simulator struct {
	occupied coordinates.CartesianSet
	pose     coordinates.Pose
	sensors  []sensor.Sensor
}

func New(occ coordinates.CartesianSet, sensors ...sensor.Sensor) *Simulator {
	return &Simulator{
		occupied: occ,
		pose:     coordinates.NewPose(0, 0, 0),
		sensors:  sensors,
	}
}

func (s *Simulator) Rotate(heading float64) {
	errMargin := (float64(-1 * rand.Intn(2))) * (rotationError * rand.Float64())
	s.pose.Heading = heading + errMargin
}

func (s *Simulator) Move(d distance.Distance) {
	percentError := 1 - (rand.Float64()*movementError)/100
	destVector := coordinates.CompassRose{
		Heading:  s.pose.Heading,
		Distance: distance.Distance(percentError * float64(d)),
	}.Cartesian()
	s.pose.Location.X += destVector.X
	s.pose.Location.Y += destVector.Y
	log.Debugf("Simulator moved to %s", s.pose.Location.String())
	log.Debug("Simulator Pose ", s.pose)
}

func (s *Simulator) readings() []sensor.Reading {
	rs := make([]sensor.Reading, 0, len(s.sensors))
	for _, sensor := range s.sensors {
		r := simulate.Reading(sensor, s.pose, s.occupied)
		rs = append(rs, r)
	}
	return rs
}

func (s *Simulator) Pose() coordinates.Pose {
	return s.pose
}

func (s *Simulator) SetPose(p coordinates.Pose) {
	s.pose = p
}

func (s *Simulator) Scan() []sensor.Reading {
	rs := make([]sensor.Reading, 0)
	degrees := 180.0
	startingHeading := s.pose.Heading
	for heading := startingHeading - degrees/2; heading <= startingHeading+degrees/2; heading += scanAngleIncrement {
		s.Rotate(math.Mod(heading, 360))
		rs = append(rs, s.readings()...)
	}
	return rs
}
