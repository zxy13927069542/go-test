package leetcode

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
// 1 <= nums.length <= 10⁴
// -100 <= nums[i] <= 100
// 1 <= k <= 10⁴
//
// Related Topics 贪心 数组 排序 👍 171 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//代表当前尚未遍历过数组
var first bool = true

//可能的最大和
var maxSum int

func largestSumAfterKNegations(nums []int, k int) int {

	for i := 0; i < k; i++ {
		//获取当前数组的最小值下标
		min := Min(nums)

		//最小值置反
		nums[min] = -nums[min]

		//判断最小值置反后是否大于最大和
		if maxSum+2*nums[min] > maxSum {
			maxSum = maxSum + 2*nums[min]
		}
	}

	return maxSum

}

//
//  min
//  @Description: 获取当前数组的最小值下标和数组之和
//  @param nums
//  @return min	数组最小值的下标
//	@return	sum	数组元素之和
//
func Min(nums []int) int {
	min := 0
	for i := 0; i < len(nums); i++ {
		//第一次遍历数组时初始化maxSum
		if first {
			maxSum += nums[i]
		}
		if nums[i] < nums[min] {
			min = i
		}
	}

	//表示已遍历过数组
	first = false
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
