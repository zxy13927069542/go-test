package json_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

//
//  TestIndent
//  @Description: 测试json.Indent()
//  @param t
//
func TestIndent(t *testing.T) {
	//创建匿名结构体变量
	var person struct {
		Name string
		Age  int
	}
	person.Name = "郑晓颖"
	person.Age = 22

	//序列化person
	jsonStr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Println("Before:")
	fmt.Println(string(jsonStr))

	//格式化json
	var buff bytes.Buffer
	err = json.Indent(&buff, jsonStr, "", "\t")
	fmt.Println("After:")
	fmt.Println(buff.String())
}
