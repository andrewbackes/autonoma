package pointcloud

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/andrewbackes/autonoma/pkg/point"
	"gonum.org/v1/gonum/mat"
)

// PointCloud is a collection of Points.
type PointCloud struct {
	Points map[point.Point]int
}

func (p *PointCloud) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`{"points":[`)
	length := len(p.Points)
	count := 0
	for key := range p.Points {
		jsonKey, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		buffer.Write(jsonKey)
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]}")
	return buffer.Bytes(), nil
}

const Dimensions = 3

func New() *PointCloud {
	return &PointCloud{
		Points: map[point.Point]int{},
	}
}

// Add a point to the cloud.
func (p *PointCloud) Add(points ...point.Point) {
	for _, t := range points {
		val := p.Points[t]
		p.Points[t] = val + 1
	}
}

func (p *PointCloud) Centroid() point.Point {
	if len(p.Points) == 0 {
		return point.Point{}
	}
	centroid := &point.Point{}
	for dim := 0; dim < 3; dim++ {
		sum := float64(0)
		for v := range p.Points {
			sum += v.Array()[dim]
		}
		centroid.SetIndex(dim, sum/float64(len(p.Points)))
	}
	return *centroid
}

func (p *PointCloud) Len() int {
	return len(p.Points)
}

/*
func (p *PointCloud) Matrix() mat.Matrix {
	if len(p.Points) == 0 {
		return &mat.Dense{}
	}
	m := mat.NewDense(p.Dimensions(), p.Len(), nil)
	for col, pt := range p.Points {
		for row, val := range pt.X {
			m.Set(row, col, float64(val))
		}
	}
	return m
}
*/

func printMatrix(m mat.Matrix) {
	r, c := m.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Print(m.At(i, j), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
