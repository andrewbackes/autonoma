package sensor

// Origin is where the sensor is located.
type Origin struct {
	// X, Y is the origin location.
	X, Y int
	// Heading is the direction relative to north.
	Heading float64
	// Xoffset is how far left or right the sensor is from the center of mass.
	Xoffset int
	// Xoffset is how far up or down the sensor is from the center of mass.
	Yoffset int
}
