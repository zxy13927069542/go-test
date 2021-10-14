package main

import "fmt"

func f3() {
	fmt.Println("f3")
}

//函数类型亦可作为参数类型
func f4(s4 func()) {
	f3()
}

func main() {
	//函数类型
	s1 := f1
	fmt.Printf("%T\n", s1) //func(string,int) int

	//1.匿名函数定义及调用
	//2.函数内部无法声明带名字的函数
	//3.调用方式一
	f2 := func() {
		fmt.Println("匿名函数")
	}
	f2()
	//4.调用方式二
	func() {
		fmt.Println("匿名函数")
	}()

	//闭包
	f6(f7(f5, 3, 2))

}

//1.第一个()表示参数
//2.第二个()表示返回值
func f1(v1 string, v2 int) (rtValue int) {
	//3.如果返回值有命名，则可以不用再声明
	// rtValue := 3
	rtValue = 3
	return //4.返回值有命名的情况下，return后面可以省略
}

//5.如果返回值只有一个并且不命名，可以不写括号
func f2() int {
	//6.如果返回值无命名，则仍需声明
	rtValue := 3
	return rtValue
}

//1.功能函数
func f5(x, y int) {
	fmt.Println(x + y)
}

//2.调用函数：欲调用功能函数，但参数与功能函数的函数类型不匹配，所以需要适配器函数来进行转换
func f6(f func()) {
	f()
}

//3.适配器函数：功能函数作为参数，输出一个无参函数
//4.参数：不仅要包含功能函数的函数地址，还需包含功能函数的参数x，y
//5.闭包 = 函数 加 外部变量的引用
func f7(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}

// f6(f7(f5,3,2))
