package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) Add() int {
	return v.X + v.Y
}

// ポインタレシーバでメソッド宣言
// **Vertexなど、ポインタのポインタはレシーバに指定できない
// レシーバ型がポインタ型の場合、そのポインタが指すインスタンスの値を直接操作ができる
// レシーバ型が値型の場合は、レシーバはその値のコピーをメソッドに渡すため、呼び出し元のインスタンスに影響する値の変更はできない
func (v *Vertex) Set(x, y int) {
	v.X = x
	v.Y = y
}

func (v Vertex) Set2(x, y int) {
	v.X = x
	v.Y = y
}

func main() {
	v := Vertex{}
	// &Vertexじゃなくても、Goはコンパイラが自動変換して(*V).Scaleにしてくれる
	v.Set(1, 2)

	// 同様に、&Vertexであっても値型レシーバを持つメソッドを呼び出す場合は自動で参照を解除する
	pv := &Vertex{}
	pv.Set(1, 2)

	fmt.Println(v.Add())
	fmt.Println(pv.Add())

	// Vertex型の値レシーバのメソッドを使うと、レシーバはコピーされるのでフィールドを変更しても呼び出し元のインスタンスには影響しない
	v2 := Vertex{}
	v2.Set2(1, 2)
	fmt.Println(v2.Add()) // 0

	// *Vertex型のポインタレシーバのメソッドを使うと、ポインタレシーバはポインタの指す実体を参照するので呼び出し元のインスタンスに変更が同期される
	v3 := &Vertex{}
	v3.Set(1, 2)
	fmt.Println(v3.Add()) //3
}
