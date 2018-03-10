package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/andrewbackes/autonoma/pkg/map/image"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Simulator Started.")

	// Load image of map
	log.Info("Loading map from image.")
	occ, err := image.Occupied("pkg/map/image/assets/box.png")
	check(err)
	log.Debugf("There are %d occupied cells.", len(occ))

	// generate poses
	log.Info("Generating poses around map.")
	poses := sensor.SimulatePoses(250, 250, 10, 30.0)
	log.Debugf("Generated %d poses.", len(poses))

	// simulate sensor readings
	log.Info("Simulating sensor readings.")
	readings := make([]sensor.Reading, len(poses))
	for _, pose := range poses {
		r := sensor.SimulateReading(sensor.UltraSonic, pose, occ)
		readings = append(readings, r)
	}
	log.Debugf("Simulated %d readings.", len(readings))

	// create occupancy grid based on sensor readings
	log.Info("Creating occupancy grid from sensor readings.")

	// output image of results
	log.Info("Simulator Ended.")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
