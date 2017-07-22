// Package sensor provides models for representing sensors.
package sensor

// Measurement is a reading from a sensor.
type Measurement struct {
	// Distance is how far the obstruction is from the sensor.
	Distance float64
	// ConeAngle is how wide in degrees the sensor reading is.
	ConeSize float64
	// Angle is the orientation in degrees from forward facing.
	Angle float64
}
