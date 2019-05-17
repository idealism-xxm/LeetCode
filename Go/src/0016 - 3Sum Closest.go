// 链接：https://leetcode.com/problems/3sum-closest/
// 题意：给定一个数组 和 一个整数 target，求一个三元组的和，使这三个数的和最接近 target

// 输入：[-1, 2, 1, -4], 1
// 输出：2

// 思路：双指针（有了上一题的经验，这题就轻车熟路了，好像也没有其他比较快的解法了）
// 		排序后，枚举最小的数 num[i]，然后剩下的两个数用双指针寻找，l = i + 1， r = length - 1
//		令 sum = nums[i] + nums[l] + nums[r]
//		每次比较 sum 和 result 与 target 的差的绝对值，若 sum 的更小，则更新 result = sum
//		再根据 sum 和 target 的大小，移动指针
//		若 sum <  target，则 nums[l] 太小，右移 l
//		若 sum >  target，则 nums[r] 太大，左移 r
//		若 sum == target，则可以不用继续更新，直接返回结果
// 		时间复杂度：O(n^2)

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums) // 升序排序，保证答案结果升序，放便使用双指针

	result := math.MaxInt32
	for i := 0; i < length; {
		l, r := i + 1, length - 1
		for ; l < r; {
			sum := nums[i] + nums[l] + nums[r]

			if abs(sum - target) < abs(result - target) {
				result = sum
			}

			if sum < target {
				for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 相同的数字直接过滤
				continue
			}
			if sum > target {
				for r--; r > l && nums[r] == nums[r + 1]; r-- {} // 相同的数字直接过滤
				continue
			}
			if sum == target {
				return result
			}
		}
		for i++; i < length && nums[i - 1] == nums[i]; i++ {} // 相同的数字直接过滤，防重
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}