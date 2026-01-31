package main


import "fmt"

func someFunc(num int){
	fmt.Println(num)
}

func main(){
   ch:=make(chan string)
   go func(){
     ch <- "data"
   }()

   msg:=<-ch
   fmt.Println("sohan",msg)
}