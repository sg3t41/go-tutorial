// https://go-tour-jp.appspot.com/basics/5
package main

import "fmt"

/*
関数の引数が、同型で複数の場合は最後の型を残して省略できます。
*/
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(23, 49))
}
