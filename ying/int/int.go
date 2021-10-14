package main

import "fmt"

func main() {
	//整形只能表示十进制，八进制，十六进制
	var a int = 1          //默认为十进制
	fmt.Printf("%d \n", a) //以十进制输出 1
	fmt.Printf("%b \n", a) //以二进制输出 01

	var b int = 076        //0开头表示八进制
	fmt.Println(b)         //以十进制输出 62
	fmt.Printf("%o \n", b) //以八进制输出 76

	var c int = 0xff       //0x开头表示十六进制
	fmt.Println(c)         //以十进制输出 255
	fmt.Printf("%x \n", c) //以八进制输出 ff

	//%T能输出变量的类型
	d := 3
	fmt.Printf("变量类型:%T \n", d) //变量类型int

}
