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

func Test748(t *testing.T) {
	str := "1s3 456"
	words := []string{"looks", "pest", "stew", "show"}
	shortestCompletingWord(str, words)
}

func Test475(t *testing.T) {
	houses := []int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}
	heaters := []int{499, 500, 501}
	fmt.Println(findRadius(houses, heaters))
}
