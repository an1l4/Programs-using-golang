package main

import "fmt"

func main() {

	flag := 0

	array := [...]int{2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 15, 17, 18, 19}

	fmt.Println("prime numbers from array : ")
	for i := 0; i < 15; i++ {
		flag = 0
		for j := 2; j < array[i]; j++ {
			if array[i]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && array[i] > 1 {
			fmt.Println(array[i])
		}
	}

}
