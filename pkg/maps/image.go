package maps

import (
	"image"
	"image/color"
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

var (
	colorModel             = color.RGBAModel
	odometryPositionColor  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	correctedPositionColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}
)

// Image implements image.Image on a Grid.
type Image Map

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i *Image) At(x, y int) color.Color {
	m := (*Map)(i)
	pt := coordinates.Vector{X: x, Y: -y}

	/*
		shade := uint8(255 / 2)
		if g.CellIsVacant(pt) {
			shade = 255
		} else if g.CellIsOccupied(pt) {
			shade = 0
		}
	*/
	p := math.Min(m.grid.Get(pt).Probability(), 1)
	shade := uint8((1.0 - p) * 255.0)
	return color.RGBA{R: shade, G: shade, B: shade, A: 255}
}

// ColorModel returns the Image's color model.
func (i *Image) ColorModel() color.Model {
	return colorModel
}

// Bounds returns the domain for which At can return non-zero color.
func (i *Image) Bounds() image.Rectangle {
	return image.Rect((*Map)(i).grid.Bounds())
}
