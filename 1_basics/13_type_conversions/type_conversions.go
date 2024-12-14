package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// 平方根を計算
	var f float64 = math.Sqrt(float64(x*x + y*y))

	// 型変換をする場合、 T(v)で vをTに変換できます
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
