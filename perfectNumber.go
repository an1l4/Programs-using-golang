package main

import (
	"fmt"
	"math"
)

func main() {

	var number int

	fmt.Println("enter a number")
	fmt.Scan(&number)

	for i := 2; i <= number; i++ {
		var flag int = 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				flag = 1
				break
			}

		}
		if flag == 0 {
			perfect := math.Pow(2, (float64(i) - 1))
			val := math.Pow(2, float64(i))
			total := perfect * (val - 1)
			fmt.Println("prime Number : ", i, "  ", "corresponding perfect number :", int64(total))
		}
	}

}
