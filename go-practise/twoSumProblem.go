package main

import "fmt"

func twoSum(){

	array:=[]int{2,3,4,5,6,7}

	sum:=10;
    m:=map[int]int{}
	for i , v := range array{
		find:=sum-v;
        
		if val , ok := m[find] ;  ok {
			fmt.Println("index",i,val)
		}
		m[v] = i
	}
}