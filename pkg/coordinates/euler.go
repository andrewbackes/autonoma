package coordinates

// Euler represents the orientation of an object in 3-space. It is represented by 3 rotations.
type Euler struct {
	// Yaw is the heading.
	Yaw float64 `json:"yaw"`
	// Pitch is the forward/backward angle.
	Pitch float64 `json:"pitch"`
	// Roll is the left/right angle.
	Roll float64 `json:"roll"`
}
