package main

import "fmt"

func main() {

	var number int
	var factorial int = 1

	fmt.Println("enter a number")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		factorial = factorial * i
	}
	fmt.Println("factorial of", number, "is : ", factorial)

}
