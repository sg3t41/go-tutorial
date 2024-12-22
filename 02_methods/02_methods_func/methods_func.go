package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// この場合はレシーバ引数を取らない、通常の"関数"
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
