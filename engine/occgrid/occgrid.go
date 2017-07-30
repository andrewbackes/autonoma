// Package occgrid implements an occupancy grid. An occupancy grip implements image.Image.
package occgrid

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
	"image"
	"image/color"

	"math"
)

// Grid represents a map. It represents the probability that an area is occupied vs open.
type Grid struct {

	// Grid implements image.Image
	probability []uint8
	occupied    []bool
	path        []bool

	height int
	width  int

	colorModel color.Model
	pathColor  color.Color

	maxProbability uint8
	cellSize       uint8
}

// NewGrid returns a new Grid of the given size.
func NewGrid(height, width, maxProbability, cellSize int) *Grid {
	g := &Grid{
		probability:    make([]uint8, height*width),
		occupied:       make([]bool, height*width),
		path:           make([]bool, height*width),
		height:         height,
		width:          width,
		colorModel:     color.RGBAModel,
		pathColor:      color.RGBA{R: 0, G: 255, B: 255, A: 255},
		maxProbability: uint8(maxProbability),
		cellSize:       uint8(cellSize),
	}
	for i := 0; i < len(g.probability); i++ {
		g.probability[i] = g.maxProbability / 2
	}
	return g
}

func NewDefaultGrid(height, width int) *Grid {
	return NewGrid(height, width, 10, 5)
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
	if g.path[g.index(x, y)] {
		return g.pathColor
	}
	p := uint8((g.maxProbability - g.probability[g.index(x, y)]) * (250 / g.maxProbability))
	o := uint8(0)
	if g.occupied[g.index(x, y)] {
		o = uint8(255 / 8)
	}
	b := uint8(math.Min(float64(256), float64(p+o)))
	return color.RGBA{R: p, G: p, B: b, A: 255}
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
	if g.probability[g.index(x, y)] == g.maxProbability {
		// TODO(andrewbackes): fill in the occupancy cell.
	}
}

func (g *Grid) IsOccupied(loc sensor.Location) bool {
	return g.occupied[g.index(loc.X, loc.Y)]
}

func (g *Grid) IsUnexplored(loc sensor.Location) bool {
	occ := g.occupied[g.index(loc.X, loc.Y)] == true
	unexplored := g.probability[g.index(loc.X, loc.Y)] < (g.maxProbability/2)-1
	return !occ && unexplored
}

// Vacant marks a square as *not* having an object in it.
func (g *Grid) Vacant(x, y int) {
	g.decreaseProbability(x, y)
}

func (g *Grid) increaseProbability(x, y int) {
	if g.probability[g.index(x, y)] != g.maxProbability {
		g.probability[g.index(x, y)] = g.probability[g.index(x, y)] + 1
	}
	if g.probability[g.index(x, y)] == g.maxProbability {
		g.occupyCell(x, y)
	}
}

func (g *Grid) decreaseProbability(x, y int) {
	if g.probability[g.index(x, y)] != 0 {
		g.probability[g.index(x, y)] = g.probability[g.index(x, y)] - 1
	}
}

func (g *Grid) occupyCell(x, y int) {
	for xMin := x - int(g.cellSize/2); xMin <= x+int(g.cellSize/2); xMin++ {
		for yMin := x - int(g.cellSize/2); yMin <= x+int(g.cellSize/2); yMin++ {
			if xMin >= g.Bounds().Min.X && xMin <= g.Bounds().Max.X &&
				yMin >= g.Bounds().Min.Y && yMin <= g.Bounds().Max.Y {
				//log.Println(xMin, yMin, g.index(xMin, yMin), len(g.occupied))
				g.occupied[g.index(xMin, yMin)] = true
			}
		}
	}
}
