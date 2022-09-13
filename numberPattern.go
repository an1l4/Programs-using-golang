package main

import "fmt"

func main() {

	var n int

	fmt.Println("enter a number")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j," ")
		}
		fmt.Print("\n")
	}

}
