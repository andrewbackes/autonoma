package sensor

import (
	"time"
)

// Reading from a sensor.
type Reading struct {
	// TimeStamp when reading was taken.
	TimeStamp time.Time

	// Value returned by sensor.
	Value float64

	// Sensor that generated the reading.
	Sensor Sensor

	// Pose is the orientation of the sensor.
	Pose Pose
}

func (r Reading) Distance() *float64 {
	if r.Value >= r.Sensor.MaxDistance {
		return nil
	}
	return &r.Value
}
