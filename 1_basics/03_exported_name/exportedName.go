// https://go-tour-jp.appspot.com/basics/3

/*
Goでは、最初の文字が大文字で始まる名前は、外部から参照できるエクスポートされた名前(exported name)です。
小文字で始まる"pi"や"hoge"は外部から参照できません。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		math.piは先頭が小文字なので内部に閉じ込められた値です。
		外部から参照できずエラーが起きます。
		Error: undefined: math.pi
	*/
	//fmt.Println(math.pi)

	// math.Piの値は円周率です。
	fmt.Println(math.Pi)
}
