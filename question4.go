//Write a program to check whether a student has passed or failed in a subject after he or she enters their mark (pass mark for a subject is 50 out of 100).
//The program should accept input from the user and output a message as “Passed” or “Failed.”
package main

import "fmt"

func main() {

	var m float32
	fmt.Println("Enter your Mark")
	fmt.Scanf("%f", &m)

	if m < 50 {
		fmt.Println("Failed")
	} else if m > 100 {
		fmt.Println("Please check,mark is out of 100")
	} else if m >= 50 {
		fmt.Println("Passed")
	}

}
