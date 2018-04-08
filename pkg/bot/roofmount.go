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

func (b *Bot) horizontalScan(vertPos int) []RoofmountScanResult {
	b.sendReceiver.send(fmt.Sprintf(`{"command": "horizontal_scan", "vertical_position": %d, "resolution": 0.0625}`, vertPos))
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
		Heading:  r0.Yaw,
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

// LidarScan spins the lidar for one rotation.
func (b *Bot) LidarScan(verticalPos int) []coordinates.Point {
	log.Info("Lidar Scan...")
	ps := make([]coordinates.Point, 0)
	orientation := b.orientation()
	origin := coordinates.Vector{X: b.pose.Location.X, Y: b.pose.Location.Y}
	scans := b.horizontalScan(verticalPos)
	for _, scan := range scans {
		p := coordinates.Point{
			Origin:      origin,
			Orientation: orientation,
			Vector:      coordinates.NewVector(scan.HorizontalPosition, scan.VerticalPosition, scan.Lidar),
		}
		if b.pointPub != nil {
			b.pointPub.Publish(p)
		}
		ps = append(ps, p)
	}
	return ps
}
