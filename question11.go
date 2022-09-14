//find the number of even numbers in an array

package main

import "fmt"

func main() {

	var size int
	fmt.Println("enter the size of an array")
	fmt.Scanf("%d", &size)

	array := make([]int, size)
	fmt.Println("enter the array elements")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println()
	var count int = 0
	for i := 0; i < size; i++ {
		if array[i]%2 == 0 {
			fmt.Println(array[i])
			count++
		}
	}
	fmt.Println("number of even numbers in array", count)

}
