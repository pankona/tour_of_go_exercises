package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ret := make(map[string]int)
	for _, field := range fields {
		ret[field] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
