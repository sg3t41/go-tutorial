package main

import (
	"fmt"
	"time"
)

// Goのプログラムはエラーの状態をerror値で表現します

/*
type error interface {
    Error() string
}
*/

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
