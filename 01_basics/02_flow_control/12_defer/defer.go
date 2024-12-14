package main

import "fmt"

// deferステートメントは、呼び出した関数がreturnされる直前に実行させます。
func main() {
	defer fmt.Println("piyo")
	fmt.Println("hoge")
	//hoge piyo
}
