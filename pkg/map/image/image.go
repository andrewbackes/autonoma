package image

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"image/color"
	"image/png"
	"os"

	"github.com/andrewbackes/autonoma/pkg/set"
)

// Occupied generates a set of all occupied squares in a black and white image.
// Black represents occupied and white represents vacant. Keys are strings of
// coordinates seperated by a comma without a space. The middle of the image is
// 0,0
//
// Examples:
//		1,1
//		1337,1
//		0,999
func Occupied(filepath string) (set.Set, error) {
	s := set.New()

	in, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	img, err := png.Decode(in)
	if err != nil {
		return nil, err
	}

	centerX := img.Bounds().Max.X / 2
	centerY := img.Bounds().Max.Y / 2
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			if img.At(x, y) != img.ColorModel().Convert(color.White) {
				key := fmt.Sprintf(`(%d,%d)`, x-centerX, y-centerY)
				s.Add(key)
				log.Debug(fmt.Sprintf(`Added %s as occupied`, key))
			}
		}
	}

	return s, nil
}
