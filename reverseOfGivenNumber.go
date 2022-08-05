package main

import "fmt"

func main() {

	var number, mod, reverse int64

	fmt.Println("enter a number")
	fmt.Scan(&number)

	for number > 0 {
		mod = number % 10
		reverse = reverse*10 + mod
		number = number / 10

	}
	fmt.Println("reversed number is : ", reverse)

}
