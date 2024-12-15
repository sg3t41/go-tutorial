package main

import "fmt"

func main() {
	// スライスは組み込み関数のmake()を使って作成することができます
	// makeはゼロ化された配列を割り当てて、その配列を参照したスライスを返します
	a := make([]int, 5)
	printSlice("a", a)

	// 長さ0, キャパ5 のスライスを作成
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
