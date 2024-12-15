package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// クロージャーを使うと、関数内に変数を閉じ込めて状態を隔離できる
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(pos(i), neg(-i))
	}
}
