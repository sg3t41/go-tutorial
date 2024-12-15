package main

import "fmt"

func main() {
	var s []int
	// スライスへの要素の追加は組み込みのappend()を使います
	// 元の配列のキャパを越えた場合は新しい配列を作り直して参照を返すから、元配列への影響はなくなる
	// 詳細: https://go.dev/blog/slices-intro
	s = append(s, 1, 2, 3)

	fmt.Println(s) // [1,2,3]
}
