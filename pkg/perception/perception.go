// Package perception uses sensor data to create an understanding the environment around the robot as well as the robot's position in that environment.
package perception

import (
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
	"github.com/andrewbackes/autonoma/pkg/vector"
)

// Perception of the vehicle given sensor data.
type Perception struct {
	EnvironmentModel EnvironmentModel  `json:"environmentModel"`
	DrivableSurface  Surface           `json:"drivableSurface"`
	Vehicle          Vehicle           `json:"vehicle"`
	Path             []vector.Vector   `json:"path"`
	Scans            [][]vector.Vector `json:"scans"`
}

func New() *Perception {
	return &Perception{
		EnvironmentModel: EnvironmentModel{
			PointCloud: pointcloud.New(),
		},
		Path:  make([]vector.Vector, 0),
		Scans: make([][]vector.Vector, 0),
	}
}
