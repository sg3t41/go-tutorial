package main

import "fmt"

func main() {
	// 固定長配列[n]Tに対して、[]Tで可変長配列"スライス"を生成できます
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice[1:4] は、sliceの要素のうちの1から3番目を含むスライスを作ります
	var s []int = primes[1:4]
	fmt.Println(s)
}
