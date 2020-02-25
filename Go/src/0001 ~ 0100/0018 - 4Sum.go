// 链接：https://leetcode.com/problems/4sum/
// 题意：给定一个整数数组和一个目标数，求 4 个数字和等于目标数的所有组合

// 输入：nums: [1, 0, -1, 0, -2, 2], target = 0
// 输出：[
//  		[-1,  0, 0, 1],
//  		[-2, -1, 1, 2],
//  		[-2,  0, 0, 2]
//		 ]

// 思路1：先排序，保证升序（O(nlog)）
// 		再循环遍历一遍求出每个数字出现的次数（O(n)）
// 		然后枚举所有两个数字和的情况，并记录不同的组合（O(n^2)）
//		最后枚举数字和及其组合，再判断对应剩余数字的组合是否满足即可
//		（不太会算时间复杂度，感觉大部分情况下是 O(n^2)，有人会估计最差情况的复杂度烦请指教）
//		综上：时间复杂度：O(n^2)，空间复杂度：O(n^2)【由于使用了很多复杂的结构，可能效率很低】

import (
	"sort"
)

type Pair struct { // 由于 map 可以用数组做键，故可以用数组，不过定义的 map 语句很复杂，不利于阅读和理解
	A int
	B int
}

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	sort.Ints(nums) // 先排序，用于放重

	numCnt := make(map[int] int)
	for _, num := range nums { // 统计每个数字出现的次数
		numCnt[num]++
	}

	sumPairs := make(map[int] []Pair)
	for i := 0; i < length; {
		for j := i + 1; j < length; {
			sum := nums[i] + nums[j]
			sumPairs[sum] = append(sumPairs[sum], Pair{A: nums[i], B: nums[j]}) // 统计两个数字和可能的组合

			for j++; j < length && nums[j] == nums[j - 1]; j++ {} // 过滤全部相同的数字，防重
		}
		for i++; i < length && nums[i] == nums[i - 1]; i++ {} // 过滤全部相同的数字，防重
	}

	var result [][]int
	exists := make(map[[4]int] bool)
	for sum, pairs := range sumPairs { // 枚举数字和，获取数字和 及 可能的组合
		remainPairs := sumPairs[target - sum] // 获取 剩余数字 的 可能的组合

		for _, remainPair := range remainPairs {  // 枚举 剩余数字 的 各个组合
			numCnt[remainPair.A]-- // 减去 1，表示已使用
			numCnt[remainPair.B]-- // 减去 1，表示已使用
			for _, pair := range pairs { // 枚举 数字和 的 各个组合
				if numCnt[pair.A] <= 0 {
					continue
				}

				numCnt[pair.A]-- // 减去 1，表示已使用
				if numCnt[pair.B] > 0 {  // 若组合的数字都还有，则满足题意
					item := [4]int{remainPair.A, remainPair.B, pair.A, pair.B} // 初始化结果
					sort.Ints(item[0:]) // 结果按升序排序
					if !exists[item] {  // 如果这个结果没出现过，则放入答案列表，并标记已出现
						result = append(result, item[0:])
						exists[item] = true
					}
				}
				numCnt[pair.A]++ // 加上 1，表示不再使用
			}
			numCnt[remainPair.A]++ // 加上 1，表示不再使用
			numCnt[remainPair.B]++ // 加上 1，表示不再使用
		}

		sumPairs[sum] = nil // 已经遍历过，则赋值为 nil，防重，并减少无用循环判断
	}

	return result
}

// 思路2：用递归，可以很容易扩展到 kSum，这也是做 3sum 时的想法
//		当最后变成 2sum 时，用双指针优化，使最终时间复杂度变为：O(n^(k - 1))

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums) // 保证升序，便于放重
	return kSum(nums, target, 4, []int{})
}

// nums：		还可以用的数（保证升序）
// target：		k 个数的目标和
// k：			需要用 k 个数（k >= 2）
// usedNums：	已使用的数
func kSum(nums []int, target int, k int, usedNums []int) [][]int {
	length := len(nums)
	if length < k {
		return [][]int{}
	}

	if k == 2 { // k == 2 时，使用双指针优化
		var result [][]int
		l, r := 0, length - 1

		for ; l < r;  {
			sum := nums[l] + nums[r]
			if sum < target { // nums[l] 太小，l 需要右移
				for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 过滤重复数字，防重
			} else if sum > target { // nums[r] 太大，r 需要左移
				for r--; l < r && nums[r] == nums[r + 1]; r-- {} // 过滤重复数字，防重
			} else { // sum == target，两个数正好，记录答案，既可以 l 右移，也可以 r 左移
				result = append(result, append(usedNums, nums[l], nums[r]))
				for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 过滤重复数字，防重
			}
		}

		return result
	}

	// 当 k > 2 时，枚举当前数，递归处理
	var result [][]int
	for i := 0; i < length; { // 将所有递归处理的结果合并
		result = append(result, kSum(nums[i + 1:], target - nums[i], k - 1, append(usedNums, nums[i]))...)
		for i++; i < length && nums[i - 1] == nums[i]; i++ {} // 过滤重复数字，防重
	}
	return result
}

// 思路3：可以按照 3sum 那个双指针的思路，枚举前两个数，后两个数用双指针确定
// 		排序后，枚举最小的数 num[i] 和次小的数 nums[j]
// 		然后剩下的两个数用双指针寻找，l = j + 1， r = length - 1
//		令 sum = nums[i] + nums[j] + nums[l] + nums[r]
//		若 sum <  0， 则 nums[l] 太小，右移 l
//		若 sum >  0， 则 nums[r] 太大，左移 r
//		若 sum == 0，则满足题意，先记录答案，再右移 l（当然左移 r 也行）
//		上述操作均要防重，若移动后的数字等于移动前的数字，则要继续移动，直到数字不等
// 		时间复杂度：O(n^3)

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	sort.Ints(nums) // 升序排序，保证答案结果升序，同时防重

	var result [][]int
	for i := 0; i < length; {
		for j := i + 1; j < length; {
			l, r := j + 1, length - 1
			for ; l < r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]

				if sum < target {
					for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 相同的数字直接过滤，防重
					continue
				}
				if sum > target {
					for r--; r > l && nums[r] == nums[r + 1]; r-- {} // 相同的数字直接过滤，防重
					continue
				}
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					for l++; l < r && nums[l - 1] == nums[l]; l++ {} // 相同的数字直接过滤，防重
					continue
				}
			}
			for j++; j < length && nums[j - 1] == nums[j]; j++ {} // 相同的数字直接过滤，防重
		}
		for i++; i < length && nums[i - 1] == nums[i]; i++ {} // 相同的数字直接过滤，防重
	}

	return result
}