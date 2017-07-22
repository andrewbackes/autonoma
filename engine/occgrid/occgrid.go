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
	occ    []uint8
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
		occ:            make([]uint8, height*width),
		path:           make([]bool, height*width),
		height:         height,
		width:          width,
		pos:            point{width / 2, height / 2},
		colorModel:     color.RGBAModel,
		colors:         make(map[uint8]color.Color),
		pathColor:      color.RGBA{R: 255, G: 0, B: 0, A: 255},
		maxProbability: uint8(maxProbability),
	}
	for i := 0; i < len(g.occ); i++ {
		g.occ[i] = g.maxProbability / 2
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
	return image.Rect(0, 0, g.width, g.height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (g *Grid) At(x, y int) color.Color {
	if g.path[g.index(x, y)] {
		return g.pathColor
	}
	return g.colors[g.occ[g.index(x, y)]]
}

// Center returns the coordates of the center of the Grid.
func (g *Grid) Center() (x, y int) {
	return g.width / 2, g.height / 2
}

func (g *Grid) Move(x, y int) {
	// TODO
}

// index converts coordinates to an array index.
func (g *Grid) index(x, y int) int {
	return y*g.width + x
}

// Mark will adjust the probability of an obstruction on the grid.
func (g *Grid) Mark(o sensor.Origin, m sensor.Measurement) error {
	// TODO(andrewbackes): offsets.
	marked := newPointset()
	startAngle := math.Mod(m.Angle-m.ConeSize/2, 360)
	endAngle := math.Mod(m.Angle+m.ConeSize/2, 360)
	for a := startAngle; a <= endAngle; a += 0.25 {
		em := sensor.Measurement{
			Distance: m.Distance,
			Angle:    a,
		}
		ep := g.point(o, em)
		if !marked.contains(ep) {
			g.increase(ep)
			marked.add(ep)
		}
		for d := float64(0); d < m.Distance; d++ {
			im := sensor.Measurement{
				Distance: d,
				Angle:    a,
			}
			ip := g.point(o, im)
			if !marked.contains(ip) {
				g.decrease(ip)
				marked.add(ip)
			}
		}
	}
	return nil
}

func (g *Grid) point(o sensor.Origin, m sensor.Measurement) point {
	angle := math.Mod(m.Angle+o.Heading-90, 360)
	destX, destY := polarToCart(m.Distance, angle)
	centerX, centerY := g.Center()
	destX += centerX + o.X
	destY += centerY - o.Y
	return point{x: destX, y: destY}
}

func polarToCart(dist, angle float64) (x, y int) {
	return int(dist * math.Cos(toRadians(angle))), int(dist * math.Sin(toRadians(angle)))
}

func toRadians(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func (g *Grid) increase(p point) {
	if g.occ[g.index(p.x, p.y)] != g.maxProbability {
		g.occ[g.index(p.x, p.y)] = g.occ[g.index(p.x, p.y)] + 1
	}
}

func (g *Grid) decrease(p point) {
	if g.occ[g.index(p.x, p.y)] != 0 {
		g.occ[g.index(p.x, p.y)] = g.occ[g.index(p.x, p.y)] - 1
	}
}
