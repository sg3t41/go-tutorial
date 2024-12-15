package main

import "fmt"

type Vertex struct {
	// 緯度・経度
	Lat, Long float64
}

var m map[string]Vertex

// map = キーと値の関連付け
func main() {
	// map のゼロ値はnilです
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	fmt.Println(m["Bell"]) // {0 0}
}
