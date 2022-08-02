package main

import "fmt"

func main() {
	var number int

	fmt.Println("enter a number")
	fmt.Scanf("%d", &number)

	i := 1

	for i <= 10 {
		result := number * i
		fmt.Printf("%d * %d = %d\n", number, i, result)
		i++

	}

}
