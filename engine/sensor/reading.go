// Package sensor provides models for representing sensors.
package sensor

// Measurement is a reading from a sensor.
type Measurement struct {
	// Distance is how far the obstruction is from the sensor.
	Distance int
	// Xoffset is how far left or right the sensor is from the center of mass.
	Xoffset int
	// Xoffset is how far up or down the sensor is from the center of mass.
	Yoffset int
	// ConeAngle is how wide in degrees the sensor reading is.
	ConeAngle int
	// Angle is the orientation in degrees from forward facing.
	Angle int
}
