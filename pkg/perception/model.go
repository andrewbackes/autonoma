package perception

import (
	"github.com/andrewbackes/autonoma/pkg/pointcloud"
)

type EnvironmentModel struct {
	PointCloud *pointcloud.PointCloud `json:"pointCloud"`
}
