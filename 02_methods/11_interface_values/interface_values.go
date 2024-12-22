package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	// interface i は内部的に (動的型, 値) を保持します。
	// 例えば、i = &T{"Hello"} の場合、内部的には (*T, &T{S: "Hello"}) という情報を持ちます。
	// i.M() が呼び出されると、動的型 (*T) の M メソッドが実行されます。
	// 次に、i = F(math.Pi) の場合、(F, math.Pi) を持ち、F 型の M メソッドが実行されます。
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
