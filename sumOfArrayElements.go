package main

import "fmt"

func main() {

	var sum int

	var array = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 5; i++ {
		sum = sum + array[i]

	}
	fmt.Println("sum of array elements", sum)

}
