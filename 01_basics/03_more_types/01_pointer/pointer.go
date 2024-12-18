package main

import "fmt"

// ポインタは値のメモリアドレスを指します
func main() {
	/*
					&記号の使い方

					&は変数が格納されているアドレスを取得するために使われる
					i := 100
					p := &i
				 では、変数iのアドレスを変数pに代入していることになる

				 *記号の使い方

				1. ポインタ型を作る
					k := 100
					var j *int
					j = &k

					jはint型のポインタ型となる
					jは、*int型と宣言することで、指す先がintの値であるポインタ(アドレス情報)を格納するための変数である、と宣言できる

				2. ポインタ型が指す値にアクセスする
					*int型のポインタ変数に対して、そのアドレスが指す値を得る
					func main() {
			 	 		t := &T{Name: "Alice"} // ポインタ型の変数t

						// *tを使って、ポインタtが指す実際の値にアクセス
		    		fmt.Println((*t).Name) // "Alice"
					}
					note: Goでは、(*t).Nameとしなくても、内部的に自動でデリファレンスするのでt.Nameでも大丈夫
	*/

	i, j := 42, 2701
	p := &i
	fmt.Println(*p) // 42

	p = &j
	*p = *p / 32 // 84
	fmt.Println(j)
}
