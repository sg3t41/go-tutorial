package main

import (
	"fmt"
	"math"
)

// 関数も変数として扱えます
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 7))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
