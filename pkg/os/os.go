// Package os is the operating system of the vehicle.
package os

import (
	"encoding/json"
	"fmt"
	"github.com/andrewbackes/autonoma/pkg/control"
	"github.com/andrewbackes/autonoma/pkg/perception"
	"github.com/andrewbackes/autonoma/pkg/planning"
	"github.com/andrewbackes/autonoma/pkg/sensing"
	"net/http"
)

// OperatingSystem is the stack used in the operation of an autonomous robot.
type OperatingSystem struct {
	Perceiver  perceiver
	Planner    planner
	Controller controller
}

type perceiver interface {
	Perceive(*sensing.SensorData) *perception.Perception
	Perception() *perception.Perception
}
type planner interface {
	Plan(*perception.Perception) *planning.Motions
	SetMission(*planning.Mission)
	Mission() *planning.Mission
}

type controller interface {
	Execute(*planning.Motions) *control.Commands
}

func (os *OperatingSystem) evaluate(s *sensing.SensorData) {
	os.Controller.Execute(os.Planner.Plan(os.Perceiver.Perceive(s)))
}

// Start the vehicle's operating stack.
func (os *OperatingSystem) Start() {
	// display the environment and pose
	http.HandleFunc("/perception", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(os.Perceiver.Perception())
		if err != nil {
			fmt.Printf("error - %v", err)
		}
	})
	// update the environment with sensorData
	http.HandleFunc("/perception/sensorData", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if r.Method == http.MethodPost {
			var s sensing.SensorData
			err = json.NewDecoder(r.Body).Decode(&s)
			defer r.Body.Close()
			os.evaluate(&s)
		}
		if err != nil {
			fmt.Printf("error - %v", err)
		}
	})
	// update the mission
	http.HandleFunc("/planning/mission", func(w http.ResponseWriter, r *http.Request) {
		var err error
		if r.Method == http.MethodPost {
			var m planning.Mission
			err = json.NewDecoder(r.Body).Decode(&m)
			defer r.Body.Close()
			os.Planner.SetMission(&m)
		} else {
			err = json.NewEncoder(w).Encode(os.Planner.Mission())
		}
		if err != nil {
			fmt.Printf("error - %v", err)
		}
	})
	fmt.Println("Starting...")
	http.ListenAndServe(":8080", nil)
}
