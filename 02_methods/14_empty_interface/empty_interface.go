package main

import "fmt"

func main() {
	var i interface{}
	// 空のインターフェース
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
	// (<nil>, <nil>)
	// (42, int)
	// (hello, string)
}
