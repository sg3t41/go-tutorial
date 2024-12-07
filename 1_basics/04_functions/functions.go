//https://go-tour-jp.appspot.com/basics/4

package main

import "fmt"

/*
関数は0以上の引数を取ることができます。
下記のadd関数はint型の2つのパラメータを取ります。

型は変数名の後ろにつけます。
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(23, 53))
}
