package main

import (
	"flag"
	"fmt"
)

//直接返回变量指针
var id = flag.String("id", "0", "id")
var name = flag.String("name", "郑晓颖", "name")
var age = flag.Int("age", 22, "age")
var other int

func init() {
	//将flag与现有变量绑定
	flag.IntVar(&other, "other", 1, "other")
}

func main() {
	//解析运行命令参数到指定的变量
	flag.Parse()
	fmt.Println("没被解析的变量:", flag.Args())
	fmt.Println("没被解析的变量数:", flag.NArg())
	fmt.Println("被解析的变量数:", flag.NFlag())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Args()[i])
	}
	fmt.Println("id:", *id)
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("other:", other)
}
