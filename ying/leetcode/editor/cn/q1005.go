package leetcode

import "sort"

//给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
//
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//
//
// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
// 以这种方式修改数组后，返回数组 可能的最大和 。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,2,3], k = 1
//输出：5
//解释：选择下标 1 ，nums 变为 [4,-2,3] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,-1,0,2], k = 3
//输出：6
//解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
//
//
// 示例 3：
//
//
//输入：nums = [2,-3,-1,5,-4], k = 2
//输出：13
//解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 104
// -100 <= nums[i] <= 100
// 1 <= k <= 104
//
// Related Topics 贪心 数组 排序
// 👍 197 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func largestSumAfterKNegations(nums []int, k int) int {
	//排序
	sort.Ints(nums)

	//最大和
	var maxSum int

	var min int
	//全是正数
	if nums[0] >= 0 {
		min = nums[0]
	}

	//全是负数
	if nums[len(nums)-1] < 0 {
		min = -nums[len(nums)-1]
	}

	for i := 0; i < len(nums); i++ {
		if k > 0 {
			if nums[i] < 0 { //负数取反
				//对正负边界判断最小值
				if nums[i] < 0 && i+1 < len(nums) && nums[i+1] >= 0 {
					if -nums[i] < nums[i+1] {
						min = -nums[i]
					} else {
						min = nums[i+1]
					}
				}
				nums[i] = -nums[i]
				k--
			} else if nums[i] >= 0 { //正数取最小值
				if k%2 == 1 {
					//当最小值等于当前元素，且k为奇数时相当于nums[i]变成负数
					if min == nums[i] {
						maxSum -= min
						k = 0
						continue
					} else { //当最小值不等于当前元素，说明有一个负数转换后小于第一个正数
						maxSum -= 2 * min
					}

				}
				k = 0
			}

		}
		maxSum += nums[i]
	}

	//遍历完数组之后k仍不为0
	if k%2 == 1 {
		maxSum -= 2 * min
	}

	return maxSum

}

//leetcode submit region end(Prohibit modification and deletion)
