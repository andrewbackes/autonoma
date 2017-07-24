package sensor

import (
	"encoding/json"
	"log"
	"math"
)

// Reading is a raw sensor reading linked to a sensor.
type Reading struct {
	// SensorId is the unique identifier of the sensor.
	SensorID string `json:"sensorId"`
	// Output is the reading of the sensor.
	Outout float64 `json:"output"`

	Heading float64 `json:"heading"`
	X       int     `json:"X"`
	Y       int     `json:"Y"`
}

// DecodeReading turns json into a struct.
func DecodeReading(payload []byte) *Reading {
	m := &Reading{}
	if err := json.Unmarshal(payload, &m); err != nil {
		log.Println("Could not decode", string(payload))
	}
	return m
}

// Process outputs occupied and vacant location sets from sensor data.
func Process(s *Sensor, r *Reading) (occupied, vacant LocationSet) {
	// TODO(andrewbackes): offsets.
	occupied = NewLocationSet()
	vacant = NewLocationSet()

	startAngle := math.Mod((r.Heading+s.AngleOffset-90)-s.ConeWidth/2, 360)
	endAngle := math.Mod((r.Heading+s.AngleOffset-90)+s.ConeWidth/2, 360)
	for a := startAngle; a <= endAngle; a += 0.25 {
		for d := float64(0); d < r.Outout; d++ {
			loc := polarToCart(d, a)
			loc.X += r.X
			loc.Y += r.Y
			if s.Inclusive {
				occupied.Add(loc)
			} else {
				vacant.Add(loc)
			}
		}
		// Endpoint:
		{
			loc := polarToCart(r.Outout, a)
			loc.X += r.X
			loc.Y += r.Y
			occupied.Add(loc)
			if vacant.Contains(loc) {
				vacant.Remove(loc)
			}
		}
	}
	log.Println("Occupied:", occupied)
	log.Println("Vacant:", vacant)
	return
}

func polarToCart(dist, angle float64) Location {
	return Location{
		X: int(dist * math.Cos(toRadians(angle))),
		Y: int(dist * math.Sin(toRadians(angle))),
	}
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}
