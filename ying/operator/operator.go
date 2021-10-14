package main

import "fmt"

func main() {
	//++和--作为一个单独的语句
	// a := 1
	// fmt.Println(a ++)

	//赋值运算符
	//测试赋值运算符的运算顺序
	var a int = 1
	a += 3 * 4
	fmt.Println(a) //先做等号后边 13

}
