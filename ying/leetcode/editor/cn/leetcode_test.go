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
