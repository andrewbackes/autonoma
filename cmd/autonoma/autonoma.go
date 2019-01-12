package main

import (
	"github.com/andrewbackes/autonoma/pkg/control"
	"github.com/andrewbackes/autonoma/pkg/os"
	"github.com/andrewbackes/autonoma/pkg/perception"
	"github.com/andrewbackes/autonoma/pkg/planning"
)

func main() {
	(&os.OperatingSystem{
		Perceiver:  perception.NewPerceiver(),
		Planner:    planning.NewPlanner(),
		Controller: control.NewController(),
	}).Start()
}
