package main

import "fmt"

func main() {
	sum := 1
	// for文の初期化と後処理ステートメントは任意です
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
