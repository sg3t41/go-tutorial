package main

import (
	"fmt"
	"io"
	"strings"
)

/*
io パッケージは、データストリームを読むことを表現する io.Reader インタフェースを規定しています。
Goの標準ライブラリには、ファイル、ネットワーク接続、圧縮、暗号化などで、このインタフェースの 多くの実装 があります。
io.Reader インタフェースは Read メソッドを持ちます

	type Reader interface {
		Read(p []byte) (n int, err error)
	}
*/
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
