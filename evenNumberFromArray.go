package main

import "fmt"

func main() {

	array := [8]int{43, 67, 23, 86, 13, 47, 94, 28}

	for i := 0; i < 8; i++ {
		if array[i]%2 == 0 {
			fmt.Println("even numbers :", array[i])
		}
	}

}
