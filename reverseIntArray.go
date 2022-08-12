package main

import "fmt"

func main() {

	var rev_array = [9]int{}
	array := [...]int{2, 67, 12, 59, 23, 7, 90, 23, 23}

	var j int = 8
	for i := 0; i < 9; i++ {
		rev_array[j] = array[i]
		j = j - 1
	}
	fmt.Println("reversed array : ")

	for i := 0; i < 9; i++ {
		fmt.Printf("%d  ", rev_array[i])
	}
	fmt.Println()

}
