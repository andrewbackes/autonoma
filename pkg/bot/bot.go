// Package bot is for interacting with a physical bot over TCP.
package bot

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Bot struct {
	pose       coordinates.Pose
	sensors    map[string]sensor.Sensor
	dimensions Dimensions
	wheels     Wheels
	address    string
	conn       net.Conn
}

// Dimensions of a bot.
type Dimensions struct {
	Height, Width, Depth distance.Distance
}

// Wheels stats.
type Wheels struct {
	Diameter distance.Distance
	RPM      int
	MaxPower int
}

// New creates a bot with the specified sensors with given IP.
func New(address string, sensors map[string]sensor.Sensor, d Dimensions, w Wheels) *Bot {
	b := &Bot{
		address:    address,
		sensors:    sensors,
		dimensions: d,
		wheels:     w,
	}
	b.calibrate()
	return b
}

func (b *Bot) Rotate(heading float64) {
	b.rotate(heading)
}

func (b *Bot) rotate(heading float64) {
	h := b.Pose().Heading
	var cwd, ccwd float64
	if h < heading {
		cwd = heading - h
		ccwd = (360 - heading) + h
	} else {
		ccwd = h - heading
		cwd = (360 - h) + heading
	}
	dir := "clockwise"
	if ccwd < cwd {
		dir = "counter_clockwise"
	}
	t := 1.0
	b.send(fmt.Sprintf(`{"command": "move", "direction": "%s", "time": %f, "power": %d}`, dir, t, b.wheels.MaxPower))
}

func (b *Bot) Move(d distance.Distance) {
	p := b.Pose()
	circ := b.wheels.Diameter * math.Pi
	rot := d / circ
	t := float64(rot) / float64(b.wheels.RPM*60)
	b.send(fmt.Sprintf(`{"command": "move", "direction": "forward", "time": %f, "power": %d}`, t, b.wheels.MaxPower))
	b.pose.Location = coordinates.Add(p.Location, coordinates.CompassRose{Heading: p.Heading, Distance: d})
}

func (b *Bot) readings() map[string]float64 {
	time.Sleep(250 * time.Millisecond)
	resp := b.sendAndReceive(`{"command": "get_readings"}`)
	readings := map[string]float64{}
	err := json.Unmarshal([]byte(resp), readings)
	if err != nil {
		panic(err)
	}
	return readings
}

func (b *Bot) Pose() coordinates.Pose {
	b.pose.Heading = b.readings()["heading"]
	return b.pose
}

func (b *Bot) SetPose(p coordinates.Pose) {
	b.pose = p
}

func (b *Bot) Scan() []sensor.Reading {
	rs := make([]sensor.Reading, 0)
	r1 := b.readings()
	initPos := coordinates.Pose{
		Heading:  r1["heading"],
		Location: b.pose.Location,
	}
	rs = append(rs, sensor.Reading{
		Sensor: b.sensors["ultrasonic"],
		Value:  distance.Distance(r1["ultrasonic"]),
		Pose:   initPos,
	})
	for deg := -90; deg <= 90; deg += 5 {
		b.send(fmt.Sprintf(`{"command": "servo", "position": %d}`, deg))
		b.readings()
		rs = append(rs, sensor.Reading{
			Sensor: b.sensors["ir"],
			Value:  distance.Distance(r1["ir"]),
			Pose:   initPos,
		})
	}
	return rs
}

func (b *Bot) calibrate() {

}
