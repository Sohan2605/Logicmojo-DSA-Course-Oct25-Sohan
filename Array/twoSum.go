package main

import "fmt"

func twoSum(nums []int , target int)[]int{
	m := make(map[int]int)
    
	for i ,num := range nums {
		if val , ok := m[target-num]; ok {
			return []int{val,i}
		}

		m[num] = i 
	}
	return []int{}
}

func main(){
	var n , target int
	fmt.Println("enter the size of the array")
	fmt.Scan(&n)

	nums := make([]int , n)
	fmt.Println("enter elements")
	for i:=0; i<n; i++{
		fmt.Scan(&nums[i])
	}
    fmt.Println("enter targets")

	fmt.Scan(&target)

	result := twoSum(nums,target)
	fmt.Println("Result:", result)
	

}