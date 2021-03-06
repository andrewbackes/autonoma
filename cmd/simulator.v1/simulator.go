package main

import (
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"
	"image/png"
	"os"
	"path/filepath"

	"github.com/andrewbackes/autonoma/pkg/bot/simulator"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/hud"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"github.com/andrewbackes/autonoma/pkg/map/image"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/sensor/simulate"
	"github.com/andrewbackes/autonoma/pkg/slam"
)

var (
	mapPath          = "pkg/map/image/assets/maze1.png"
	poseSpacing      = 10
	poseAngleSpacing = 5.0
	gridCellSize     = 5 * distance.Centimeter
	sensors          = []sensor.Sensor{sensor.SharpGP2Y0A21YK0F}
	// sensors  = []sensor.Sensor{sensor.UltraSonicHCSR04}
	logLevel = log.InfoLevel
)

func main() {
	log.SetLevel(logLevel)
	fixedReadingsSimulator()
	mappingSimulator()
}

func mappingSimulator() {
	// mapName := filepath.Base(mapPath)
	bot := simulator.New(mapPath, sensors...)
	grid := grid.New(gridCellSize)
	go slam.RandomlyMap(grid, bot)
	hud.ListenAndServe(grid)
}

func fixedReadingsSimulator() {
	mapName := filepath.Base(mapPath)
	occ, err := image.Occupied(mapPath)
	if err != nil {
		panic(err)
	}
	poses := makePoses(mapPath)
	readings := simulateSensorReadings(poses, occ)

	// Create occupancy grid based on sensor readings
	log.Info("Creating occupancy grid from sensor readings.")
	g := grid.New(gridCellSize)
	g.Apply(readings...)

	os.MkdirAll("output", os.ModePerm)
	// Output image of results
	log.Info("Saving image.")
	img, err := os.Create(fmt.Sprintf("output/%s-map.png", mapName))
	check(err)
	err = png.Encode(img, (*grid.Image)(g))
	img.Close()
	check(err)

	/*
		// Output text file of results
		log.Info("Saving text file.")
		txt, err := os.Create(fmt.Sprintf("output/%s-probabilities.txt", mapName))
		check(err)
		_, err = txt.WriteString(g.String())
		txt.Close()
		check(err)
	*/
	log.Info("Simulator Ended.")
}

func makePoses(mapFilePath string) []coordinates.Pose {
	log.Info("Generating poses around map.")
	mapMaxX, mapMaxY := image.Bounds(mapFilePath)
	poses := simulate.Poses(mapMaxX/2, mapMaxY/2, poseSpacing, poseAngleSpacing)
	log.Debugf("Generated %d poses.", len(poses))
	return poses
}

func simulateSensorReadings(poses []coordinates.Pose, occ coordinates.CartesianSet) []sensor.Reading {
	log.Info("Simulating sensor readings.")
	readings := make([]sensor.Reading, 0, len(poses))
	for _, pose := range poses {
		for _, s := range sensors {
			r := simulate.Reading(s, pose, occ)
			readings = append(readings, r)
		}
	}
	log.Debugf("Simulated %d readings.", len(readings))
	return readings
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
