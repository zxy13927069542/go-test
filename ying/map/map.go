package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
生成随机数的value，并按照key的顺序打印键值对
*/
func main() {
	//1.测试printOrderly函数
	//2.初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	//3.初始化map
	s2 := make(map[string]int, 20)
	//4.往map中填充元素
	var key string //将key与value放在for循环外避免重复定义
	var value int
	for i := 0; i < 20; i++ {
		key = fmt.Sprintf("第%02d", i) // %02d : 不足两位的整数用0补足，超出或等于两位的整数正常输出
		value = rand.Intn(20)         //生产0~19的整数
		s2[key] = value
	}
	//5.调用printOrderly函数
	printOrderly(s2)
}

//将map中的元素按序打印
func printOrderly(s2 map[string]int) {
	//1.初始化一个容量为s2长度的切片
	s3 := make([]string, 0, len(s2))
	//2.遍历s2，将key存入切片中
	for k := range s2 {
		s3 = append(s3, k)
	}
	//3.为切片中的key排序
	sort.Strings(s3)
	//4.输出
	for _, v := range s3 {
		fmt.Println(v, s2[v])
	}
}
