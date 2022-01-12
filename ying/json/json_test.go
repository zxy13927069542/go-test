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

//
// TestUnMarshal
//  @Description: 测试反序列化
//  @param t
//
func TestUnMarshal(t *testing.T) {
	type Person struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	//  测试类型不同能否正确转换
	jsonStr := `{
    "id": 9999,
    "name": "zxy"
	}`

	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("person: %+v\n", person)
}
