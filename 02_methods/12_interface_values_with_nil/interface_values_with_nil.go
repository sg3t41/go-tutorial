package main

import "github.com/sg3t41/go-tutorial/02_methods/12_interface_values_with_nil/t"

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Tがpublicの場合はnilポインターの*Tを使ってMを呼び出し得るので養生が必要
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	// 構造体を隠蔽化してNewを作ればnilチェック不要
	tt := tp.New()
	tt.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
