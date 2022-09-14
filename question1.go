//Accept a char input from the user and display it on the console.

package main

import "fmt"

func main() {

	var n string

	fmt.Println("enter your name")
	fmt.Scanf("%s", &n)

	fmt.Println("Hello", n)

}
