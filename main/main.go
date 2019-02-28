package main

import "fmt"
func test(){
	x,y := 1,2
	x += 10
	y += 10
	fmt.Println(x,y)
	defer func (i int){
		fmt.Println(i,y)
	}(x)
}

func main(){
	test()
}