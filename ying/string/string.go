package main

import (
	"fmt"
)

func main() {
	//单引号，表示一个字符，rune类型，也叫int32
	//go使用utf-8编码，一个字符占用4个字节，一个表示类型，另外3个字节表示值
	s1 := 'a'
	fmt.Printf("s1的值为%v,类型为%T\n", s1, s1)  //97	int32
	fmt.Printf("s1的值为%#v,类型为%T\n", s1, s1) //97	int32

	//双引号，表示一个字符串,类型为string
	s2 := "哈哈哈"
	fmt.Printf("s2的值为%s,类型为%T\n", s2, s2)  //哈哈哈	string
	fmt.Printf("s2的值为%#v,类型为%T\n", s2, s2) //"哈哈哈"	string

	//反引号，里面的值都会原样输出，不会被转义
	s3 := `/t`
	fmt.Printf("s3的值为%s,类型为%T\n", s3, s3)  ///t	string
	fmt.Printf("s3的值为%v,类型为%T\n", s3, s3)  ///t	string
	fmt.Printf("s3的值为%#v,类型为%T\n", s3, s3) //"/t"	string

	//输出s3的长度
	fmt.Println("s3的长度为:" + fmt.Sprintf("%d", len(s3))) //int转string 以及string之间的拼接

	//将string转换为字符数组即(rune[])
	s4 := []rune(s3)
	s4[0] = 'A'
	fmt.Println(s4)
	for i := 0; i < len(s4); i++ {
		fmt.Println(string(s4[i]))
	}
	fmt.Printf("s4的类型为：%T\n", s4) //[]int32 即[]rune

	//将rune数组转换成string
	s5 := string(s4)
	fmt.Println(s5)

	//测试字符串中英文字符与中文字符的区别
	s6 := "AAA哈哈哈"

	//直接遍历原数组
	for _, v := range s6 {
		fmt.Printf("当前字符为%c,类型为%T\n", v, v)
	}

	//遍历切片
	s7 := []rune(s6)
	for _, v := range s7 {
		fmt.Printf("当前字符为%c,类型为%T\n", v, v)
	}

}
