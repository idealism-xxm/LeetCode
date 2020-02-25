// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// 题意：给定一个升序的整型数组，删除多余两个的数字，
//		将剩余的数字都放在原数组前面，并返回剩余数字的个数？

// 输入：nums = [1,1,1,2,2,3]
// 输出：nums = [1,1,2,2,3,any], length = 5

// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：nums = [0,0,1,1,2,3,3,any,any], length = 7

// 思路：双指针
//		维护两个指针，后指针指向当前放入数字的位置，前指针指向待放入的数字
//		并维护当前数字计数，
//		1. 若待放入的数字超过两次，则直接移动前指针
//		2. 否则：放入数字，然后移动前后指针

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	count := 1
	back := 1
	for front := 1; front < length; front++ {
		if nums[front] == nums[front - 1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[back] = nums[front]
			back++
		}
	}
	return back
}
