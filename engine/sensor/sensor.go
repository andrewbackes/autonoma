package sensor

import (
	"encoding/json"
	"log"
)

// Sensor represents a physical sensor on the bot.
type Sensor struct {
	// Id is the sensor's unique identifier.
	ID string `json:"id"`

	// MaxDistance is the maximum range of the sensor.
	MaxDistance float64 `json:"maxDistance"`
	// ConeWidth is how wide the sensor is in degrees.
	ConeWidth float64 `json:"coneWidth"`
	// Inclusive determines how to handle a measurement. When 'true',
	// everything within the range of the sensor will be marked as an obstruction.
	Inclusive bool `json:"inclusive"`

	// AngleOffset is the direction the sensor is mounted.
	AngleOffset float64
	// Xoffset is how far left or right the sensor is from the center of mass.
	XOffset int `json:"xOffset"`
	// Xoffset is how far up or down the sensor is from the center of mass.
	YOffset int `json:"yOffset"`
}

func DecodeSensor(payload []byte) *Sensor {
	s := &Sensor{}
	if err := json.Unmarshal(payload, &s); err != nil {
		log.Println("Could not decode", string(payload))
	}
	return s
}
