package main

import "fmt"

func solutionFunction(nums []int) int {
	return 0
}

func main() {
	var n int
	fmt.Println("Enter the size of array")
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println("array", nums)
}
