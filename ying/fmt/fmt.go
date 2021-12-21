package main

import "fmt"

func main() {
	// //1.scan
	// var s1 string
	// //2.获取用户输入，赋值给变量
	// fmt.Scan(&s1)	//输入asdf asdf
	// fmt.Println(s1)	//输出asda

	//1.scanf：需要按照格式输入
	var (
		name string
		age  int
		sex  string
	)
	fmt.Scanf("1:%s 2:%d 3:%s\n", &name, &age, &sex) //输入1:name 2:23 3:nan
	fmt.Println(name, age, sex)                      //输出name 23 nan
}

