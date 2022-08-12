package main

import "fmt"

func main() {
	var arr3 = [11]int{}

	arr1 := [...]int{1, 2, 3, 4, 4}

	arr2 := [...]int{5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		arr3[i] = arr1[i]
	}
	var j int = 5
	for i := 0; i < 6; i++ {
		arr3[j] = arr2[i]
		j = j + 1
	}

	fmt.Println("merged array is :")
	for i := 0; i < 11; i++ {
		fmt.Printf("%d  ", arr3[i])
	}
	fmt.Println()

}
