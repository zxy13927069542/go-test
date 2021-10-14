package main

import "fmt"

func main() {
	//定义一个切片
	s1 := []int{1, 2}
	//往切片追加元素
	// s1[2] = 3	//失败
	fmt.Println(s1)

	//切片为引用类型，未初始化的切片为nil，相当于null
	var s2 []int
	fmt.Println(s2)                                    //[]
	fmt.Println(s2 == nil)                             //true
	fmt.Printf("s2的长度为:%d,容量为:%d\n", len(s2), cap(s2)) //s2的长度为:0,容量为:0
	//初始化过的切片不为nil，即使长度与容量均为0
	s6 := []int{}
	fmt.Println(s6 == nil) //false

	//由数组得到切片
	s3 := [...]int{3, 3, 3}
	s4 := s3[:]
	s4[0] = 1 //数组相应的值也会发生变化
	fmt.Println(s3, s4)

	//make函数
	//可以指定切片的类型，长度，容量
	s5 := make([]int, 5, 10) //s5的类型为[]int，长度为5，容量为10
	fmt.Printf("s5的类型为:%T,长度为:%d,容量为:%d,值为:%v\n", s5, len(s5), cap(s5), s5)

	//切片扩容
	//容量不够情况下直接追加元素 索引越界
	// s7 := []int{1, 2}
	// s1[2] = 3
	// fmt.Println(s7)

	//容量足够下扩容,也无法直接追加元素，索引越界
	// s8 := make([]int,3,4)
	// s8[3] = 1

	//使用append函数扩容,扩容后容量依照策略进行增加
	s9 := make([]int, 3, 3)
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(s9), cap(s9), s9)
	//返回值必须使用s9接受
	s9 = append(s9, 3, 3)
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(s9), cap(s9), s9)

	//扩容后不使用s10接受,而是新定义一个变量来接受
	s10 := make([]int, 3)
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(s10), cap(s10), s10)
	s11 := append(s10, 11)
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(s10), cap(s10), s10) //s10任指向原来的数组
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(s11), cap(s11), s11) //新建了一个数组由s11来指向

	//使用一个切片扩容
	s12 := append(s10, s11...) //...表示拆开
	fmt.Println(s12)
	fmt.Println(s10, s11) //原切片仍然存在

	//无法使用数组扩容
	// s13 := [...]int{3,3}
	// s14 := append(s12,s13...)

	//练习题
	s15 := make([]string, 5, 10)
	for i := 1; i < 10; i++ {
		s15 = append(s15, fmt.Sprint(i))
	}
	fmt.Println(s15) //[     1 2 3 4 5 6 7 8 9]

	//删除切片元素
	s16 := [...]int{3}
	s17 := s16[:]
	s17 = append(s17[0:0], s17[1:]...)
	fmt.Println(s17, len(s17), cap(s17))
	fmt.Println(s16)

}
