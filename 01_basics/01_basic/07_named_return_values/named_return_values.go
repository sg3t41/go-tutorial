//https://go-tour-jp.appspot.com/basics/7

package main

import "fmt"

func split(sum int) (x, y int) {
	// int型の引数sumの値を4:3の比率に分けます。
	x = sum * 4 / 7
	y = sum - x
	return
}

func main() {
	fmt.Println(split(70))
}
