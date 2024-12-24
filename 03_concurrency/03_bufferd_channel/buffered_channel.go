package main

import "fmt"

func main() {
	// チャネルはバッファとして使える
	// make の２つ目の引数にバッファの長さを指定できる

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1個バッファを消費
	ch <- 3
	//	ch <- 4 // バッファがいっぱいだから送信がロックされるのでエラー
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // バッファが空になって受信がロックされるのでエラー

}
