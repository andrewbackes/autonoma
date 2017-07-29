package actions

import (
	"fmt"
)

func Move(dist, x, y float64) string {
	r := fmt.Sprintf(`{ "action": "move", "distance": %f, "destination": {"x": %f, "y": %f }}`, dist, x, y)
	return r
}

func Rotate(heading float64) string {
	r := fmt.Sprintf(`{ "action": "rotate", "heading": %f }`, heading)
	return r
}

func Look(degrees float64) string {
	r := fmt.Sprintf(`{ "action": "look", "degrees": %f }`, degrees)
	return r
}

func Read(sensor_id string) string {
	r := fmt.Sprintf(`{ "action": "read", "sensor_id": "%s" }`, sensor_id)
	return r
}
