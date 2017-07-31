// Package occgrid implements an occupancy grid. An occupancy grip implements image.Image.
package occgrid

import (
	"github.com/andrewbackes/autonoma/engine/sensor"
	"image"
	"image/color"
	"log"
)

const occThreshold = 0.55
const vacantThreshold = 0.15
const initProbability = 0.5
const occWeight = 5

// Grid represents a map. It represents the probability that an area is occupied vs open.
type Grid struct {

	// Grid implements image.Image
	probability    []uint8
	scannedCounter []int
	blockedCounter []int
	path           []bool
	position       sensor.Location
	height         int
	width          int

	colorModel color.Model
	pathColor  color.Color
	botColor   color.Color

	cellSize uint8
}

// NewGrid returns a new Grid of the given size.
func NewGrid(height, width, cellSize int) *Grid {
	g := &Grid{
		scannedCounter: make([]int, height*width),
		blockedCounter: make([]int, height*width),
		path:           make([]bool, height*width),
		position:       sensor.Location{X: 0, Y: 0},
		height:         height,
		width:          width,
		colorModel:     color.RGBAModel,
		pathColor:      color.RGBA{R: 255, G: 0, B: 0, A: 255},
		botColor:       color.RGBA{R: 0, G: 0, B: 255, A: 255},
		cellSize:       uint8(cellSize),
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
	//if g.path[g.index(x, y)] {
	//	return g.pathColor
	//}
	// Draw the bot's location:
	if (x-5 < g.position.X && g.position.X < x+5) && (y-5 < g.position.Y && g.position.Y < y+5) {
		return g.botColor
	}
	p := uint8((1 - g.cellProbability(x, y)) * 250)
	return color.RGBA{R: p, G: p, B: p, A: 255}
}

// Center returns the coordates of the center of the Grid.
func (g *Grid) Center() (x, y int) {
	return g.width / 2, g.height / 2
}

// index converts coordinates to an array index. Rounds down to the nearest cell.
func (g *Grid) index(x, y int) int {
	x2 := x + g.width/2
	y2 := -y + g.height/2 // don't forget to flip y

	return y2*g.width + x2
}

func (g *Grid) cellIndex(x, y int) int {
	x2 := x + g.width/2
	y2 := -y + g.height/2 // don't forget to flip y
	x2 = (x2 / int(g.cellSize)) * int(g.cellSize)
	y2 = (y2 / int(g.cellSize)) * int(g.cellSize)
	return y2*g.width + x2
}

// Path registers a path from the current posisiton.
func (g *Grid) Path(x, y int) {
	g.path[g.index(x, y)] = true
}

// Occupied marks a square as having an object in it.
func (g *Grid) Occupied(x, y int) {
	log.Println("Occupied: (", x, y, ")")
	g.blockedCounter[g.cellIndex(x, y)] += occWeight
	g.scannedCounter[g.cellIndex(x, y)] += occWeight
}

// Vacant marks a square as *not* having an object in it.
func (g *Grid) Vacant(x, y int) {
	g.scannedCounter[g.cellIndex(x, y)]++
}

func (g *Grid) cellProbability(x, y int) float64 {
	if g.scannedCounter[g.cellIndex(x, y)] == 0 {
		return initProbability
	}
	return float64(g.blockedCounter[g.cellIndex(x, y)]) / float64(g.scannedCounter[g.cellIndex(x, y)])
	/*
		xp := (x / int(g.cellSize)) * int(g.cellSize)
		yp := (y / int(g.cellSize)) * int(g.cellSize)
		var max float64
		for col := xp; col < xp+int(g.cellSize) && col < g.Bounds().Max.X; col++ {
			for row := yp; row < yp+int(g.cellSize) && row < g.Bounds().Max.Y; row++ {
				prob := 0.5
				if g.scannedCounter[g.index(col, row)] != 0 {
					prob = float64(g.blockedCounter[g.index(col, row)]) / float64(g.scannedCounter[g.index(col, row)])
				}
				if prob > max {
					max = prob
				}
			}
		}
		return max
	*/
}

func (g *Grid) IsOccupied(loc sensor.Location) bool {
	return g.cellProbability(loc.X, loc.Y) >= occThreshold
}

func (g *Grid) IsUnexplored(loc sensor.Location) bool {
	prob := g.cellProbability(loc.X, loc.Y)
	return prob < occThreshold && prob > vacantThreshold
}

func (g *Grid) SetPosition(x, y int) {
	g.position = sensor.Location{X: x, Y: y}
}
