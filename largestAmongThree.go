package main

import "fmt"

func main() {

	var (
		number1 int
		number2 int
		number3 int
	)
	fmt.Println("Enter 3 numbers :")
	fmt.Scanf("%d%d%d", &number1, &number2, &number3)

	if number1 > number2 && number1 > number3 {
		fmt.Printf("%d is greatest\n ", number1)
	}
	if number2 > number1 && number2 > number3 {
		fmt.Printf("%d is greatest\n ", number2)
	}
	if number3 > number1 && number3 > number2 {
		fmt.Printf("%d is greatest\n ", number3)

	}

}
