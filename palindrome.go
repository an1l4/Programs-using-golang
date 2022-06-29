package main

import "fmt"

func isPalindrome(x int) bool {

	result := 0
	newValue := x
	for x > 0 {
		//fmt.Println(num)
		remainder := x % 10
		//fmt.Println(remainder)
		result = (result * 10) + remainder
		//fmt.Println(res)
		x /= 10
		//fmt.Println(num)
		fmt.Println(x)

	}
	if newValue == result {
		return true
	}
	return false

}

func main() {
	fmt.Println(isPalindrome(168))
	//fmt.Println(isPalindrome(121))
	//fmt.Println(isPalindrome(12321))
}
