package main

import "fmt"

func main() {
	// スライスリテラルは長さのない配列宣言
	// スライスを宣言すると、配列が生成されてそれに対する参照を得たスライスを生成する
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
