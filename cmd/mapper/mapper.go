package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"

	//"github.com/andrewbackes/autonoma/pkg/bot/simulator"
	"github.com/andrewbackes/autonoma/pkg/bot"
	"github.com/andrewbackes/autonoma/pkg/distance"
	"github.com/andrewbackes/autonoma/pkg/pointfeed"
	"github.com/andrewbackes/autonoma/pkg/pointfeed/subscribers/file"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/slam"
	"github.com/andrewbackes/autonoma/pkg/web"
)

// settings:
var (
	logLevel     = log.DebugLevel
	gridCellSize = 5 * distance.Centimeter
)

// bot:
var (
	address = "192.168.86.74:9091"
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
	d := pointfeed.New()
	//b := slamBot(d)
	//go slam.ThreeD(b)
	//go slam.TwoD(b)
	web.NewAPI(d).Start()
}

func slamBot(d *pointfeed.PointFeed) slam.Bot {
	useSim := len(os.Args) > 1 && strings.Contains(os.Args[1], "simulat")
	var b slam.Bot
	if useSim {
		var s []sensor.Sensor
		for _, sensor := range sensors {
			s = append(s, sensor)
		}
		//b = simulator.New("pkg/map/image/assets/maze1.png", s...)
	} else {
		b = bot.New(address, sensors, dimensions, wheels, d)
		go file.Subscribe(fmt.Sprintf("output/mapper-%d.json", time.Now().Unix()), d)
	}
	return b
}
