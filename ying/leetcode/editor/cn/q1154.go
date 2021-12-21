package leetcode

import (
	"strconv"
)

//给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。
//
// 通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。
//
//
//
// 示例 1：
//
//
//输入：date = "2019-01-09"
//输出：9
//
//
// 示例 2：
//
//
//输入：date = "2019-02-10"
//输出：41
//
//
// 示例 3：
//
//
//输入：date = "2003-03-01"
//输出：60
//
//
// 示例 4：
//
//
//输入：date = "2004-03-01"
//输出：61
//
//
//
// 提示：
//
//
// date.length == 10
// date[4] == date[7] == '-'，其他的 date[i] 都是数字
// date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日
//
// Related Topics 数学 字符串 👍 80 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//
// dayOfYear
//  @Description: 构造一个月份字符串，里面的每一个元素代表该月有多少天
// 	闰年二月会多一天，判定方法，年份是400的倍数，或者年份是4的倍数并且不是100的倍数
//  @param date
//  @return int
//
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	// 每年月份天数是固定的
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// 闰年二月多加一天
	if year%400 == 0 || (year%4 == 0 && year&100 != 0) {
		days[1] += 1
	}

	// 得到给定日期前一个月的所有天数
	rtValue := 0
	for i := 1; i < month; i++ {
		rtValue += days[i-1]
	}

	rtValue += day
	return rtValue
}

//leetcode submit region end(Prohibit modification and deletion)
