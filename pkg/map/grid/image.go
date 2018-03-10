package grid

import (
	"image"
	"image/color"
	"math"

	"github.com/andrewbackes/autonoma/pkg/coordinates"
)

var (
	colorModel = color.RGBAModel
	pathColor  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
)

// Image implements image.Image on a Grid.
type Image Grid

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x, y int) color.Color {

	/*
		// Draw the bot's location:
		if (x-5 < g.position.X && g.position.X < x+5) && (y-5 < g.position.Y && g.position.Y < y+5) {
			return g.botColor
		}
	*/
	if Grid(i).path.Contains(coordinates.Cartesian{X: x, Y: y}) {
		return pathColor
	}

	// p := uint8((1 - math.Min(1, g.cellProbability(x, y)/occThreshold)) * 255)
	p := math.Min(Grid(i).Get(coordinates.Cartesian{X: x, Y: y}).Probability(), 1)
	shade := uint8((1.0 - p) * 255.0)
	return color.RGBA{R: shade, G: shade, B: shade, A: 255}
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return colorModel
}

// Bounds returns the domain for which At can return non-zero color.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(Grid(i).bounds())
}
