package main

import "strings"

const repeat = 5

func Repeat(s string) string {

	var repeated strings.Builder
	// var repeated string

	for range repeat {
		// repeated += s
		// Bench mark of this code is very bad because string are immutable hence we are using strings.Builder
		repeated.WriteString(s)
	}

	return repeated.String()
	// return repeated

}

func main() {

}