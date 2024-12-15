package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapリテラルにはキーが必要です
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	// Vertexを省くことも可能
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
