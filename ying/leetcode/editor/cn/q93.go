package leetcode

import "strings"

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
//和 "192.168@1.1" 是 无效 IP 地址。
//
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你不能重新排序
//或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//
//
// 示例 3：
//
//
//输入：s = "1111"
//输出：["1.1.1.1"]
//
//
// 示例 4：
//
//
//输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
//
//
// 示例 5：
//
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 20
// s 仅由数字组成
//
// Related Topics 字符串 回溯
// 👍 750 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	return solution(s)

}

func solution(s string) []string {
	var ip []string
	var result []string

	var f func(nums, index int, ip []string)

	f = func(nums, index int, ip []string) {
		if nums == 0 {
			ip = []string{}
		}

		if index >= len(s) {
			return
		}

		if nums == 4 && index != len(s)-1 {
			return
		}

		if nums == 4 {
			resultIp := strings.Join(ip, ".")
			result = append(result, resultIp)
			return
		}

		if s[index] == '0' {
			ip = append(ip, string(s[index]))
			f(nums+1, index+1, ip)
		}

		sum := int(s[index]) - 48
		for sum <= 255 {
			ip = append(ip, string(s[index]))
			f(nums+1, index+1, ip)
			index++
			if index >= len(s) {
				return
			}
			sum = sum + int(s[index]-48)
			//if sum > 255 {
			//	return
			//}
		}
		//return
	}

	f(0, 0, ip)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
