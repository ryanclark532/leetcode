package main

import (
	"fmt"
	"slices"
	"strings"
)

var bracketPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isValid(s string) bool {
	ss := strings.Split(s, "")
	slices.Sort(ss)

	for i := 0; i < len(ss)-1; i = i + 2 {

		if bracketPairs[ss[i]] != ss[i+1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
