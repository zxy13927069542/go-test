package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	//指定id作为json键值即key
	ID int `json:"id"`
	//key默认与字段名一致
	Name string
	//私有字段无法被json读取，因为私有属性在相同包内被访问
	gender string
}

func main() {
	stu := &Student{
		ID:     3,
		Name:   "郑晓颖",
		gender: "男",
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("序列化错误,错误原因为: %v\n", err)
		return
	}
	fmt.Printf("%s\n", data)
}
