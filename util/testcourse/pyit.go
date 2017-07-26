package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	in, err := os.Open("courses/box_obstacles.png")
	check(err)
	out, err := os.Create("courses/box_obstacles.py")
	check(err)

	img, err := png.Decode(in)
	check(err)

	out.WriteString("#!/usr/bin/env python3\n\n")
	out.WriteString("occupied = set({\n")

	check(err)
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			if img.At(x, y) != img.ColorModel().Convert(color.White) {
				out.WriteString(fmt.Sprintf("  (%d,%d),\n", x, y))
				fmt.Println(x, y, img.At(x, y))
			}
		}
	}
	out.WriteString("})\n")
	in.Close()
	out.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
