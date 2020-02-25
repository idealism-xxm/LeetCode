// 链接：https://leetcode.com/problems/two-sum/
// 题意：给定一个数组 nums 和一个目标数 target，求和为 target 的两个数的下标
// 思路：用 map 保存每个数出现的下标，遍历数组求目标数的下标，存在则直接返回，不存在则保存入 map 中
// 额外收获：map 的返回值可以有两个，第二个表示键是否存在

// 第一次竟然错了，亏我还运行了一遍，一看错误结果发现是输出了两个数而不是下标，可见正确理解题意及检查很重要
// 跑两次时间差异居然这么大，看来数据集太小了，第一次 8ms，第二次 4ms
func twoSum(nums []int, target int) []int {
	// 保存每个数出现时的 index 的map，由于只有求两个数的和，且题目保证只有一个答案，所以不必记录每个数的下标
	numIndexMap := make(map[int]int)
	for index, num := range nums {
		// 当前 num 需要的另一个数的的下标
		anotherIndex, exists := numIndexMap[target - num]
		// 如果这个数存在，则直接返回
		if exists {
			return []int {anotherIndex, index}
		}
		// 这个数不存在，则标记 num 的位置为 index
		numIndexMap[num] = index
	}
	return nil
}