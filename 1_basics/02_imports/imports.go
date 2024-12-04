// https://go-tour-jp.appspot.com/basics/2

package main

/*
import "fmt"
import "math/rand"

複数のパッケージをインポートする場合は上記のように書くこともできますが、
それらをグループ化した下記の書き方が推奨されています。
*/
import (
	"fmt"
	"math"
)

func main() {
	/*
		%gはGoのフォーマット指定子と呼ばれるものです。
		%gは数字を最適な形式(小数点形式、指数表記、末尾のゼロ省略などを数値に自動で合わせる)で表現します。

		math.Sqrtは平方根を計算します。
	*/
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
