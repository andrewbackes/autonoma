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
	r0 := b.readings()
	initPos := coordinates.Pose{
		Heading:  r0["heading"],
		Location: b.pose.Location,
	}

	for deg := -90; deg <= 90; deg += 1 {
		b.sendReceiver.send(fmt.Sprintf(`{"command": "servo", "position": %d}`, deg))
		readings := b.readings()
		h := int(initPos.Heading+float64(deg)) % 360
		if h < 0 {
			h = h + 360
		}
		r := sensor.Reading{
			Sensor: b.sensors["lidar"],
			Value:  distance.Distance(readings["lidar"]),
			Pose: coordinates.Pose{
				Location: initPos.Location,
				Heading:  float64(h),
			},
			RelativeHeading: float64(deg),
		}
		rs = append(rs, r)
		log.Info(r)
	}

	return rs
}

func (b *Bot) Reset() {
	b.sendReceiver.send(`{"command": "reset"}`)
}

func (b *Bot) calibrate() {
	log.Info("Calibrating...")
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "clockwise", "time": 1, "speed": %d}`, b.wheels.MaxPower))
	time.Sleep(1500 * time.Millisecond)
	b.sendReceiver.send(fmt.Sprintf(`{"command": "move", "direction": "counter_clockwise", "time": 1, "speed": %d}`, b.wheels.MaxPower))
	time.Sleep(1500 * time.Millisecond)
}
