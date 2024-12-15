package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// indexのみ必要な場合
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// valueのみ必要な場合
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
