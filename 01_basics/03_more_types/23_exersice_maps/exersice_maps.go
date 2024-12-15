package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	m := make(map[string]int)
	for _, sv := range ss {
		if _, ok := m[sv]; ok {
			m[sv]++
		} else {
			m[sv] = 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
