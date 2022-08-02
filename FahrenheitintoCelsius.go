package main

import "fmt"

func main() {
	var userNumber int

	fmt.Println("enter a number")
	fmt.Scanf("%d", &userNumber)

	Fahrenheit_into_Celsius := (userNumber - 32) * 5 / 9

	fmt.Println(Fahrenheit_into_Celsius)
}
