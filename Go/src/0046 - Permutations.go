// 链接：https://leetcode.com/problems/permutations/
// 题意：给定一个不含重复数字的数组，求全排列？

// 输入：[1,2,3]
// 输出：
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

// 思路：递归模拟即可
//		每次枚举未使用的数字，放入排列对应的位置，当所有位置都放入后，即找到一个不重复且合法的排列

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))  // 表示当前数字是否正在使用
	list := make([]int, len(nums))  // 表示当前结果排列
	return dfs(nums, used, 0, list)
}

func dfs(nums []int, used []bool, current int, list []int) [][]int {
    if current == len(nums) {  // 若所有位置都放入数字，则当前排列放入结果列表中
    	return [][]int {append(list[:0:0], list...)}
	}

	var result [][]int  // 收集循环中产生的所有结果
	for i := 0; i < len(nums); i++ {
		if !used[i] {  // 如果当前数字未使用
			used[i] = true  // 使用当前数字
			list[current] = nums[i]  // current 位放入 nums[i]
			result = append(result, dfs(nums, used, current + 1, list)...)  // 收集所有结果
			used[i] = false  // 不再使用当前数字
		}
	}
	return result
}