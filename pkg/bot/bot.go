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

type pointPub interface {
	Publish(coordinates.Point)
}

type Bot struct {
	pose         coordinates.Pose
	sensors      map[string]sensor.Sensor
	dimensions   Dimensions
	wheels       Wheels
	sendReceiver sendReceiver

	pointPub      pointPub
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
func New(address string, sensors map[string]sensor.Sensor, d Dimensions, w Wheels, p pointPub) *Bot {
	b := &Bot{
		sendReceiver:  &tcpSendReceiver{address: address},
		sensors:       sensors,
		dimensions:    d,
		wheels:        w,
		rotationError: 10,
		pointPub:      p,
	}
	return b
}

func (b *Bot) Rotate(heading float64) {
	log.Info("Rotating to ", heading)
	b.rotate(heading, 150)
}

func (b *Bot) rotate(target, turnRate float64) {
	start := b.Pose().Heading
	dist, dir := smallestAngle(start, target)
	t := math.Max(dist/turnRate, 0.1)
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "%s", "time": %f, "speed": %d}`, dir, t, b.wheels.MaxPower))
	end := b.Pose().Heading
	distFromTarget, _ := smallestAngle(end, target)
	if distFromTarget > b.rotationError {
		//travelDist, _ := smallestAngle(start, end)
		//newTurnRate := travelDist / t
		newTurnRate := 150.0
		b.rotate(target, newTurnRate)
	}
}

func smallestAngle(from, to float64) (dist float64, dir string) {
	var cwd, ccwd float64
	if from < to {
		cwd = math.Abs(to - from)
		ccwd = math.Abs((360 - to) + from)
	} else {
		ccwd = math.Abs(to - from)
		cwd = math.Abs((360 - to) + from)
	}
	dir = "clockwise"
	dist = cwd
	if ccwd < cwd {
		dir = "counter_clockwise"
		dist = ccwd
	}
	log.Debug("Smallest angle: ", dist)
	return
}

/*
Move the bot a distance.

Experiments:
	1sec@20% = 7 7/8 in  = 20.0025 cm  => 20    cm/sec
	2sec@20% = 19 5/8 in = 49.8475 cm  => 25    cm/sec
	3sec@20% = 32 in     = 81.28 cm    => 27.09 cm/sec
*/
func (b *Bot) Move(d distance.Distance) {
	p := b.Pose()
	circ := b.wheels.Diameter * math.Pi
	rot := d / circ
	t := float64(rot) / float64(b.wheels.RPM*60)
	log.Info("Moving ", "forward", " for ", t, " seconds")
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "forward", "time": %f, "speed": %d}`, t, b.wheels.MaxPower))
	time.Sleep(time.Duration(t*1000) * time.Millisecond)
	b.pose.Location = coordinates.Add(p.Location, coordinates.CompassRose{Heading: p.Heading, Distance: d})
}

func (b *Bot) readings() map[string]float64 {
	b.sendReceiver.send(`{"command": "get_readings"}`)
	time.Sleep(20 * time.Millisecond)
	resp, _ := b.sendReceiver.receive()
	readings := map[string]float64{}
	err := json.Unmarshal([]byte(resp), &readings)
	if err != nil {
		panic(err)
	}
	return readings
}

func (b *Bot) orientation() coordinates.Euler {
	b.sendReceiver.send(`{"command": "get_orientation"}`)
	resp, _ := b.sendReceiver.receive()
	euler := coordinates.Euler{}
	err := json.Unmarshal([]byte(resp), &euler)
	if err != nil {
		panic(err)
	}
	return euler
}

func (b *Bot) Pose() coordinates.Pose {
	b.pose.Heading = b.readings()["heading"]
	return b.pose
}

func (b *Bot) SetPose(p coordinates.Pose) {
	b.pose = p
}
