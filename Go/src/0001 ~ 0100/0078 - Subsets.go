// 链接：https://leetcode.com/problems/subsets/
// 题意：给定一个数字集合，求所有子集？

// 输入：nums = [1,2,3]
// 输出：
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

// 思路：递归
//		按照顺序递归处理每一个数，每一层可做两个操作
//		1. 放入结果待选列表 2. 不放入结果待选列表
//		等递归到结束条件是，将其放入结果集合中即可
// 		时间复杂度： O(2^n)

func subsets(nums []int) [][]int {
	list := make([]int, len(nums))
	return dfs(nums, 0, list)
}

func dfs(nums []int, count int, list []int) [][]int {
	if len(nums) == 0 {  // 没有需要处理的数，则将现有结果放入集合中
		return [][]int{append(list[:0:0], list[:count]...)}
	}

	// 不选用第一个数
	result := dfs(nums[1:], count, list)
	// 选用第一个数
	list[count] = nums[0]
	result = append(result, dfs(nums[1:], count + 1, list)...)
	return result
}