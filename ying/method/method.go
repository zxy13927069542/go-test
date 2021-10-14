package main

import "fmt"

//结构体Person
//struct是值类型
type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

//Person构造函数
func NewPerson(name, sex string, age int, hobby []string) Person {
	return Person{
		name:  name,
		age:   age,
		sex:   sex,
		hobby: hobby,
	}
}

//方法，类似于java中的成员函数
//p Person指定此方法属于哪种类型，只有该类型能调用此方法
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
//接收者为值类型
func (p Person) GetAge() int {
	return p.age
}

//接收者为指针类型
//结构体为值类型，想要在方法中修改其属性并且保存就需要用到指针类型
func (p *Person) SetAge(age int) {
	p.age = age
}

//非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
//可以使用 type myInt int将int声明为本地类型从而使用方法
type MyInt int

func (m MyInt) Hello() {
	fmt.Println("Hello")
}

func main() {
	p1 := NewPerson("郑晓颖", "男", 18, []string{"打篮球"})
	fmt.Println(p1.GetAge()) //如果p1是值类型，则正常调用，如果p1是指针类型，则相当于(*p1).GetAge()
	p1.SetAge(20)            //如果p1是值类型，则相当于(&p1).SetAge(),如果p1是指针类型，则正常调用
	fmt.Println(p1)          //这即是go语言的语法糖
}
