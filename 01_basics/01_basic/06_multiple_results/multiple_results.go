// https://go-tour-jp.appspot.com/basics/6

package main

import "fmt"

func swap(x, y string) (string, string) {
	// 関数は複数の値をreutrnすることができます。
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello
}
