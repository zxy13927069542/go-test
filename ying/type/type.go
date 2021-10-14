package main

import "fmt"

//type关键字:自定义类型或者别名或者结构体
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
type myInt int     //自定义类型:编译后仍然有效
type yourInt = int //类型别名:仅在编译时有效，例如字符类型rune是int32的别名
type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

//Person构造函数
func NewPerson(name, sex string, age int, hobby []string) *Person {
	return &Person{
		name:  name,
		age:   age,
		sex:   sex,
		hobby: hobby,
	}
}

//面试题
type student struct {
	name string
	age  int
}

//结构体匿名字段
//字段名默认为类型名,同类型的匿名字段只能有一个,可以与非匿名字段混合使用
//相当于
//var string string
//var int int
type Girl struct {
	string
	int
}

func main() {
	var a myInt = 100
	fmt.Printf("%T\n", a) //main.myInt

	var b yourInt = 100
	fmt.Printf("%T\n", b) //int

	//初始化方式
	//1.基本初始化
	var p Person
	//未赋值时结构体中的属性都为对应的零值
	fmt.Println(p) //{ 0  []}
	p.name = "郑晓颖"
	p.age = 12
	p.sex = "男"
	p.hobby = []string{"打篮球", "打足球"}
	fmt.Println(p)

	//2.使用new进行初始化
	p1 := new(Person)             //返回结构体的地址，p1为指针类型
	fmt.Printf("p1的类型为：%T\n", p1) //p1的类型为：*main.Person
	//亦可直接使用.来进行赋值与访问
	p1.age = 12
	p1.hobby = []string{"打球"}
	p1.name = "郑晓颖"
	p1.sex = "男"
	fmt.Println(p1) //&{郑晓颖 12 男 [打球]}

	//3.使用&取地址符相当于new操作
	p2 := &Person{}
	fmt.Printf("p2的类型为：%T\n", p2) //p2的类型为：*main.Person
	//亦可直接使用.来进行赋值与访问
	p2.age = 12
	p2.hobby = []string{"打球"}
	p2.name = "郑晓颖"
	p2.sex = "男"
	fmt.Println(p2) //&{郑晓颖 12 男 [打球]}

	//4.使用键值对初始化
	p3 := Person{
		age:   16,
		name:  "p3",
		hobby: []string{"打篮球"},
		sex:   "男",
	}
	fmt.Printf("p3的类型为：%T,值为：%v\n", p3, p3) //p3的类型为：main.Person,值为：{p3 16 男 [打篮球]}
	fmt.Println(&p3.name)
	fmt.Printf("%p\n", &p3)

	//匿名结构体
	var boy struct {
		name string
		age  int
	}
	boy.name = "郑晓颖"
	boy.age = 18
	fmt.Println(boy)

	//结构体匿名字段
	g1 := Girl{
		"郑晓颖",
		18,
	}
	fmt.Println(g1)

	//面试题
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		//	娜扎 => 大王八
		//	大王八 => 大王八
		//	小王子 => 大王八
	}
}
