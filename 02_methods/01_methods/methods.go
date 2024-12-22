package main

import (
	"fmt"
	"math"
)

// Goは他のOOP言語と違って、クラスがメソッドを持つというよりも関数として独立して、その関数が操作する型がある、みたいな関数ファーストチックな雰囲気

// クラスの変わりに型にメソッドを定義できる
type Vertex struct {
	X, Y float64
}

// メソッドは関数と違い、自身をレシーバ引数を関数に取る
// v: レシーバ Vertex: 型 (v Vertex): レシーバ引数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 構造体じゃなくてもいける
type MyInt int

func (mi MyInt) print() {
	fmt.Println(mi)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	var i MyInt = 99
	i.print()
}
