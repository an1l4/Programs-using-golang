package main

import "fmt"

func main() {

	var number int

	fmt.Println("enter a number")
	fmt.Scan(&number)

	for i := 2; i <= number; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}

}
