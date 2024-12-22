package tp

import "fmt"

type t struct {
	S string
}

// 直接インスタンス化させるのを隠蔽化すればnilチェック不要
func New() *t {
	return &t{}
}

func (t *t) M() {
	fmt.Println(t.S)
}
