package main

import (
	"fmt"
	"image"
)

/*
Image インターフェース

// Image is a finite rectangular grid of [color.Color] values taken from a color
// model.
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}

color.Color と color.Model は共にインタフェースですが、
定義済みの color.RGBA と color.RGBAModel を使うことで、このインタフェースを無視できます。
これらのインタフェースは、image/color パッケージで定義されています。
*/

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
