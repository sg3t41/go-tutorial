package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// ユーザー定義型もレシーバとなり得る
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 他のパッケージで宣言されたレシーバ型のメソッドは作れない
// プリミティブ型でも同様
//func (i int) Abs() float64 {}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
