package main

import "fmt"

func main() {
	var maptest map[int]string
	maptest = make(map[int]string)
	maptest[0] = "郑晓颖"
	fmt.Println(maptest)

}
