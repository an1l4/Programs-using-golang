//The program should accept 3 inputs from the user and calculate simple interest for the given inputs. Formula: SI=(P*R*n)/100)

package main

import "fmt"

func main() {

	var p int
	var r, n float32
	fmt.Println("Enter the Principal amount (P)")
	fmt.Scanf("%d", &p)

	fmt.Println("Enter the Interest Rate (R)")
	fmt.Scanf("%f", &r)

	fmt.Println("Enter the number of years (n)")
	fmt.Scanf("%f", &n)

	simpleInterest := (float32(p) * r * n) / 100

	fmt.Println("Simple nterest SI =", simpleInterest)
}
