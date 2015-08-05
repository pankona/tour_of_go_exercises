package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (img *Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x % 256), uint8(y % 256), uint8(x * y % 256), 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(&m)
}
