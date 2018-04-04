package bot

import (
	"encoding/json"
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/distance"
	log "github.com/sirupsen/logrus"
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
	"github.com/andrewbackes/autonoma/pkg/sensor"
)

type RoofmountScanResult struct {
	HorizontalPosition float64 `json:"horizontal_position"`
	VerticalPosition   float64 `json:"vertical_position"`
	Lidar              float64 `json:"lidar"`
}

func (b *Bot) horizontalScan(vertPos float64) []RoofmountScanResult {
	b.sendReceiver.send(fmt.Sprintf(`{"command": "horizontal_scan", "vertical_position": %f, "resolution": 1.0}`, vertPos))
	resp := b.sendReceiver.receive()
	readings := []RoofmountScanResult{}
	err := json.Unmarshal([]byte(resp), &readings)
	if err != nil {
		panic(err)
	}
	return readings
}

func (b *Bot) Scan() []sensor.Reading {
	log.Info("Scanning.")
	rs := make([]sensor.Reading, 0)
	r0 := b.orientation()
	currentPose := coordinates.Pose{
		Heading:  r0["heading"],
		Location: b.pose.Location,
	}
	scans := b.horizontalScan(0)
	for _, scan := range scans {
		h := math.Mod(currentPose.Heading+scan.HorizontalPosition, 360.0)
		r := sensor.Reading{
			Sensor: b.sensors["lidar"],
			Value:  distance.Distance(scan.Lidar),
			Pose: coordinates.Pose{
				Location: currentPose.Location,
				Heading:  float64(h),
			},
			RelativeHeading: scan.HorizontalPosition,
		}
		rs = append(rs, r)
		log.Info(r)
	}
	return rs
}
