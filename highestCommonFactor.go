package main

import "fmt"

var result []int

func main() {

	var number1, number2, largervalue, temp int
	var slice1, slice2 []int

	fmt.Println("enter 2 numbers")
	fmt.Scanf("%d", &number1)
	fmt.Scanf("%d", &number2)

	for i := 1; i <= number1; i++ {
		if number1%i == 0 {
			slice1 = append(slice1, i)
		}
	}
	//fmt.Println(slice1)

	for i := 1; i <= number2; i++ {
		if number2%i == 0 {
			slice2 = append(slice2, i)

		}
	}
	//fmt.Println(slice2)

	s := intersection(slice1, slice2)
	//fmt.Println(s)

	for _, element := range s {
		if element > temp {
			temp = element
			largervalue = temp
		}

	}
	fmt.Printf("Highest common factor of %d and %d is %d\n", number1, number2, largervalue)

}

func intersection(slice1, slice2 []int) []int {
	for _, i := range slice1 {
		flag := false
		for _, j := range slice2 {
			if i == j {
				flag = true
				break
			}
		}
		if flag {
			result = append(result, i)

		}

	}
	return result
}
