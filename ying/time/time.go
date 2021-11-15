package main

import (
	"fmt"
	"time"
)

/**
测试time包
*/

func main() {
	now := time.Now()
	unix := now.Unix()
	nano := now.UnixNano()
	fmt.Println(unix, nano)
}
