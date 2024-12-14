package main

import "fmt"

func main() {
	var i, j int = 1, 2

	// 関数内ではvarを省略して変数宣言することができます
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
