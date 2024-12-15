package main

import "golang.org/x/tour/pic"

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
}
