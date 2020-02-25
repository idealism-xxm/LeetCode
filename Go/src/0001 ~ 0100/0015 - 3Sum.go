// 链接：https://leetcode.com/problems/3sum/
// 题意：给定一个数组，求所有的三元组，满足三元组的三个数和为 0

// 输入：[-1, 0, 1, 2, -1, -4]
// 输出：[[-1, 0, 1], [-1, -1, 2]]

// 思路1：模拟即可，注意防重（保证结果升序可以很方便地进行防重），也可以改成递归的方式扩展成更多数的和
// 		时间复杂度：O(n^2)

import "sort"

func threeSum(nums []int) [][]int {
	numCntMap := make(map[int]int) // 记录每个数字出现的次数
	for _, num := range nums {
		numCntMap[num]++
	}
	length := len(nums)
	sort.Ints(nums) // 升序排序，保证答案结果升序，同时防重

	var result [][]int
	for i := 0; i < length; {
		if nums[i] > 0 { // 剪枝，由于第一个数最小，所以必定不可能是 正数
			break
		}
		numCntMap[nums[i]]-- // 标记当前数已用一个
		for j := i + 1; j < length; {
			target := 0 - nums[i] - nums[j] // 还需要的数字
			if target < 0 || target < nums[j] { // 剪枝，由于最后一个数最大，所以必定不可能是 负数
				break
			}
			numCntMap[nums[j]]--// 标记当前数已用一个
			// 如果需要的数字大于等于第二个数字（保证升序，便于放重），且还有可用的数字，则满足题意
			if numCntMap[target] > 0 {
				result = append(result, []int {nums[i], nums[j], target})
			}

			numCntMap[nums[j]]++
			for j++; j < length && nums[j - 1] == nums[j]; j++ {} // 相同的第二个数字直接过滤，防重
		}

		numCntMap[nums[i]]++
		for i++; i < length && nums[i - 1] == nums[i]; i++ {} // 相同的第一个数字直接过滤，防重
	}

	return result
}

// 思路2：双指针，不需要都枚举时的 map
// 		排序后，枚举最小的数 num[i]，然后剩下的两个数用双指针寻找，l = i + 1， r = length - 1
//		若 nums[i] + nums[l] + nums[r] < 0， 则 nums[l] 太小，右移 l
//		若 nums[i] + nums[l] + nums[r] > 0， 则 nums[r] 太大，左移 r
//		若 nums[i] + nums[l] + nums[r] == 0，则满足题意，先记录答案，再右移 l（当然左移 r 也行）
//		上述操作均要防重，若移动后的数字等于移动前的数字，则要继续移动，直到数字不等
// 		时间复杂度：O(n^2)

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums) // 升序排序，保证答案结果升序，同时防重

	var result [][]int
	for i := 0; i < length; {
		if nums[i] > 0 { // 剪枝，由于第一个数最小，所以必定不可能是 正数
			break
		}
		l, r := i + 1, length - 1
		for ; l < r; {
			if nums[r] < 0 { // 剪枝，由于最后一个数最大，所以必定不可能是 负数
				break
			}

			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 相同的数字直接过滤，防重
				continue
			}
			if sum > 0 {
				for r--; r > l && nums[r] == nums[r + 1]; r-- {} // 相同的数字直接过滤，防重
				continue
			}
			if sum == 0 {
				result = append(result, []int {nums[i], nums[l], nums[r]})
				for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 相同的数字直接过滤，防重
				continue
			}
		}
		for i++; i < length && nums[i - 1] == nums[i]; i++ {} // 相同的数字直接过滤，防重
	}

	return result
}