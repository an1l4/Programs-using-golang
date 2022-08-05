package main

import "fmt"

func main() {

	var number, power int
	var result int = 1

	fmt.Println("enter a number")
	fmt.Scan(&number)

	fmt.Println("enter the power")
	fmt.Scan(&power)

	for i := 1; i <= power; i++ {
		result = result * number
	}
	fmt.Println("power of", number, "is : ", result)

}
