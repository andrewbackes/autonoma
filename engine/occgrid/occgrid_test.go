package occgrid

import (
	//"github.com/andrewbackes/autonoma/engine/sensor"
	"github.com/stretchr/testify/assert"
	"image/jpeg"
	"os"
	"testing"
)

func TestIndexConversionSmall(t *testing.T) {
	/*
		Quadrants:
		4 4 | 1 1
		4 4 | 1 1
		- - - - -
		3 3 | 2 2
		3 3 | 2 2

		Indices:
		0  1  2  3  4
		5  6  7  8  9
		10 11 12 13 14
		15 16 17 18 19
		20 21 22 23 24

		Note: Don't forget the axis
	*/
	var assert = assert.New(t)
	g := NewDefaultGrid(5, 5)
	assert.Equal(12, g.index(0, 0))

	assert.Equal(8, g.index(1, 1))
	assert.Equal(16, g.index(-1, -1))
	assert.Equal(6, g.index(-1, 1))
	assert.Equal(18, g.index(1, -1))

	assert.Equal(4, g.index(2, 2))
	assert.Equal(0, g.index(-2, 2))
	assert.Equal(24, g.index(2, -2))
	assert.Equal(20, g.index(-2, -2))
}

/*
func TestExampleFullCircle(t *testing.T) {
	g := NewGrid(500, 500, 2)
	dist := float64(249)
	for d := dist; d > 0; d-- {
		m := sensor.Measurement{Distance: d, Angle: 0, ConeSize: 360}
		g.Mark(sensor.Origin{}, m)
		g.Mark(sensor.Origin{}, m)
	}
	save(g, "grid_fullcircle.jpeg")
}
*/

/*
func TestExampleImage(t *testing.T) {
	g := NewDefaultGrid(1000, 1000)
	save(g, "test_output/grid_empty.jpeg")
}



func TestExampleHollowCircle(t *testing.T) {
	g := NewGrid(500, 500, 10)
	dist := float64(249)
	m := sensor.Measurement{Distance: dist, Angle: 0, ConeSize: 360}
	g.Mark(sensor.Origin{}, m)
	save(g, "test_output/grid_hollowcircle.jpeg")
}

func TestExampleLine(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(50)
	for d := dist; d >= 1; d-- {
		pos := sensor.Measurement{Distance: d, Angle: 45}
		g.Mark(sensor.Origin{}, pos)
		neg := sensor.Measurement{Distance: d, Angle: 225}
		g.Mark(sensor.Origin{}, neg)
	}
	save(g, "test_output/grid_line.jpeg")
}

func TestExampleRay(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(50)
	pos := sensor.Measurement{Distance: dist, Angle: 45}
	g.Mark(sensor.Origin{}, pos)
	neg := sensor.Measurement{Distance: dist, Angle: 225}
	g.Mark(sensor.Origin{}, neg)
	save(g, "test_output/grid_ray.jpeg")
}

func TestExampleConeRight(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(49)
	pos := sensor.Measurement{Distance: dist, Angle: 90, ConeSize: 30}
	g.Mark(sensor.Origin{}, pos)
	save(g, "test_output/grid_cone_right.jpeg")
}

func TestExampleConeLeft(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(49)
	pos := sensor.Measurement{Distance: dist, Angle: -90, ConeSize: 30}
	g.Mark(sensor.Origin{}, pos)
	save(g, "test_output/grid_cone_left.jpeg")
}

func TestExampleConeUp(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(49)
	pos := sensor.Measurement{Distance: dist, Angle: 0, ConeSize: 30}
	g.Mark(sensor.Origin{}, pos)
	save(g, "test_output/grid_cone_up.jpeg")
}

func TestExampleConeDown(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(49)
	pos := sensor.Measurement{Distance: dist, Angle: 180, ConeSize: 30}
	g.Mark(sensor.Origin{}, pos)
	save(g, "test_output/grid_cone_down.jpeg")
}

func TestExampleOffsetUp(t *testing.T) {
	g := NewGrid(200, 200, 2)
	dist := float64(49)
	m := sensor.Measurement{Distance: dist, Angle: 0, ConeSize: 30}
	g.Mark(sensor.Origin{Xoffset: -25}, m)
	g.Mark(sensor.Origin{Xoffset: 25}, m)
	save(g, "test_output/grid_offset_up.jpeg")
}

func TestExampleOffset45(t *testing.T) {
	g := NewGrid(200, 200, 2)
	dist := float64(49)
	m := sensor.Measurement{Distance: dist, Angle: 45, ConeSize: 30}
	g.Mark(sensor.Origin{Xoffset: -25}, m)
	g.Mark(sensor.Origin{Xoffset: 25}, m)
	save(g, "test_output/grid_offset_45.jpeg")
}

func TestExampleOffsetRight(t *testing.T) {
	g := NewGrid(200, 200, 2)
	dist := float64(49)
	m := sensor.Measurement{Distance: dist, Angle: 90, ConeSize: 30}
	g.Mark(sensor.Origin{Xoffset: -25}, m)
	g.Mark(sensor.Origin{Xoffset: 25}, m)
	save(g, "test_output/grid_offset_right.jpeg")
}

func TestExampleMultiOrigin(t *testing.T) {
	g := NewGrid(100, 100, 2)
	dist := float64(20)
	g.Mark(sensor.Origin{Y: 20}, sensor.Measurement{Distance: dist, Angle: 90, ConeSize: 30})
	g.Mark(sensor.Origin{X: 20}, sensor.Measurement{Distance: dist, Angle: 180, ConeSize: 30})
	g.Mark(sensor.Origin{Y: -20}, sensor.Measurement{Distance: dist, Angle: 270, ConeSize: 30})
	g.Mark(sensor.Origin{X: -20}, sensor.Measurement{Distance: dist, Angle: 0, ConeSize: 30})

	save(g, "test_output/grid_multi_origin.jpeg")
}

*/

func save(g *Grid, filename string) {
	f, _ := os.Create(filename)
	err := jpeg.Encode(f, g, nil)
	if err != nil {
		panic("Could encode jpeg")
	}
}
