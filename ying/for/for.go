package main

import "fmt"

//打印乘法表
func main() {
	i := 1
	y := 1
	for ; i <= 9; i++ {
		for y = 1; y <= i; y++ {
			fmt.Printf("%d * %d = %d ", y, i, y*i)
		}
		fmt.Print("\n")
	}
}
