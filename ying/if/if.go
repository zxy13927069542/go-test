package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "小龙"
	//判断name前缀是否为李
	if strings.HasPrefix(name, "李") {
		fmt.Println("姓李")
	} else {
		fmt.Println("不姓李")
	}
}
