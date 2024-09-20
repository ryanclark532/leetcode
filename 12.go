package main

var vals = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var chars = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	i := 0
	final := ""
	for num > 0 {
		if num >= vals[i] {
			final = final + chars[i]
			num -= vals[i]
		} else {
			i++
		}
	}
	return final
}

func main() {
	intToRoman(3749)
}
