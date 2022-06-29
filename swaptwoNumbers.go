package main

import "fmt"

func main() {
	var (
		a    = 20
		b    = 10
		temp int
	)
	temp = a
	a = b
	b = temp
	fmt.Printf("value of a %d and value of b %d \n", a, b)
	swap(23, 45)
}

func swap(x, y int) {
	fmt.Printf("Before swaping %d %d \n", x, y)
	y = x + y
	x = y - x
	y = y - x
	fmt.Printf("After swap %d %d \n", x, y)

}
