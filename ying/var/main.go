package main

import "fmt"

var (
	name   string
	age    int
	sex    string
	height float32 = 0.32 //全局变量即使未被使用也不会编译报错 非全局变量声明后必须使用
	weight         = 0.32 //自动推导类型
)

func main() {
	name = "郑晓颖"
	age = 22
	sex = "男"

	//简短变量声明 只能在函数内使用
	// a := "aaa"
	// fmt.Println(a)

	// fmt.Println(name, age, sex)

	var height = 0.31
	fmt.Println(height)
	printHeigth()

	//使用s1,s2接受函数返回的值
	s1, s2 := ReturnTwoValue()
	fmt.Println(s1, s2)

}

func printHeigth() {
	fmt.Println(height)
}

//一个返回两个字符串的函数
func ReturnTwoValue() (s1 string, s2 string) {
	return "aaa", "aaa"
}
