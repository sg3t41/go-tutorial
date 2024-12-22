package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// String()メソッドで文字列表示用のフォーマットを作ればfmt.Stringer インターフェースを実装できる
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// StringerインターフェースとしてString()が実装されていればそれが呼ばれる
	fmt.Println(a, z)
}
