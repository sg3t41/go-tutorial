package main

import "fmt"

// constを使って定数宣言ができます。
const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("hello", world)
	fmt.Println(Pi)

	const Trust = true
	fmt.Println("Go Rules?", Trust)
}
