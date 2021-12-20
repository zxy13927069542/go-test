package leetcode

import (
	"math"
	"sort"
)

//冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
//
// 在加热器的加热半径范围内的每个房屋都可以获得供暖。
//
// 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
//
// 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
//
//
//
// 示例 1:
//
//
//输入: houses = [1,2,3], heaters = [2]
//输出: 1
//解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
//
//
// 示例 2:
//
//
//输入: houses = [1,2,3,4], heaters = [1,4]
//输出: 1
//解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
//
//
// 示例 3：
//
//
//输入：houses = [1,5], heaters = [2]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= houses.length, heaters.length <= 3 * 10⁴
// 1 <= houses[i], heaters[i] <= 10⁹
//
// Related Topics 数组 双指针 二分查找 排序 👍 305 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findRadius(houses []int, heaters []int) int {
	return solution1(houses, heaters)
}

//
// solution1
//  @Description: 遍历房子，对每一个房子判断距离最近的供暖器，如果该房子与该供暖器的距离大于初始半径，则将该距离赋值给半径
//  @param houses
//  @param heaters
//  @return int
//
func solution1(houses []int, heaters []int) int {
	sort.Ints(heaters)
	radius := 0
	for i := 0; i < len(houses); i++ {
		//获取离房子最近的供暖器的距离
		distance := binarySearch(heaters, houses[i])
		if distance > radius {
			radius = distance
		}
	}

	return radius
}

//
// binarySearch
//  @Description: 寻找距离房子target最近的供暖器
//  @param heaters
//  @param target
//  @return int
//
func binarySearch(heaters []int, target int) int {
	left := 0
	right := len(heaters) - 1
	for left < right {
		//避免溢出
		mid := int(uint(left+right) >> 1)
		if heaters[mid] == target {
			return 0
		}

		if right-left == 1 {
			if math.Abs(float64(target-heaters[left])) <= math.Abs(float64(heaters[right]-target)) {
				return int(math.Abs(float64(target - heaters[left])))
			} else {
				return int(math.Abs(float64(heaters[right] - target)))
			}
		}

		if heaters[mid] > target {
			right = mid
			continue
		}

		if heaters[mid] < target {
			left = mid
			continue
		}
	}
	return int(math.Abs(float64(heaters[left] - target)))
}

//leetcode submit region end(Prohibit modification and deletion)
