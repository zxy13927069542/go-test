package error

import "fmt"

func ErrorMsg(errorMsg string, err error) {
	fmt.Println("ERROR! ", errorMsg, "错误类型为：", err)
	return
}
