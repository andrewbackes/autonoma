package perception

import (
	"github.com/andrewbackes/autonoma/pkg/vector"
)

type Vehicle struct {
	Location vector.Vector `json:"location"`
	Odometer float64       `json:"odometer"`
}
