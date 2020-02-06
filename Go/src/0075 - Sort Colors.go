// 链接：https://leetcode.com/problems/sort-colors/
// 题意：给定一个只含有 0, 1, 2 的整型数组，将其按照升序排序？

// 输入：[2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]

// 思路1：计数
//		先遍历统计每种数字的个数，确定每种数字在数组中的起始下标，
//		然后遍历放入对应的位置即可
//		时间复杂度：O(n) ，空间复杂度： O(1)

func sortColors(nums []int)  {
	length := len(nums)
	counts := make([]int, 3)
	for i := 0; i < length; i++ {
		counts[nums[i]]++
	}
	index := 0
	for i := 0; i < 3; i++ {
		for ; counts[i] > 0; counts[i]-- {
			nums[index] = i
			index++
		}
	}
}

// 思路2：三路快排
//		设置 pivot1 = pivot2 = 1
//		则进行一次 partition 后，数组元素大小关系如下：
//		| nums[i] < pivot1 | pivot1 <= nums[i] <= pivot2 | pivot2 < nums[i]
//		由于只有三个数，则三块区域恰好分别是 0, 1, 2 ，只遍历一次即完成了排序
//		时间复杂度：O(n) ，空间复杂度： O(1)
func sortColors(nums []int)  {
	// l 以左都都是 0 ， r 以右的都是 2
	l, r := 0, len(nums) - 1
	for cur := 0; cur <= r; {
		if nums[cur] == 0 {  // 如果当前数是 0，则将其放入最左边
			// 由于 l <= cur ，则必有 nums[l] == 0 || nums[l] == 1
			// 且 l < cur 时，必有 nums[l] == 1
			nums[l], nums[cur] = nums[cur], nums[l]
			l++
			cur++
		} else if nums[cur] == 2 {  // 如果当前数是 2，则将其放入最右边
			nums[cur], nums[r] = nums[r], nums[cur]
			r--
		} else {  // 如果当前数是 1，则处理直接下一个数
			cur++
		}
	}
}
