//Incuome tax is calculated

package main

import "fmt"

func main() {

	var income int64

	fmt.Println("Enter the anual income")
	fmt.Scanf("%d", &income)

	if income <= 250000 {
		fmt.Println("No Tax")
	} else if income > 1000000 {
		fmt.Println("Tax amount is : ", (income * 5 / 100))
	} else if income > 500000 {
		fmt.Println("Tax amount is : ", (income * 20 / 100))
	} else if income > 250000 {
		fmt.Println("Tax amount is : ", (income * 30 / 100))
	}
}
