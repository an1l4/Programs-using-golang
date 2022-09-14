//Write a program to print the multiplication table of given numbers.
package main

import "fmt"

func main() {
	var n int

	fmt.Println("enter a number")
	fmt.Scanf("%d", &n)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "*", n, "=", i*n)
	}

}
