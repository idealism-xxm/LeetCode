// 链接：https://leetcode.com/problems/pascals-triangle-ii/
// 题意：给定数字 k ，返回杨辉三角的第 k 层（使用 O(k) 的空间）？

// 输入： 3
// 输出： [1,3,3,1]

// 思路：DP
//
//		思路和 0118 一致，
//		由于只需要最后一层的数据，所以我们完全可以用一维数组，
//		利用原有的数据直接更新即可
//
//		时间复杂度： O(k ^ 2) 【当然可以利用组合数公式直接计算，进一步降低复杂度】
//		空间复杂度： O(k)

func getRow(k int) []int {
	result := make([]int, k + 1)
	result[0] = 1
	// 处理后面的行
	for i := 1; i <= k; i++ {
		// 最后一个数字为 1
		result[i] = 1

		// 通过上一行的数字更新本行的数字
		// 由于后一个需要知道前一个的上一行数字，所以从后往前更新
		for j := i - 1; j > 0; j-- {
			result[j] += result[j - 1]
		}
	}
	return result
}
