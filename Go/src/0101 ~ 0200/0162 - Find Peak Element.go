// 链接：https://leetcode.com/problems/find-peak-element/
// 题意：给定一个整数数组，相邻的数不相等，返回一个数组下标，使得其对应的数字比两边都大 ？

// 输入： nums = [1,2,3,1]
// 输出： 2
// 解释： 3 比两边都大，对应的下标为 2

// 输入： nums = [1,2,1,3,5,6,4]
// 输出： 1 或 5
// 解释： 2 和 6 比两边都大，对应的下标为 1 和 5

// 思路： 二分
//
//		题目要求时间复杂度为 O(logn) ，首先肯定就会想到是二分，
//		然后就开始思考如何处理，
//		因为相邻的数不相等，所以可以先这样归纳思考：
//		为什么所有长度的数组都有满足题意的答案？
//		1. 数组长度为 1 时，必定满足题意
//		2. 假设数组长度为 i 时，满足题意，
//			那么转换成长度为 i + 1 的数组就需要在最后添加一个数字，
//			(1) 如果原数组满足题意的数不是最后一个，那么添加任意数字仍满足题意
//			(2) 如果原数组满足题意的数是最后一个：
//				则有 nums[i - 2] < nums[i - 1]
//				此时添加的数字 nums[i] 有两种情况：
//				(a) nums[i - 1] < nums[i] ，则 nums[i] 是满足题意的数
//				(b) nums[i - 1] > nums[i] ，则 nums[i - 1] 是满足题意的数
//
//		思考到这其实就很容易想到如何二分处理了：
//		每次找到中点 mid ，并比较 nums[mid] 和 nums[mid + 1] 的大小
//		1. nums[mid] < nums[mid + 1] ，
//			则有 nums[mid + 1] 左侧已满足题意，
//			根据上述思考可以知道 nums[mid + 1:] 必定有满足题意的数，
//			可以直接处理右半部分即可
//		2. nums[mid] > nums[mid + 1] ，
//			则有 nums[mid] 右侧已满足题意，
//			根据上述思考可以知道 nums[:mid] 必定有满足题意的数，
//			可以直接处理左半部分即可
//		当二分区间只有一个数时，必定是满足题意的数
//
//		时间复杂度： O(logn)
//		空间复杂度： O(1)

func findPeakElement(nums []int) int {
	l, r := 0, len(nums) - 1
	for ; l < r; {
		mid := (l + r) >> 1
		if nums[mid] < nums[mid + 1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
