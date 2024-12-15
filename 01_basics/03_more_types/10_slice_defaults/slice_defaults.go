package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// [0:2]と同じ
	s = s[:2]
	fmt.Println(s)

	// [1:最後まで]の意
	s = s[1:]
	fmt.Println(s)
}
