// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// 题意：给你一个升序数组 nums ，原地 删除重复出现的数字，
//      使每个数字最多出现 2 次 ，将剩余的数字都放在原数组前面，
//		返回删除后数组的新长度。
//      
//      要求：不使用额外数组，空间复杂度为 O(1)

// 数据限制：
//  1 <= nums.length <= 3 * 10 ^ 4
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4
//  nums 已按升序排序

// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 
//      并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 
//      不需要考虑数组中超出新长度后面的元素。

// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 解释：函数应返回新长度 length = 7, 
//      并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 
//      不需要考虑数组中超出新长度后面的元素。
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
