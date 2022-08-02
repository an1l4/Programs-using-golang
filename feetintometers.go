package main

import "fmt"

func main() {
	var userNumber int
	var feet1 float64 = 0.3048

	fmt.Println("enter a number")
	fmt.Scanf("%d", &userNumber)

	feet_into_meters := float64(userNumber) * feet1

	fmt.Println(feet_into_meters)

}
