package main

import "fmt"

func main() {

	s1 := []string{"郑晓颖", "郑兵颖"}
	for _, val := range s1 {
		val = "正"
		fmt.Println(val)
	}
	fmt.Println(s1) //[郑晓颖 郑兵颖]

}
