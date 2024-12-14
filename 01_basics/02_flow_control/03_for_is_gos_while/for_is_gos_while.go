package main

import "fmt"

func main() {
	sum := 1
	// ;を省略してC言語で言うwhileを実現できます
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
