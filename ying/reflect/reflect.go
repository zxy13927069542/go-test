package main

import (
	"fmt"
	"reflect"
)

var A struct {
	name string
}

//使用反射获取空接口类型
//TypeOf(i interface{}) Type
func GetType(v interface{}) string {
	s := fmt.Sprintln(reflect.TypeOf(v))
	return s
}

//获取值reflect.Value
//ValueOf(i interface{}) Value
func GetValue(v interface{}) string {
	value := reflect.ValueOf(v)
	return value.String()
}

//运行中反射修改变量
//func (v Value) Elem() Value
//函数Elem()可以获取指针的值，从而达到修改源变量而不是副本
func change(i interface{}) {
	//IsNil()判断是否为nil
	//IsValid()判断是否为零值
	if !reflect.ValueOf(i).IsNil() && !reflect.ValueOf(i).IsValid() {
		if reflect.TypeOf(i).Kind() == reflect.Int32 {
			value := reflect.ValueOf(i)
			value.Elem().SetInt(300)
		}
	}

}

func main() {
	s := "李大钊"
	fmt.Println(GetType(s))

}
