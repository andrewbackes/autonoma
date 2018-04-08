package main

import (
	"github.com/andrewbackes/autonoma/pkg/bot/simulator"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/hud"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/slam"
)

// settings:
var (
	logLevel     = log.DebugLevel
	gridCellSize = 5 * distance.Centimeter
)

// bot:
var (
	address = "192.168.86.52:9091"
	sensors = map[string]sensor.Sensor{
		// key is the sensor's id sent by the bot.
		// "ir":         sensor.SharpGP2Y0A21YK0F,
		// "ultrasonic": sensor.UltraSonicHCSR04,
		"lidar": sensor.GarminLidarLiteV3,
	}
	dimensions = bot.Dimensions{
		Height: 20 * distance.Centimeter,
		Depth:  20 * distance.Centimeter,
		Width:  20 * distance.Centimeter,
	}
	wheels = bot.Wheels{
		Diameter: 6.477 * distance.Centimeter,
		RPM:      140,
		MaxPower: 50, // percentage
	}
)

func main() {
	log.SetLevel(logLevel)
	log.Info("Mapper started.")
	// b := slamBot()
	g := grid.New(gridCellSize)
	// go slam.Manual(g, b)
	// slam.ThreeD(g, b)
	hud.ListenAndServe(g)
}

func slamBot() slam.Bot {
	useSim := len(os.Args) > 1 && strings.Contains(os.Args[1], "simulat")
	var b slam.Bot
	if useSim {
		var s []sensor.Sensor
		for _, sensor := range sensors {
			s = append(s, sensor)
		}
		b = simulator.New("pkg/map/image/assets/maze1.png", s...)
	} else {
		b = bot.New(address, sensors, dimensions, wheels)
	}
	return b
}
