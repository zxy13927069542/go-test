package main

import "fmt"

func main() {

	//删除切片元素
	s16 := [...]int{3, 4, 5}
	s17 := s16[:]
	s17 = append(s17[0:0], s17[1:]...)
	fmt.Println(s17, len(s17), cap(s17))
	fmt.Println(s16)

}
