//program to interchange the values of two arrays.

package main

import "fmt"

func main() {

	var size int

	fmt.Println("enter the size of arrays")
	fmt.Scanf("%d", &size)

	elements1 := make([]int, size)
	elements2 := make([]int, size)

	fmt.Println("enter the elements of array 1")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &elements1[i])

	}

	fmt.Println("enter the elements of array 2")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &elements2[i])

	}
	fmt.Println("before swap")
	fmt.Println()

	fmt.Println("Array1", elements1)
	fmt.Println("Array2", elements2)
	fmt.Println()

	temp := make([]int, size)

	temp = elements1
	elements1 = elements2
	elements2 = temp

	fmt.Println("after swap")
	fmt.Println()

	fmt.Println("Array1", elements1)
	fmt.Println("Array2", elements2)

}
