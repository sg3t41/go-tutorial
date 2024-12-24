package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	ch := make(chan int)     // チャネルを作成
	go Walk(tree.New(1), ch) // 新しい木を作成してWalkを実行

	// チャネルから10個の値を受け取り表示
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}

// Walkは、与えられた二分木tを歩き、すべての値をチャネルchに送信します。
func Walk(t *tree.Tree, ch chan int) {
	// 二分木を深さ優先で走査する再帰関数
	if t == nil {
		return
	}
	Walk(t.Left, ch)  // 左の子ノード
	ch <- t.Value     // 現在のノードの値を送信
	Walk(t.Right, ch) // 右の子ノード
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 並行して2つの木をウォーク
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// 両方のチャネルから値を受け取って比較
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}
