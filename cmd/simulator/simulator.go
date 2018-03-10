package main

import (
	log "github.com/sirupsen/logrus"
	"image/png"
	"os"

	"github.com/andrewbackes/autonoma/pkg/map/grid"
	"github.com/andrewbackes/autonoma/pkg/map/image"
	"github.com/andrewbackes/autonoma/pkg/sensor"
	"github.com/andrewbackes/autonoma/pkg/sensor/simulate"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Simulator Started.")

	// Load image of map
	log.Info("Loading map from image.")
	occ, err := image.Occupied("pkg/map/image/assets/box.png")
	check(err)
	log.Debugf("There are %d occupied cells.", len(occ))

	// Generate poses
	log.Info("Generating poses around map.")
	poses := simulate.Poses(125, 125, 18, 30.0)
	log.Debugf("Generated %d poses.", len(poses))

	// Simulate sensor readings
	log.Info("Simulating sensor readings.")
	readings := make([]sensor.Reading, 0, len(poses))
	for _, pose := range poses {
		r := simulate.Reading(sensor.UltraSonic, pose, occ)
		readings = append(readings, r)
	}
	log.Debugf("Simulated %d readings.", len(readings))

	// Create occupancy grid based on sensor readings
	log.Info("Creating occupancy grid from sensor readings.")
	g := grid.New()
	g.Apply(readings...)

	// Output text file of results
	log.Info("Saving text file.")
	txt, err := os.Create("box-output.txt")
	check(err)
	defer txt.Close()
	_, err = txt.WriteString(g.String())
	check(err)

	// Output image of results
	log.Info("Saving image.")
	img, err := os.Create("box-output.png")
	check(err)
	defer img.Close()
	err = png.Encode(img, grid.Image(g))
	check(err)
	log.Info("Simulator Ended.")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
