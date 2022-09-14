//Accept two inputs from the user and output its sum.

package main

import "fmt"

func main() {

	var integer int
	var floating float64

	fmt.Println("enter a non decimal numbers")
	fmt.Scanf("%d", &integer)
	fmt.Println("enter a decimal numbers")
	fmt.Scanf("%f", &floating)

	fmt.Println("Sum of the values is :", float64(integer)+floating)

}
