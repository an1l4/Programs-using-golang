package main

import "fmt"

func main() {

	var num1, num2, result int

	fmt.Println("enter 2 numbers")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	for i := 1; i <= num2; i++ {
		result = result + num1
	}
	fmt.Println("multiplication is : ", result)

}
