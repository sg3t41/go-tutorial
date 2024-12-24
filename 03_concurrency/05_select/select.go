package main

import "fmt"

func fibonacci(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		// selectは、goroutineの複数の通信操作で待たせる
		// select は、複数ある case のいずれかが準備できるようになるまでブロックし、準備ができた case を実行します。
		// もし、複数の case の準備ができている場合、 case はランダムに選択されます。
		select {
		case c <- x:
			x, y = y, x+y
		case q := <-quit:
			if q {
				fmt.Println("quit")
				return
			}
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibonacci(c, quit)

}
