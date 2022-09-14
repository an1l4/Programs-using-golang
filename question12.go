//sort an array in descending order
package main

import "fmt"

func main() {

	var size int
	fmt.Println("enter the size of array")
	fmt.Scanf("%d", &size)

	array := make([]int, size)
	fmt.Println("enter array elements")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &array[i])

	}
	fmt.Println("array before sorting", array)
	var temp int

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if array[i] < array[j] {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	fmt.Println("array after sorting (desc)", array)

}
