package main

import "fmt"

// チャネルは ごルーチン間で <- を使って値の送受信をするための通り道
// 送信側がデータを送ると、受信側が受け取るまでブロッキング(待機)する
// 同様に、受信側はデータが送られてくるまでブロッキングする
// このブロッキング機能によってごルーチン間の同期がとれる
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func print(s chan string) {
	s <- "hello"
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	for range 10 {
		s = append(s, s...)
	}

	c := make(chan int)
	// 配列を分割して並行処理で作業分担を行う
	go sum(s[:len(s)/2], c)
	st := make(chan string)
	go print(st)
	ss := <-st
	fmt.Println(ss)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
