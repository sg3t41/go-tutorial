// https://go-tour-jp.appspot.com/basics/9
package main

import "fmt"

/*
var宣言で変数ごとに初期化子を(initiarizer)を指定して、
変数宣言と同時に初期値を設定できます。

初期化子が指定されている場合、変数の型を省略することもできます。
その場合に変数の型は、初期値の型から自動判別で定義されます。
*/
var i, j int = 1, 2

func main() {
	// 変数の型を省略する場合
	var c, java, python = false, true, "no!"
	fmt.Println(i, j, c, java, python)
	//1 2 false true no!
}
