package main

import "fmt"

func main() {

	var limit int

	fmt.Println("enter the array limit")
	fmt.Scan(&limit)

	var array = make([]int, limit)

	for i := 0; i < limit; i++ {
		fmt.Println("enter the", i, "element")
		fmt.Scan(&array[i])
	}

	fmt.Println(array)

}
