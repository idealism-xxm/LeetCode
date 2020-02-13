// 链接：https://leetcode.com/problems/subsets-ii/
// 题意：给定一个数字列表（可能存在重复数字），求所有子集？

// 输入：nums = [1,2,2]
// 输出：
// [
//   [2],
//   [1],
//   [1,2,2],
//   [2,2],
//   [1,2],
//   []
// ]

// 思路：递归
//		思路和 0078 基本一样，只需要修改递归部分逻辑即可
//
//		统计每个数出现的次数，然后初始化一个新的待选数组，保证每个元素各不相同
//		按照顺序递归处理每一个数，每一层可做两个操作
//		1. 放入结果待选列表 (1 ~ numToCount[num]) 次 2. 不放入结果待选列表
//		等递归到结束条件是，将其放入结果集合中即可
// 		时间复杂度： O(2^n)

func subsetsWithDup(nums []int) [][]int {
	list := make([]int, len(nums))

	// 统计每个数字出现次数
	numToCount := make(map[int]int)
	for _, num := range nums {
		numToCount[num]++
	}
	// 初始化一个新的待选数组，每个数字都不同
	distinctNums := make([]int, len(numToCount))
	i := 0
	for num := range numToCount {
		distinctNums[i] = num
		i++
	}

	return dfs(distinctNums, numToCount,0, list)
}

func dfs(distinctNums []int, numToCount map[int]int, count int, list []int) [][]int {
	if len(distinctNums) == 0 {  // 没有需要处理的数，则将现有结果放入集合中
		return [][]int{append(list[:0:0], list[:count]...)}
	}

	// 不选用第一个数
	result := dfs(distinctNums[1:], numToCount, count, list)
	// 选用第一个数 1 ～ numToCount[distinctNums[0]] 次
	num := distinctNums[0]
	for i := 0; i < numToCount[num]; i++ {
		list[count + i] = distinctNums[0]
		result = append(result, dfs(distinctNums[1:], numToCount, count + i + 1, list)...)
	}
	return result
}