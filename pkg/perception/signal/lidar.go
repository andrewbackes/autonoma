package signal

import (
	"github.com/andrewbackes/autonoma/pkg/vector"
)

type LidarScan struct {
	Orientation Orientation     `json:"orientation"`
	Odometer    float64         `json:"odometer"`
	Origin      vector.Vector   `json:"origin"`
	Vectors     []vector.Vector `json:"vectors"`
}

type Orientation struct {
	Yaw   float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll  float64 `json:"roll"`
}
