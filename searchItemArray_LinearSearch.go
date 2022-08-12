package main

import "fmt"

func main() {

	var array = [5]int{}
	var flag int = 0
	var item int = 0

	fmt.Println("enter array elements : ")
	for i := 0; i < 5; i++ {
		fmt.Println("Enter", i, "th", "element")
		fmt.Scan(&array[i])
	}
	fmt.Println("enter the item you want to search")
	fmt.Scan(&item)

	for i := 0; i < 5; i++ {
		if item == array[i] {
			flag = 1
			fmt.Printf("item %d found at index %d\n", item, i)
		}
	}
	if flag == 0 {
		fmt.Println("item not found")

	}

}
