//program to find the grade of a student during his academic year.

package main

import "fmt"

func main() {

	var wt int

	fmt.Println("Please enter your written test score")
	fmt.Scanf("%d", &wt)

	var le int

	fmt.Println("Please enter your lab exam score")
	fmt.Scanf("%d", &le)

	var assignments int

	fmt.Println("Please enter your assignments score")
	fmt.Scanf("%d", &assignments)

	var grade float32

	grade = (float32(wt)*70)/100 + (float32(le)*20)/100 + (float32(assignments)*10)/100

	fmt.Println("your Grade is : ", grade)

}
