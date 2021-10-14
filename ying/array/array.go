package main

import "fmt"

func main() {
	//数组类型
	//数组长度为类型的一部分
	var s1 [3]int
	var s2 [4]int
	fmt.Printf("s1的类型为%T\n", s1) //s1的类型为[3]int
	fmt.Printf("s2的类型为%T\n", s2) //s2的类型为[4]int

	//数组初始化
	//第一种
	var s3 [3]int = [3]int{1, 1, 1}
	fmt.Println(s3)
	//第二种
	s4 := [3]int{1, 1, 1}
	fmt.Println(s4)
	//第三种
	s5 := [...]int{1, 1, 1}
	fmt.Printf("s5的类型为%T\n值为%v\n", s5, s5) //类型为[3]int
	//第四种
	s6 := []int{1, 1, 1}
	fmt.Printf("s6的类型为%T\n值为%v\n", s6, s6) //类型为[]int,切片，长度可变
	//第五种
	s7 := [5]int{0: 1, 4: 3}               //根据索引来进行初始化
	fmt.Printf("s7的类型为%T\n值为%v\n", s7, s7) //[1 0 0 0 3]

	//数组遍历
	s8 := [...]int{3, 4, 5, 6}
	//第一种
	for i := 0; i < len(s8); i++ {
		fmt.Println(s8[i])
	}

	//第二种
	for i1, v1 := range s8 {
		fmt.Printf("\n索引为%v的值为%v\n", i1, v1)
	}

}
