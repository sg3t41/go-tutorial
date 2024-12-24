package main

import (
	"fmt"
)

func fibonacchi(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// これ以上値を送信しない場合はcloseして受け手に通知する
	close(c) // チャネルをクローズして終了を通知
}

func main() {
	c := make(chan int, 20)
	go fibonacchi(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

