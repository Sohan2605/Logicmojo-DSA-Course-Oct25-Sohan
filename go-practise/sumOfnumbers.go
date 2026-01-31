package main

import "fmt"

func sumOfNumbers(){
	number:=[]int{2,3,4,5}
	sum:=0

    for _,val := range number{
		sum+=val;
	}
  fmt.Println("sum:",sum)
}