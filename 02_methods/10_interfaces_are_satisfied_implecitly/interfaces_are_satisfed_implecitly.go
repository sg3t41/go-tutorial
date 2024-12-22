package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// implements などの宣言なしに、interfaceで取り決められた関数名を作成すれば、インターフェースを実装することができる
// 外部のパッケージでMを実装したinterface Tを満たす構造体を作ることもできる
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
