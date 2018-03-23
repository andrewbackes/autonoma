// Package bot is for interacting with a physical bot over TCP.
package bot

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"time"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type Bot struct {
	pose         coordinates.Pose
	sensors      map[string]sensor.Sensor
	dimensions   Dimensions
	wheels       Wheels
	sendReceiver sendReceiver

	rotationError float64
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
		sendReceiver:  &tcpSendReceiver{address: address},
		sensors:       sensors,
		dimensions:    d,
		wheels:        w,
		rotationError: 10,
	}
	b.calibrate()
	return b
}

func (b *Bot) Rotate(heading float64) {
	log.Info("Rotating to ", heading)
	b.rotate(heading, 150)
}

func (b *Bot) rotate(target, turnRate float64) {
	start := b.Pose().Heading
	dist, dir := smallestAngle(start, target)
	t := dist / turnRate
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "%s", "time": %f, "power": %d}`, dir, t, b.wheels.MaxPower))
	end := b.Pose().Heading
	distFromTarget, _ := smallestAngle(end, target)
	if distFromTarget > b.rotationError {
		travelDist, _ := smallestAngle(start, end)
		newTurnRate := travelDist / t
		b.rotate(target, newTurnRate)
	}
}

func smallestAngle(from, to float64) (dist float64, dir string) {
	var cwd, ccwd float64
	if from < to {
		cwd = to - from
		ccwd = (360 - to) + from
	} else {
		ccwd = to - from
		cwd = (360 - to) + from
	}
	dir = "clockwise"
	dist = cwd
	if ccwd < cwd {
		dir = "counter_clockwise"
		dist = ccwd
	}
	return
}

func (b *Bot) Move(d distance.Distance) {
	p := b.Pose()
	circ := b.wheels.Diameter * math.Pi
	rot := d / circ
	t := float64(rot) / float64(b.wheels.RPM*60)
	log.Info("Moving ", "forward", " for ", t, " seconds")
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "forward", "time": %f, "power": %d}`, t, b.wheels.MaxPower))
	b.pose.Location = coordinates.Add(p.Location, coordinates.CompassRose{Heading: p.Heading, Distance: d})
}

func (b *Bot) readings() map[string]float64 {
	time.Sleep(250 * time.Millisecond)
	b.sendReceiver.send(`{"command": "get_readings"}`)
	resp := b.sendReceiver.receive()
	readings := map[string]float64{}
	err := json.Unmarshal([]byte(resp), &readings)
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
	log.Info("Scanning.")
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
		b.sendReceiver.send(fmt.Sprintf(`{"command": "servo", "position": %d}`, deg))
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
