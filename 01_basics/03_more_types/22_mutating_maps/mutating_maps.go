package main

import "fmt"

func main() {
	// makeで作成
	m := make(map[string]int)

	// 値の更新はキーを指定
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 削除にはdelete()関数を使用します
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// ok は存在していればtrue 存在しない場合はfalse
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

