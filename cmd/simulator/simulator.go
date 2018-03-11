package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"image/png"
	"os"
	"path/filepath"

	"github.com/andrewbackes/autonoma/pkg/bot/simulator"
	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/hud"
	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"github.com/andrewbackes/autonoma/pkg/map/image"
	"github.com/andrewbackes/autonoma/pkg/mapper"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/sensor/simulate"
)

const (
	mapPath          = "pkg/map/image/assets/maze1.png"
	poseSpacing      = 10
	poseAngleSpacing = 15
)

func main() {
	// mappingSimulator()
	fixedReadingsSimulator()
}

func mappingSimulator() {
	// mapName := filepath.Base(mapPath)
	occ := getOccupied(mapPath)
	bot := simulator.New(occ, sensor.UltraSonic)
	grid := grid.New()
	go mapper.Map(&grid, bot)
	hud.ListenAndServe(&grid)
}

func fixedReadingsSimulator() {
	mapName := filepath.Base(mapPath)
	occ := getOccupied(mapPath)
	poses := makePoses(mapPath)
	readings := simulateSensorReadings(poses, occ)

	// Create occupancy grid based on sensor readings
	log.Info("Creating occupancy grid from sensor readings.")
	g := grid.New()
	g.Apply(readings...)

	// Output text file of results
	log.Info("Saving text file.")
	os.MkdirAll("output", os.ModePerm)
	txt, err := os.Create(fmt.Sprintf("output/%s-probabilities.txt", mapName))
	check(err)
	defer txt.Close()
	_, err = txt.WriteString(g.String())
	check(err)

	// Output image of results
	log.Info("Saving image.")
	img, err := os.Create(fmt.Sprintf("output/%s-map.png", mapName))
	check(err)
	defer img.Close()
	err = png.Encode(img, grid.Image(g))
	check(err)
	log.Info("Simulator Ended.")
}

func getOccupied(mapFilePath string) coordinates.CartesianSet {
	mapName := filepath.Base(mapFilePath)
	log.Infof("Using map %s", mapName)

	log.SetLevel(log.DebugLevel)
	log.Info("Simulator Started.")

	// Load image of map
	log.Info("Loading map from image.")
	occ, err := image.Occupied(mapFilePath)
	check(err)
	log.Debugf("There are %d occupied cells.", len(occ))
	return occ
}

func makePoses(mapFilePath string) []sensor.Pose {
	log.Info("Generating poses around map.")
	mapMaxX, mapMaxY := image.Bounds(mapFilePath)
	poses := simulate.Poses(mapMaxX/2, mapMaxY/2, poseSpacing, poseAngleSpacing)
	log.Debugf("Generated %d poses.", len(poses))
	return poses
}

func simulateSensorReadings(poses []sensor.Pose, occ coordinates.CartesianSet) []sensor.Reading {
	log.Info("Simulating sensor readings.")
	readings := make([]sensor.Reading, 0, len(poses))
	for _, pose := range poses {
		r := simulate.Reading(sensor.UltraSonic, pose, occ)
		readings = append(readings, r)
	}
	log.Debugf("Simulated %d readings.", len(readings))
	return readings
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
