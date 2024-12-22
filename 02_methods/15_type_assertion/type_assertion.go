package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// interface変数.(型)で、interfaceがその型を保持している場合に値を受け取れる
	s := i.(string)
	fmt.Println(s) // hello

	// 第二戻り値は、アサーションの結果としてinterfaceがその型を持っているかどうかを判定できる
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	// 第二戻り値を指定して、指定の型をインターフェースが保持しない場合は指定値のゼロ値とfalseが返ってくる
	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	// 第二戻り値を指定していないかつ保持していない型をアサーションした結果、パニックになる
	f = i.(float64) // panic
	fmt.Println(f)
}
