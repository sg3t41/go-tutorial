package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// if文はfor文と同様に()は不要で、{}は必要です。
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
