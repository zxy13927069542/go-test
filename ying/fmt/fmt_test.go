package main

import (
	"fmt"
	"testing"
)

//
// TestPlaceholder
//  @Description: 测试占位符
//  @param t
//
func TestPlaceholder(t *testing.T) {
	//通用占位符
	type Person struct {
		Name string
		Age int
	}

	person := Person{
		Name: "郑晓颖",
		Age: 22,
	}

	fmt.Printf("person: %v\n", person)	//%v	person: {郑晓颖 22}
	fmt.Printf("person: %#v\n", person)	//%#v	person: main.Person{Name:"郑晓颖", Age:22}
	fmt.Printf("person: %+v\n", person)	//%+v	person: {Name:郑晓颖 Age:22}

	//布尔占位符
	isTrue := true
	fmt.Printf("isTrue: %t\n", isTrue)	//%t	isTrue: true

	//整型占位符
	i := 15
	fmt.Printf("i: %b\n", i)	//%b	i: 1111
	fmt.Printf("i: %c\n", i)	//%c	i: 
	fmt.Printf("i: %d\n", i)	//%d	i: 15
	fmt.Printf("i: %o\n", i)	//%o	i: 17
	fmt.Printf("i: %x\n", i)	//%x	i: f
	fmt.Printf("i: %X\n", i)	//%X	i: F
	fmt.Printf("i: %U\n", i)	//%U	i: U+000F
	fmt.Printf("i: %q\n", i)	//%q	i: '\x0f'

	//浮点数
	f := 12.34
	fmt.Printf("f: %b\n", f)	//%b	f: 6946802425218990p-49
	fmt.Printf("f: %e\n", f)	//%e	f: 1.234000e+01
	fmt.Printf("f: %E\n", f)	//%E	f: 1.234000E+01
	fmt.Printf("f: %f\n", f)	//%f	f: 12.340000
	fmt.Printf("f: %F\n", f)	//%F	f: 12.340000
	fmt.Printf("f: %g\n", f)	//%g	f: 12.34
	fmt.Printf("f: %G\n", f)	//%G	f: 12.34
}
