package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var number string
	var result int64
	var temp int64

	fmt.Println("enter a number")
	fmt.Scan(&number)

	for i := 0; i < len(number); i++ {

		n, err := strconv.Atoi(string(number[i]))

		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(n)
		}

		power := math.Pow(float64(n), float64(len(number)))
		result = result + int64(power)
	}

	temp, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(temp)
	}

	if result == temp {
		fmt.Println("number is armstrong ")
	} else {
		fmt.Println("number is not armstrong")
	}

}
