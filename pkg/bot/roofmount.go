package bot

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

type RoofmountScanResult struct {
	HorizontalPosition float64 `json:"horizontal_position"`
	VerticalPosition   float64 `json:"vertical_position"`
	Lidar              float64 `json:"lidar"`
}

// LidarScan spins the lidar for one rotation.
func (b *Bot) LidarScan(verticalPos int) []coordinates.Point {
	log.Info("Lidar Scan...")
	ps := make([]coordinates.Point, 0)
	orientation := b.orientation()
	origin := coordinates.Vector{X: b.pose.Location.X, Y: b.pose.Location.Y}
	b.sendReceiver.send(fmt.Sprintf(`{"command": "horizontal_scan", "vertical_position": %d, "resolution": 0.0625}`, verticalPos))
	for {
		resp, err := b.sendReceiver.receive()
		if err != nil {
			break
		}
		if strings.Contains(resp, "complete") {
			log.Debug("Received scan complete signal.")
			break
		}
		result := RoofmountScanResult{}
		err = json.Unmarshal([]byte(resp), &result)
		if err != nil {
			log.Error(err)
		} else {
			p := coordinates.Point{
				Origin:      origin,
				Orientation: orientation,
				Vector:      coordinates.NewVector(result.HorizontalPosition, result.VerticalPosition, result.Lidar),
			}
			// --
			ps = append(ps, p)
			b.pointPub.Publish(p)
		}
	}
	return ps
}
