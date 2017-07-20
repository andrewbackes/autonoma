// Package occgrid implements an occupancy grid. An occupancy grip implements image.Image.
package occgrid

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
	"image"
	"image/color"
)

// Grid represents a map. It represents the probability that an area is occupied vs open.
type Grid struct {

	// Grid implements image.Image
	occ    []uint8
	path   []bool
	height int
	width  int

	// Grid implements engine.mapper

}

var colorModel = color.RGBAModel

var pathColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var occColors = map[uint8]color.Color{
	0:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
	1:  color.RGBA{R: 25, G: 25, B: 25, A: 255},
	2:  color.RGBA{R: 50, G: 50, B: 50, A: 255},
	3:  color.RGBA{R: 75, G: 75, B: 75, A: 255},
	4:  color.RGBA{R: 100, G: 100, B: 100, A: 255},
	5:  color.RGBA{R: 125, G: 125, B: 125, A: 255},
	6:  color.RGBA{R: 150, G: 150, B: 150, A: 255},
	7:  color.RGBA{R: 175, G: 175, B: 175, A: 255},
	8:  color.RGBA{R: 200, G: 200, B: 200, A: 255},
	9:  color.RGBA{R: 225, G: 225, B: 225, A: 255},
	10: color.RGBA{R: 250, G: 250, B: 250, A: 255},
}

const initPixelProbability = 5
const probabilityIncrement = 1

// NewGrid returns a new Grid of the given size.
func NewGrid(height, width int) *Grid {
	o := make([]uint8, height*width)
	p := make([]bool, height*width)
	for i := 0; i < len(o); i++ {
		o[i] = initPixelProbability
	}
	return &Grid{
		occ:    o,
		path:   p,
		height: height,
		width:  width,
	}
}

// ColorModel returns the Image's color model.
func (g *Grid) ColorModel() color.Model {
	return colorModel
}

// Bounds returns the domain for which At can return non-zero color.
func (g *Grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, g.width, g.height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (g *Grid) At(x, y int) color.Color {
	if g.path[g.index(x, y)] {
		return pathColor
	}
	return occColors[g.occ[g.index(x, y)]]
}

// Center returns the coordates of the center of the Grid.
func (g *Grid) Center() (x, y int) {
	return g.width / 2, g.height / 2
}

// index converts coordinates to array index.
func (g *Grid) index(x, y int) int {
	return y*g.width + x
}

func (g *Grid) increase(x, y int) {
	if g.occ[g.index(x, y)] == 10 {
		return
	}
	g.occ[g.index(x, y)] = g.occ[g.index(x, y)] + probabilityIncrement
}

func (g *Grid) decrease(x, y int) {
	if g.occ[g.index(x, y)] == 0 {
		return
	}
	g.occ[g.index(x, y)] = g.occ[g.index(x, y)] - probabilityIncrement
}

// Mark with adjust the probability of an obstruction on the grid.
// x,y are the origin locations.
func (g *Grid) Mark(x, y int, m sensor.Measurement) error {
	return nil
}
