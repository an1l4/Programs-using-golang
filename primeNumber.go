package main

import "fmt"

func main() {
	var n int
	var flag int = 0

	fmt.Println("Enter a number")
	fmt.Scanf("%d", &n)

	for i := 2; i < n/2; i++ {

		if n%i == 0 {
			flag = 1
			break
		}

	}
	if flag == 0 {
		fmt.Println("number is prime")
	} else {
		fmt.Println("number is non prime")
	}

}
