package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (m Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)           // 色を計算 (例: (x + y) / 2)
	return color.RGBA{v, v, 255, 255} // RGBA 色を返す
}

func Pic(dx, dy int) [][]uint8 {
	// dy の長さのスライスを作成
	result := make([][]uint8, dy)

	for y := range dy {
		// 各 y に対して dx の長さのスライスを割り当て
		result[y] = make([]uint8, dx)
		for x := range dx {
			// ピクセル値を計算 (ここでは (x+y)/2)
			result[y][x] = uint8((x * y))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
	m := Image{}
	pic.ShowImage(m)
}
