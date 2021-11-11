package main

import "fmt"

//当对象实现了接口，就可以用接口类型来接收该对象，相当于JAVA中的多态
func main() {

	//shape = &Triangle{}
	//shape.Draw()

	//实现方法为值类型,既可以接收值类型，也可以接收指针类型
	//var s1 Sayer
	//s1 = Dog{}	//可以接收
	//s1 = &Dog{}	//可以接收
	//s1.Say()

	//实现方法为指针类型,无法接收值类型，只能接收指针类型
	//var s1 Sayer
	//s1 = Dog{}		//无法接收
	//s1 = &Dog{}		//可以接收

	////接口亦可定义在函数内部
	//type B interface {
	//	Hello()
	//}

	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x) //type:string value:Hello 沙河
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x) //	type:int value:100
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x) //	type:bool value:true

	//使用断言进行类型判断
	switch v := x.(type) {
	case bool:
		fmt.Println(v)
	}
}
