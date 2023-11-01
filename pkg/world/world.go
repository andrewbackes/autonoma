package world

import (
	"time"

	"github.com/andrewbackes/autonoma/pkg/pointcloud"
)

type World struct {
	Name         string                 `json:"name"`
	Created      time.Time              `json:"created"`
	Obstructions *pointcloud.PointCloud `json:"obstructions"`
}

func New() *World {
	return &World{
		Name:         "newworld",
		Created:      time.Now(),
		Obstructions: pointcloud.New(),
	}
}
