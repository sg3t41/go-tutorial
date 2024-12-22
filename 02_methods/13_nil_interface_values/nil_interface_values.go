package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// nilインターフェースは型も値も保持しないのでエラー
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
