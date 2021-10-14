package main

import "fmt"

func main() {
	age := 13
	switch age {
	case 17:
		fmt.Println("当前年龄为17岁")
	case 16:
		fmt.Println("当前年龄为16岁")
	default:
		fmt.Println("当前年龄既不是17岁，也不是16岁")

	}
}
