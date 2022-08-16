package main

import "fmt"

func main() {

	var array = [6]int{}
	var middle int
	var first int
	var last int
	var item int
	var flag int

	fmt.Println("enter sorted array element")
	for i := 0; i <= 5; i++ {
		fmt.Println("enter", i, "th array")
		fmt.Scan(&array[i])
	}
	fmt.Println("enter the item")
	fmt.Scan(&item)

	first = 0
	last = 5
	middle = (first + last) / 2

	for i := 0; i <= 5; i++ {
		if array[middle] < item {
			first = middle + 1
		} else if array[middle] == item {
			flag = 1
			fmt.Printf("The iteam %d found at %d index", item, middle)
			break

		} else {
			first = middle + 1
		}
		middle = (first + last) / 2

	}
	if flag == 0 {
		fmt.Println("item not found")
	}

}
