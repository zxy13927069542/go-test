package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRadixSort(t *testing.T) {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < 20; i++ {
		value := rand.Intn(1000)
		arr = append(arr, value)
	}
	target := 645
	//arr = append(arr, target)
	fmt.Println(arr)
	sort := RadixSort(arr)
	fmt.Println(sort)

	//测试二分查找
	search := BinarySearch(sort, target)
	fmt.Println(search)
}

func TestSolutionArr(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < 20; i++ {
		value := rand.Intn(100)
		arr = append(arr, value)
	}
	//arr := []int{950, 710, 778, 701, 245, 315, 24, 857, 452, 754,
	//	687, 357, 124, 908, 552, 197, 426, 122, 655, 227}
	fmt.Println(arr)
	//m := arr[0] + arr[3] + arr[7] + arr[9]
	m := 33
	result := SolutionArr(arr, m, len(arr))
	fmt.Println(result)
}
