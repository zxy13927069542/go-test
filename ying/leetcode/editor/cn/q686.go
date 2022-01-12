package leetcode

//给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
//
// 注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
//
//
//
// 示例 1：
//
// 输入：a = "abcd", b = "cdabcdab"
//输出：3
//解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
//
//
// 示例 2：
//
// 输入：a = "a", b = "aa"
//输出：2
//
//
// 示例 3：
//
// 输入：a = "a", b = "a"
//输出：1
//
//
// 示例 4：
//
// 输入：a = "abc", b = "wxyz"
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= a.length <= 10⁴
// 1 <= b.length <= 10⁴
// a 和 b 由小写英文字母组成
//
// Related Topics 字符串 字符串匹配 👍 170 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedStringMatch(a string, b string) int {
	return circulArr(a, b)
}

//
// circulArr
//  @Description: 循环数组
//  @param a
//  @param b
//  @return int
//
func circulArr(a string, b string) int {
	lenA := len(a)
	lenB := len(b)

	//  记录b[0]元素在a数组中的每一个坐标
	var startArr []int

	//  记录a数组需要重叠几次
	var nums int

	// 记录起始元素在字符串a中的所有位置
	for i := 0; i < lenA; i++ {
		if a[i] == b[0] {
			startArr = append(startArr, i)
		}
	}

	for i := 0; i < len(startArr); i++ {
		//  初始化重叠次数
		nums = 0
		flag := true

		//  循环数组判断
		for y := 0; y < lenB; y++ {
			index := (startArr[i] + y) % lenA

			// 不符合，跳过,从新的起始坐标开始
			if a[index] != b[y] {
				flag = false
				break
			}

			//  计算重叠次数
			if index == lenA-1 && y+1 < lenB && b[y+1] == a[0] {
				nums++
			}
		}

		if flag {
			return nums + 1
		}
	}

	//  字符串中不包含b[0],直接返回-1
	return -1

}

//func subStr(a string, b string) int {
//	str := ""
//	nums := 0
//	for  {
//		str = str + a
//		nums++
//		if strings.Index(str, a) != -1 {
//			return nums
//		}
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
