// Package perception uses sensor data to create an understanding the environment around the robot as well as the robot's position in that environment.
package perception

// Perception of the vehicle given sensor data.
type Perception struct {
	EnvironmentModel EnvironmentModel `json:"environmentModel"`
	DrivableSurface  Surface          `json:"drivableSurface"`
	VehiclePose      VehiclePose      `json:"vehiclePose"`
}

func New() *Perception {
	return &Perception{}
}
