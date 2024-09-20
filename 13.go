package main

import (
	"fmt"
	"strings"
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	romans := strings.Split(s, "")
	total := 0
	current := 0
	prev := 0
	lastSubtracted := false
	for i := len(romans) - 1; i >= 0; i-- {
		if lastSubtracted {
			lastSubtracted = false
			continue
		}
		current = romanMap[romans[i]]
		if i == 0 {
			total += current
		} else {
			prev = romanMap[romans[i-1]]
			if prev < (current / 2) {
				total = total + (current - prev)
				lastSubtracted = true
			} else {
				total += current
				lastSubtracted = false
			}
		}

	}

	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
