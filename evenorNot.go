package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter anumber")
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Printf("%d is even \n", n)
	} else {
		fmt.Printf("%d is odd \n", n)
	}

}
