//calculate the area of a given object.

package main

import (
	"fmt"
	"math"
)


func circleArea(radius float64) float64 {
	return math.Pi * (radius * radius)
}

func squareArea(length float64) float64 {
	return length * length
}

func rectangleArea(lengthrec, width float64) float64 {
	return lengthrec * width
}

func triangleArea(base, height float64) float64 {
	return (base * height) / 2
}

func main() {

	var choice int

	fmt.Println("enter the choice")
	fmt.Println("1.circle")
	fmt.Println("2.square")
	fmt.Println("3.rectangle")
	fmt.Println("4.triangle")

	fmt.Scanf("%d", &choice)
	var radius, length, lengthrec, width, base, height float64

	switch choice {
	case 1:
		fmt.Println("enter the radius")
		fmt.Scanf("%f", &radius)
		fmt.Println("Area of circle :",circleArea(radius))
	case 2:
		fmt.Println("enter the length")
		fmt.Scanf("%f", &length)
		fmt.Println("Area of square :",squareArea(length))
	case 3:
		fmt.Println("enter the length and width")
		fmt.Scanf("%f %f", &lengthrec, &width)
		fmt.Println("Area of rectangle :",rectangleArea(lengthrec, width))
	case 4:
		fmt.Println("enter the base and height")
		fmt.Scanf("%f %f", &base, &height)
		fmt.Println("Area of triangle :",triangleArea(base, height))
	default:
		fmt.Println("invalid entry")

	}

}
