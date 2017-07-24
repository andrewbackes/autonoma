// Package occgrid implements an occupancy grid. An occupancy grip implements image.Image.
package occgrid

import (
	"image"
	"image/color"
)

// Grid represents a map. It represents the probability that an area is occupied vs open.
type Grid struct {

	// Grid implements image.Image
	prob   []uint8
	path   []bool
	height int
	width  int
	pos    point

	colorModel color.Model
	pathColor  color.Color
	colors     map[uint8]color.Color

	maxProbability uint8
}

// NewGrid returns a new Grid of the given size.
func NewGrid(height, width int, maxProbability int) *Grid {
	g := &Grid{
		prob:           make([]uint8, height*width),
		path:           make([]bool, height*width),
		height:         height,
		width:          width,
		pos:            point{width / 2, height / 2},
		colorModel:     color.RGBAModel,
		colors:         make(map[uint8]color.Color),
		pathColor:      color.RGBA{R: 255, G: 0, B: 0, A: 255},
		maxProbability: uint8(maxProbability),
	}
	for i := 0; i < len(g.prob); i++ {
		g.prob[i] = g.maxProbability / 2
	}
	for i := uint8(0); i <= g.maxProbability; i++ {
		c := uint8((g.maxProbability - i) * (250 / g.maxProbability))
		g.colors[i] = color.RGBA{R: c, G: c, B: c, A: 255}
	}
	return g
}

func NewDefaultGrid(height, width int) *Grid {
	return NewGrid(height, width, 10)
}

// ColorModel returns the Image's color model.
func (g *Grid) ColorModel() color.Model {
	return g.colorModel
}

// Bounds returns the domain for which At can return non-zero color.
func (g *Grid) Bounds() image.Rectangle {
	return image.Rect(-g.width/2+1, -g.height/2+1, g.width/2-1, g.height/2-1)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (g *Grid) At(x, y int) color.Color {
	//log.Println(x, y, g.index(x, y))
	if g.path[g.index(x, y)] {
		return g.pathColor
	}
	return g.colors[g.prob[g.index(x, y)]]
}

// Center returns the coordates of the center of the Grid.
func (g *Grid) Center() (x, y int) {
	return g.width / 2, g.height / 2
}

// index converts coordinates to an array index.
func (g *Grid) index(x, y int) int {
	x2 := x + g.width/2
	y2 := -y + g.height/2 // don't forget to flip y
	return y2*g.width + x2
}

// Path registers a path from the current posisiton.
func (g *Grid) Path(x, y int) {
	g.path[g.index(x, y)] = true
}

// Occupied marks a square as having an object in it.
func (g *Grid) Occupied(x, y int) {
	g.increaseProbability(x, y)
}

// Vacant marks a square as having an object in it.
func (g *Grid) Vacant(x, y int) {
	g.decreaseProbability(x, y)
}

func (g *Grid) increaseProbability(x, y int) {
	if g.prob[g.index(x, y)] != g.maxProbability {
		g.prob[g.index(x, y)] = g.prob[g.index(x, y)] + 1
	}
}

func (g *Grid) decreaseProbability(x, y int) {
	if g.prob[g.index(x, y)] != 0 {
		g.prob[g.index(x, y)] = g.prob[g.index(x, y)] - 1
	}
}
