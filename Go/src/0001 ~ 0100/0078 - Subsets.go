// 链接：https://leetcode.com/problems/subsets/
// 题意：给定一个数字集合，求所有子集？


// 数据限制：
//  1 <= nums.length <= 10
//  -10 <= nums[i] <= 10
//  nums 中的数字各不相同


// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 输入：nums = [0]
// 输出：[[],[0]]


// 思路：回溯
//
//		按照顺序递归处理每一个数，每一层可做两个操作：
//		    1. 放入结果待选列表 list
//          2. 不放入结果待选列表 list
//
//      递归终止条件是：所有数字都已枚举完毕，
//      这时候结果待选列表 list 就是一种合法的子集
//
//      时间复杂度：O(2 ^ n) 。每个数字都要进行两种选择，总共有 n 个数字要这样处理。
//      空间复杂度：O(n * 2 ^ n) 。总共有 O(2 ^ n) 个子集，每个子集最长为 O(n) 。
//          实际计算所有子集的元素个数的公式为： sum(i * C(n, i)) = n * 2 ^ (n - 1)

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