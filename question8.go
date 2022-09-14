//Write a program to find the sum of all the odd numbers for a given limit
package main

import "fmt"

func main() {

	var limit int
	var sum int = 0

	fmt.Println("enter a limit")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("sum of odd numbers within the limit is:", sum)
}
