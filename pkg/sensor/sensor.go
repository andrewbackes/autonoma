package sensor

// Sensor is all of the meta data for a physical sensor.
type Sensor struct {
	// ViewAngle in degrees.
	ViewAngle float64
	// MinDistance
	MinDistance float64
	// MaxDistance is the furthest the sensor can read in cm.
	MaxDistance float64
	// A binary sensor is on when an object is within range and
	// off otherwise.
	Binary bool
	// Reliability of the sensor from 0 to 1.
	Reliability float64
}

var (
	// UltraSonic represents a physical UltraSonic sensor.
	UltraSonic = Sensor{
		ViewAngle:   15.0,
		MaxDistance: 18.0,
		Binary:      false,
		Reliability: 1,
	}
)
