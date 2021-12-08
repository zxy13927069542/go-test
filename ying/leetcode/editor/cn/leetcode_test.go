package leetcode

import (
	"fmt"
	"testing"
)

func Test1005(t *testing.T) {
	arr := []int{2, -3, -1, 5, -4}
	max := largestSumAfterKNegations(arr, 2)
	fmt.Println(max)
}

func Test1034(t *testing.T) {
	arr := [][]int{
		{1, 2, 1, 2, 1, 2},
		{2, 2, 2, 2, 1, 2},
		{1, 2, 2, 2, 1, 2},
	}
	arr = colorBorder(arr, 1, 3, 1)
	fmt.Println(arr)
}
