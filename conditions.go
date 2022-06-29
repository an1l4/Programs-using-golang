package main

import "fmt"

func main() {
	var (
		number1 int
		number2 int
		choice  int
		result  int
	)
	fmt.Println("Enter two numbers :")
	fmt.Scanf("%d%d", &number1, &number2)
	fmt.Print("1 for addition\n2 for substraction\n3 for multiplication\n4 for division\nEnter your choice:")
	fmt.Scanf("%d", &choice)

	if choice == 1 {
		result = number1 + number2
		fmt.Printf("result is :%d", result)

	} else if choice == 2 {
		result = number1 - number2
		fmt.Printf("result is :%d", result)

	} else if choice == 3 {
		result = number1 * number2
		fmt.Printf("result is :%d", result)

	} else if choice == 4 {
		result = number1 / number2
		fmt.Printf("result is :%d \n", result)

	} else {
		fmt.Println("Open your eyes and read the choice again")
	}

}
