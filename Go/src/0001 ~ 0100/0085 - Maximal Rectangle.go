// 链接：https://leetcode.com/problems/maximal-rectangle/
// 题意：给定 m * n 的 01 矩阵，求全为 1 的子矩阵的最大面积？

// 输入：
// [
//   ["1","0","1","0","0"],
//   ["1","0","1","1","1"],
//   ["1","1","1","1","1"],
//   ["1","0","0","1","0"]
// ]
// 输出：6

// 思路1：单调栈
//		从第一行开始维护一个数组 h , h[j] 表示第 j 列从当前行网上连续的 1 的数量
//		然后就转化为 0084 这题了
//
//		时间复杂度： O(m * n) ，空间复杂度： O(n)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	h := make([]int, len(matrix[0]))
	result := 0
	for _, row := range matrix {
		for j, item := range row {  // 计算当前行 h 数组的值
			if item == '1' {  // 如果是 1 ，则 + 1
				h[j]++
			} else {  // 否则置零
				h[j] = 0
			}
		}
		result = max(result, largestRectangleArea(h))
	}
	return result
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	heights = append(heights, -1)  // 最后放入 -1 ，让所有数字出栈
	length := len(heights)
	index := make([]int, length)  // 单调递增栈（存储该最小值能到达最左端的下标）
	value := make([]int, length)  // 单调递增栈（存储该最小值）
	top := 0
	index[top], value[top] = 0, -1
	result := 0
	for i, height := range heights {
		if height >= value[top] {  // 当前元素大于等于栈顶元素，值和下标直接入栈
			top++
			index[top], value[top] = i, height
			continue
		}

		for ; height < value[top] ; {  // 当前 栈不空 且 数字小于栈顶数字时，开始出栈
			// 获取栈顶元素的下标作为矩形左边，站定元素高度为矩形高度
			result = max(result, (i - index[top]) * value[top])
			top--  // 出栈
		}

		// 入栈，维持单调递增
		top++
		value[top] = height  // 由于弹出顶元素都 >= height ，所以认为 height 能从当前下标开始
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路2：DP
//		思路和 0084 一样
//		出了维护两个数组： left 和 right ，还需要维护 h
//		h[i] 表示当前行第 i 列的高度，初始均为 0
//		left[i] 表示从 i 开始向左，最后一个大于等于 h[i] 的值的下标，初始均为 0
//		right[j] 表示从 j 开始向右，最后一个大于等于 h[j] 的值的下标，初始均为 n - 1
//		现在是逐行计算，所以当前行 left[i] 未计算之前为前一行以第 i 列的高度能到达的最左下标
//		若：matrix[i][j] == '0'
//			则当前行 left[j] 无意义，初始化为 0 方便下一行进行计算
//		否则：
//			left[j] 为 上一行对应高度能往最左端的下标 和 当前行该列能往最左端的下标 的最大值
//		同理处理 right[j]
//
//		由于需要逐行更新，所以内部循环被优化了，可准确得到
//		时间复杂度： O(n) ，空间复杂度： O(n)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	h := make([]int, n)
	left := make([]int, n)  // 以当前行的高度，最左能够到达的下标
	right := make([]int, n)  // 以当前行的高度，最右能够到达的下标
	for i := 0; i < n; i++ {
		right[i] = n - 1
	}

	result := 0
	for _, row := range matrix {
		rowLeft, rowRight := 0, n - 1  // 当前行
		for i, j := 0, n - 1; i < n; i, j = i + 1, j - 1 {
			if row[i] == '1' {
				h[i]++
				// 若为 1 ，则当前高度下能往最左端的下标等于
				// 上一行对应高度能往最左端的下标 和 当前行该列能往最左端的下标 的最大值
				left[i] = max(left[i], rowLeft)
			} else {
				h[i] = 0
				// 若为 0 ，则当前行 left[i] 无意义，初始化为 0 方便下一行进行计算
				left[i] = 0
				rowLeft = i + 1  // 当前行以后的位置连续 1 最左端最多只能到 i + 1
			}

			if row[j] == '1' {
				// 上一行对应高度能往最右端的下标 和 当前行该列能往最右端的下标 的最小值
				right[j] = min(right[j], rowRight)
			} else {
				// 若为 0 ，则当前行 right[j] 无意义，初始化为 n - 1 方便下一行进行计算
				right[j] = n - 1
				rowRight = j - 1  // 当前行以前的位置连续 1 最右端最多只能到 j - 1
			}
		}
		for i := 0; i < n; i++ {
			result = max(result, (right[i] - left[i] + 1) * h[i])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
