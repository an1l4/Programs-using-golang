package main

import "fmt"

func main() {

	var limit int = 1
	var n int
	fmt.Println("enter a number")
	fmt.Scanf("%d", &n)

	for i := n; i >= 1; i-- {
		for j := n; j >= limit; j-- {
			fmt.Print(j, " ")
		}
		limit = limit + 1
		fmt.Print("\n")
	}
}
