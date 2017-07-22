package occgrid

import (
	"image/jpeg"
	"os"
)

// ExampleImage attempts to create an image from an occgrid.
func ExampleImage() {
	g := NewDefaultGrid(1000, 1000)
	f, err := os.Create("occgrid.jpeg")
	if err != nil {
		panic("Could not create file.")
	}
	err = jpeg.Encode(f, g, nil)
	if err != nil {
		panic("Could encode jpeg")
	}
}
