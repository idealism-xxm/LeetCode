// 链接：https://leetcode.com/problems/pascals-triangle/
// 题意：给定数字 numRows ，返回杨辉三角的前 numRows 层？

// 输入： 5
// 输出：
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

// 思路：DP
//
//		杨辉三角的生成方法就是 DP 的状态转移方程
//
//		时间复杂度： O(numRows ^ 2)
//		空间复杂度： O(numRows ^ 2)

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	// 处理后面第行
	for i := 1; i < numRows; i++ {
		// 第 i 行有 i + 1 个数字
		row := make([]int, i + 1)
		result[i] = row

		// 第一个和最后一个数字为 1
		row[0], row[i] = 1, 1

		// 通过上一行的数字更新本行的数字
		lastRow := result[i - 1]
		for j := 1; j < i; j++ {
			row[j] = lastRow[j - 1] + lastRow[j]
		}
	}
	return result
}
