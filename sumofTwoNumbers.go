package main

import "fmt"

func main() {

	var (
		number1 int
		number2 int
		sum     int
	)

	fmt.Println("Welcome")
	fmt.Println("Please enter two numbers")
	fmt.Scanf("%d%d", &number1, &number2)
	sum = number1 + number2
	fmt.Printf("sum is %d \n", sum)

	//average
	var (
		a       float64
		b       float64
		c       float64
		average float64
	)
	fmt.Println("Please enter 3 numbers")
	fmt.Scanf("%f%f%f", &a, &b, &c)
	average = (a + b + c) / 3
	fmt.Printf("Average is %f\n", average)

}
