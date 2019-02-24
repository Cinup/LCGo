package main

import "fmt"

//本地测试使用
func main() {
	isexit := make(map[uint8]int)
	isexit['s'] =1
	fmt.Println(isexit['s'])
	s,y := isexit['a']
	fmt.Println(s,y)
}
