package main

import "fmt"

func main() {
	//定义一个指针
	// var s1 *int	//错误
	//s1指向一个内存地址，相当于初始化
	s1 := new(int) //new函数：给基础类型分配地址	make函数：给引用类型分配地址
	//假设该内存地址为变量a的地址，*s1相当于获取a的值
	fmt.Println(*s1) //初始值0
	fmt.Println(s1)  //a的地址，相当于&a
	fmt.Println(&s1) //s1的地址

	a := 3
	s2 := &a
	fmt.Println(&a, s2) //a的内存地址
	fmt.Println(&s2)    //s2的内存地址

}
