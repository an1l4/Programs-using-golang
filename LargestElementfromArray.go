package main

import "fmt"

func main() {

	var limit int

	fmt.Println("enter the array limit")
	fmt.Scanf("%d", &limit)

	var array = make([]int, limit)

	for i := 0; i < limit; i++ {
		fmt.Println("enter", i, "element")
		fmt.Scan(&array[i])
	}
	fmt.Println(array)

	temp := array[0]

	for i := 1; i < limit; i++ {
		if temp < array[i] {
			temp = array[i]
		}
	}
	fmt.Println("Largest number in the array is :", temp)

}
