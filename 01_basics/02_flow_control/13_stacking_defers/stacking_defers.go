package main

import "fmt"

// deferが複数ある場合、LIFOの順で呼び出される
func main() {

	// 呼び出し元の関数単位での実行なので、LIFOといえども関数が終わる順番が早いものから呼び出される
	func() {
		defer fmt.Println("innner function")
	}()

	// LIFO: 最初に入れたものが最後に取り出される
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("counting")

	/*
		counting
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0
		done
	*/
}
