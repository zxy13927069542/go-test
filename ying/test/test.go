package main

import (
	"fmt"
	"strings"
)

func main() {
	//°: 2个字节 ′: 三个字节 ″: 三个字节
	s := `300°330′330″`

	//验证这几个字符占用几个字节
	//bytes := []byte(s)
	//fmt.Println(bytes)

	index0 := strings.Index(s, "°")
	index1 := strings.Index(s, "′") - 1
	index2 := strings.Index(s, "″") - 1 - 2
	fmt.Println(index0, index1, index2)
}
