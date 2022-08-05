package main

import "fmt"

func main() {

	var num1, num2, temp int

	fmt.Println("enter 2 number")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	if num1 > num2 {
		for num2 != 0 {
			temp = num1 % num2
			num1 = num2
			num2 = temp
		}
		fmt.Println("highest common factor is : ", num1)
	} else if num2 > num1 {
		for num1 != 0 {
			temp = num2 % num1
			num2 = num1
			num1 = temp
		}
		fmt.Println("highest common factor is : ", num2)

	}
}
