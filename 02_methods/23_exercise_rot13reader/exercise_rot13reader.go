package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return 0, err
	}

	// ROT13変換
	// i,v にしてvを使ったら値のコピーを使われるので今回はiを使う
	for i := range b {
		switch {
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		default:
			continue
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!!!") // 入力データ
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r) // 出力先は os.Stdout
	if err != nil {
		fmt.Println("Error:", err) // エラーがあれば表示
	}
}

