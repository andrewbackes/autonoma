// Package util contains utility functions.
package util

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
	"math"
)

func LocationOf(origin sensor.Location, heading float64, dist float64) sensor.Location {
	return sensor.Location{
		X: int(dist*math.Cos(toRadians(heading))) + origin.X,
		Y: int(dist*math.Sin(toRadians(heading))) + origin.Y,
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func PointsBetween(origin sensor.Location, heading, dist float64) sensor.LocationSet {
	locs := sensor.NewLocationSet()
	for d := float64(0); d < dist; d++ {
		loc := LocationOf(origin, heading, dist)
		loc.X += origin.X
		loc.Y += origin.Y
		locs.Add(loc)
	}
	return locs
}
