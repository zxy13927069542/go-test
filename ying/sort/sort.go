package sort

// RadixSort
//  @Description: 	基数排序
//  @param arr		待排序的数组指针
//  @return []int	排好序的数组
//
func RadixSort(arr []int) []int {
	//存放arr，用来计算什么时候处理完所有位数
	//用来存放每次按位排序后的结果，作为下一次排序的输入
	tmpArr1 := make([]int, len(arr))
	copy(tmpArr1, arr)
	//遍历数组时进行判断，如果有元素取一位后仍不为0，则继续排序
	flag := false
	ten := 1

	for ; !flag; ten *= 10 {
		flag = true
		//每次循环重新生成一个桶
		var buckets [10][]int
		for i := 0; i < len(tmpArr1); i++ {
			bit := (tmpArr1[i] / ten) % 10
			//判断元素取一位后是否为0，不为0则继续排序
			if tmpArr1[i]/(ten*10) != 0 {
				flag = false
			}
			buckets[bit] = append(buckets[bit], tmpArr1[i])
		}
		tmpArr1 = GetArr(&buckets)
	}
	return tmpArr1

}

//
//  GetArr
//  @Description: 	获取桶中的数据，并按从小到大排成一个切片作为输出
//  @param buckets	待获取数据的桶
//  @return []int	排序好的数组
//
func GetArr(buckets *[10][]int) []int {
	var arr []int
	for i := 0; i < 10; i++ {
		for y := 0; y < len(buckets[i]); y++ {
			arr = append(arr, buckets[i][y])
		}
	}
	return arr
}

//
//  BinarySearch
//  @Description: 	二分查找
//  @param arr		从小到大排序好的数组
//  @param target	待查找的元素
//  @return int		返回找到元素的下标，如果元素不存在则返回-1
////
func BinarySearch(arr []int, target int) int {
	if len(arr) < 1 {
		return -1
	}

	index := len(arr) / 2
	//判断目标元素是否刚好在中间
	if arr[index] == target {
		return index

		//只剩一个元素时返回-1
	} else if len(arr) == 1 {
		return -1

		//对左边的数组递归查找
	} else if arr[index] > target {
		return BinarySearch(arr[:(len(arr)/2)], target)

		//对右边的数组递归查找
	} else {
		result := BinarySearch(arr[(len(arr)/2)+1:], target)
		if result == -1 {
			return result
		}
		return index + 1 + result
	}
}

//
// SolutionArr
//  @Description: 通过两重循环后构建两数之和数组，对该数组进行基数排序，
// 遍历数组中的每一个元素，二分查找 m - sum是否存在，存在则返回"YES",不存在则返回"NO"
//  @param arr		小球数组
//  @param m		期待的四球之和
//  @param len		小球数组长度
//  @return string	期待的四球之和则返回"YES"，不存在则返回"NO"
//

func SolutionArr(arr []int, m int, len int) string {
	//构建两数之和数组
	var sums []int
	for i := 0; i < len; i++ {
		for y := 0; y < len; y++ {
			sums = append(sums, arr[i]+arr[y])
		}
	}

	//使用基数排序对数组进行排序
	orderedArr := RadixSort(sums)

	//遍历数组中的每一个元素，二分查找 m - sum是否存在，存在则返回"YES",不存在则返回"NO"
	for i := 0; orderedArr[i] <= m; i++ {
		result := BinarySearch(orderedArr, m-orderedArr[i])
		if result != -1 {
			return "YES"
		}
	}
	return "NO"
}
