// https://go-tour-jp.appspot.com/basics/1

/*
Goのプログラムはパッケージ(package)で構成されています。
プログラムはmainパッケージから開始されます。
このプログラムでは"fmt"と"math/rand"のパッケージをimportしています。
*/
package main

import "fmt"
import "math/rand"

func main() {
	/*
		言語のルールで、パッケージ名はインポートパスの最後の要素と同じ名前になります。
		インポートパスがmath/randのパッケージは、package randのファイル群で構成されている、ということになります。
	*/
	// rand.Intnは擬似乱数を返す関数です。
	randomNumber := rand.Intn(10)
	fmt.Println("My favorite number is", randomNumber)
}
