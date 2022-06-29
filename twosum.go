package main

import "fmt"

func main() {
	fmt.Println("twosum")
	var arr1 = []int{2, 7, 11, 15}
	var arr2 = []int{3, 2, 4}
	var arr3 = []int{3, 3}
	fmt.Println(twoSum(arr1, 9))
	fmt.Println(twoSum(arr2, 6))
	fmt.Println(twoSum(arr3, 6))

}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		var a int = nums[i]
		for j := i + 1; j < len(nums); j++ {
			var b int = nums[j]
			if a+b == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}
