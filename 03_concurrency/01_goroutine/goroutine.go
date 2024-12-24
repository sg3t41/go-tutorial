package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// go を関数呼び出しの前につければマルチタスク処理ができる
// goroutineは共有のメモリを使うので共通メモリへのアクセスには同期が必要です(以降のチャプターにて解説)
func main() {
	say("piyo")
	go say("goo")
	say("hoge")
	go say("foo")
	go say("world")
	say("hello")
}
