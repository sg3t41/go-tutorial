package main

import "fmt"

// 変数に初期値を設定しないと、デフォルトでゼロ値が与えられます。
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 0 0 false ""
}
