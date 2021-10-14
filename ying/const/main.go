package main

import "fmt"

const PI = 3.14

//批量声明
const (
	B = 1
	C = 2
	D = 3
)

//F与H自动赋值为"aaa"
const (
	G = "aaa"
	F
	H
)

func main() {
	//常量可以在不同作用域中重名
	const PI = 3.12
	fmt.Println(PI)

	//在函数中常量即使不使用也不会编译报错
	const A = 3.12

	//F与H自动赋值为"aaa"
	fmt.Println(F, H)

}
