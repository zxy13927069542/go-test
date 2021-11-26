package main

import "fmt"

func main() {
	//每遇到一个const关键字，iota初始化为0，每新增一行变量声明iota加一
	const (
		A1 = iota //0
		A2        //1
		A3        //2
	)
	fmt.Println(A1, A2, A3)

}
