package main

import "fmt"

func main() {

	var large1, large2 int = 0, 0

	array := [5]int{46, 364, 45, 1356, 968}

	large1 = array[0]

	for i := 1; i < 5; i++ {
		if large1 < array[i] {
			large2 = large1
			large1 = array[i]
		} else if large2 < array[i] {
			large2 = array[i]

		}
	}

	fmt.Println("second largest element is  :", large2)

}
