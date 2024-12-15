package main

import "fmt"

func fibonacci() func() {
	a, b := 0, 1

	return func() {
		c := a
		defer fmt.Println(c)
		a, b = b, a+b
	}
}

func main() {
	f := fibonacci()
	for range 5 {
		f()
	}

	f()
	f()
	f()

	for range 2 {
		f()
	}

	defer f()

	f()

	func() {
		defer f()
		func() {
			f()
			f()
		}()
	}()
}

