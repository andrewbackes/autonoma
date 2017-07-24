package receiver

import (
	"github.com/andrewbackes/autonoma/engine/occgrid"
	"testing"
)

func TestProcessSensor(t *testing.T) {
	s := []byte(`{"type":"sensor", "id":"sensor-id"}`)
	g := occgrid.NewDefaultGrid(10, 10)
	r := New(g)
	r.process(s)
}

func TestProcessMeasurement(t *testing.T) {
	//m := []byte(`{"type":"measurement", "distance": 1.2}`)
}
