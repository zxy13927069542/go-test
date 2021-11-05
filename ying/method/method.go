package main

import "fmt"

type Address struct {
	province string
	city     string
}

type Person struct {
	name    string
	age     int
	address Address
}

//嵌套匿名结构体
type Person1 struct {
	name string
	age  int
	Address
}

func main() {
	person := &Person{
		name: "郑晓颖",
		age:  22,
		address: Address{
			province: "广东",
			city:     "珠海",
		},
	}
	fmt.Printf("%v\n", person.address.province)

	//嵌套匿名结构体
	person1 := &Person1{
		//同一层次的字段实例化时，要么全部匿名，要么全部不匿名，否则会报错
		"郑晓颖",
		22,
		Address{
			province: "广东",
			city:     "珠海",
		},
	}
	fmt.Printf("%v\n", person1.Address.province)
	fmt.Printf("%v\n", person1.city)
}
