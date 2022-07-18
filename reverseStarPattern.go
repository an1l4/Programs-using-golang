package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number")
	fmt.Scanf("%d", &n)

	for i := n; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}
