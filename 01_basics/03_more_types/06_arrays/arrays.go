package main

import "fmt"

func main() {
	// [n]T で T型のn個の配列を生成します
	// nは固定です
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
