package main

import (
	"fmt"
)

var phoneMap = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	backtrack(&result, "", digits, 0)
	return result
}

func backtrack(result *[]string, combination string, digits string, index int) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	d := rune(digits[index])
	letters := phoneMap[d]

	for _, l := range letters {
		combination += string(l)
		backtrack(result, combination, digits, index+1)
		combination = combination[:len(combination)-1] // Remove last character
	}
}

func main() {
	digits := "234"
	combinations := letterCombinations(digits)
	fmt.Println(combinations) // Output: ["ad" "ae" "af" "bd" "be" "bf" "cd" "ce" "cf"]
}
