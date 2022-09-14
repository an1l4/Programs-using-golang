//string palindrome

package main

import "fmt"

func main() {

	var word string

	fmt.Println("enter a word")
	fmt.Scanf("%s", &word)

	var reverse string

	for i := len(word) - 1; i >= 0; i-- {
		reverse = reverse + string(word[i])

	}
	if word == reverse {
		fmt.Println("Entered string is a palindrome")
	} else {
		fmt.Println("Entered string is not a palindrome")
	}

}
