package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // X:1 Y:2
	v2 = Vertex{X: 1}  // Name: 構文
	v3 = Vertex{}      // X:0  Y:0 ゼロ値になる
	p  = &Vertex{1, 2} // 初期化してポインターを返す *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

