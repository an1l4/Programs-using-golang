package main

import "fmt"

func fibonacciNumber(n int) {
	var n3, n1, n2 = 0, 0, 1

	for i := 1; i <= 10; i++ {
		fmt.Println(n1)
		n3 = n1 + n2
		n1 = n2
		n2 = n3
	}

}
func main() {
	fibonacciNumber(10)

}
