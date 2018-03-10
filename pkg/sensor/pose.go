package sensor

// Pose of sensor.
type Pose struct {
	// X, Y are coordinates.
	X, Y int
	// Heading is the direction the sensor is facing.
	Heading float64
}
