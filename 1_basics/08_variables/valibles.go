// https://go-tour-jp.appspot.com/basics/8
package main

import "fmt"

/*
varで変数を宣言できます。
varステートメントは、package内、関数内で使用できます。

関数の引数と同様に、同型かつ複数の場合は最後に型を指定するだけで変数リストを宣言できます。
*/
var c, java, python bool

func main() {
	var i int
	fmt.Println(i, c, java, python)
	//0 false false false
}
