package main

import "fmt"

// struct(構造体)とは、フィールドの集まりです
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
