//Using the ‘switch case,’ write a program to accept an input number from the user and output the day
package main

import "fmt"

func main() {

	var d int

	fmt.Println("Enter a number within the range of 1-7")
	fmt.Scanf("%d", &d)

	switch d {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid Entry")
	}
}
