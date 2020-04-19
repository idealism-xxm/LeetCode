// 链接：https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// 题意：给定一个升序整数数组，存在两个数字加起来为 target （唯一），
//		返回这两个数字的下标（每个数字只能用一次，且下标从 1 开始） ？

// 输入： numbers = [2,7,11,15], target = 9
// 输出： [1, 2]

// 思路： 双指针
//
//		维护左右指针 l 和 r ，每次求和 sum = numbers[l] + numbers[r]
//		然后比较 sum 和 target 的大小：
//		1. sum == target: 找到结果，直接返回对应的下标即可
//		2. sum <  target: 和更小，需要一个更大的数，移动左指针
//		3. sum >  target: 和更大，需要一个更小的数，移动右指针
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers) - 1; l < r; {
		sum := numbers[l] + numbers[r]
		// 找到结果，直接返回
		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum < target {
			// 和更小，需要一个更大的数，移动左指针
			l++
		} else {
			// 和更大，需要一个更小的数，移动右指针
			r--
		}
	}
	// 按照题意不可能走到这
	return nil
}
