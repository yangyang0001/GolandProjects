package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MineImage struct{}

func main() {

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	im := MineImage{}
	pic.ShowImage(im)
}

func (im MineImage) ColorModel() color.Model {
	return nil
}

func (im MineImage) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (im MineImage) At(x, y int) color.Color {
	return nil
}
