//menu driven program to do the basic mathematical operations such as addition, subtraction, multiplication and division

package main

import "fmt"

func main() {

	var option, num1, num2 int

	fmt.Println("enter 2 numbers")
	fmt.Scanf("%d %d", &num1, &num2)

	fmt.Println("select your option")
	fmt.Println("Addition? Press 1")
	fmt.Println("Subtraction? Press 2")
	fmt.Println("Multiplication? Press 3")
	fmt.Println("Division? Press 4")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println("Addition is : ", num1+num2)
	case 2:
		fmt.Println("Subtraction is : ", num1-num2)
	case 3:
		fmt.Println("Multiplication is : ", num1*num2)
	case 4:
		fmt.Println("Division is : ", num1/num2)
	default:
		fmt.Println("invalid option please check")

	}
}
