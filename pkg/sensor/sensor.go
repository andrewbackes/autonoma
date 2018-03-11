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
}

var (
	// UltraSonic represents a physical UltraSonic sensor.
	UltraSonic = Sensor{
		ViewAngle:   15.0,
		MaxDistance: 18.0,
		Binary:      false,
	}

	IRDistance = Sensor{
		ViewAngle:   0.0,
		MaxDistance: 80.0,
		MinDistance: 10.0,
		Binary:      false,
	}
)
