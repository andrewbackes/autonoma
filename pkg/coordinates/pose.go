package coordinates

// Pose is a combination of location and heading.
type Pose struct {
	// X, Y are coordinates.
	Location Cartesian
	// Heading is the direction the sensor is facing.
	Heading float64
}

func NewPose(x, y int, heading float64) Pose {
	return Pose{
		Location: Cartesian{
			X: x,
			Y: y,
		},
		Heading: heading,
	}
}
