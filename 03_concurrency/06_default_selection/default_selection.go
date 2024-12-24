package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return

		// 期待する値が受信する以外のときの処理＝耐機中の処理
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
