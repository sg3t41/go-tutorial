package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex = Vertex{1, 2}
	var p *Vertex = &v

	// 構造体のフィールドには、構造体のポインターを通してアクセスすることも可能です。
	(*p).X = 1e9
	fmt.Println(v) // {1000000000 2}

	// (*p).X はp.Xと書いてもコンパイラによってよしなに自動変換されます
	p.X = 43
	fmt.Println(v) // {43 2}
}
