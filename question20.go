//number pattern

package main

import "fmt"

func main() {
	var value int = 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(value, " ")
			value++
		}
		fmt.Println()
	}

}
