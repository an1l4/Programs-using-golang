package main

import "fmt"

func romanToInt(s string) int {
	var romanNumbers = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var (
		convertedNumber = 0
		total           = 0
		lastNumber      = 0
	)
	for i := 0; i < len(s); i++ {
		position := s[len(s)-(i+1) : len(s)-i]
		convertedNumber = romanNumbers[position]

		if convertedNumber < lastNumber {
			total = total - convertedNumber
		} else {
			total = total + convertedNumber
		}
		lastNumber = convertedNumber

	}
	return total

}
func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))

}
