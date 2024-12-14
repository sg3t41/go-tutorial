package main

import "fmt"

// 省略宣言をすると、値によってコンパイラが自動的に型を推測します。
func main() {
	// v := 42           // int
	// v := 3.142        // float64
	v := 0.867 + 0.5i // complex128
	fmt.Printf("v is of type %T\n", v)
}
