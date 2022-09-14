//Write a program to show the grade obtained by a student after they enter their total mark percentage.
package main

import "fmt"

func main() {

	var m float32
	fmt.Println("Enter your Mark in percentage")
	fmt.Scanf("%f", &m)

	if m > 90 {
		fmt.Println("A Grade")
	} else if m >= 80 {
		fmt.Println("B Grade")
	} else if m >= 70 {
		fmt.Println("C Grade")
	} else if m >= 60 {
		fmt.Println("D Grade")
	} else if m >= 50 {
		fmt.Println("E Grade")
	} else if m < 50 {
		fmt.Println("Failed")
	}
}
