// 链接：https://leetcode.com/problems/largest-rectangle-in-histogram/
// 题意：给定一个非负数整型数组，表示柱状图每处的高度，求能形成的最大矩形的面积？

// 输入：[2,1,5,6,2,3]
// 输出：10

// 思路1：单调栈
//		单调栈入门题目
//		维护一个单调递增栈
//		1. 若当前数字能够是栈维持单调递增，则直接入栈
//		2. 否则，取前一个元素作为矩形右边（因为当前元素之前的元素都满足大于等于栈顶），
//			开始出栈，并不停以出栈元素为矩形左边，更新矩形最大面积，
//			直至当前数字大于等于栈顶元素对应都数字，然后当前元素入栈
//
//		为了方便操作，开始在栈中放入高度 -1 ，在入参后加上 -1 ，最后让所有元素出栈
//		时间复杂度： O(n) ，空间复杂度： O(n)

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
//		看到还可以用 DP 做，基本想到了，但是实现时还是卡壳，最终看了具体思路
//		维护两个数组： left 和 right
//		left[i] 表示从 i 开始向左，最后一个大于等于 heights[i] 的值的下标
//		right[j] 表示从 j 开始向右，最后一个大于等于 heights[j] 的值的下标
//		每次初始化： left[i] = i
//		若：heights[left[i] - 1] >= heights[i]
//			则：i 对应的最左端至少为 left[i] - 1 对应的最左端
//			即： left[i] = left[left[i - 1]]
//		同理处理 right[j]
//
//		虽然每次内部都有一个循环，但一段连续的区间最多只会被遍历一次
//		（不太会证明，画图寻找最差情况可以发现，不知道有没有特例）
//		时间复杂度： O(n) ，空间复杂度： O(n)

func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}

	left := make([]int, length)  // left[i] 表示从 i 开始向左，最后一个大于等于 heights[i] 的值的下标
	right := make([]int, length)  // right[j] 表示从 j 开始向右，最后一个大于等于 heights[j] 的值的下标
	for i, j := 0, length - 1; i < length; i, j = i + 1, j - 1 {
		left[i], right[j] = i, j  // 初始最左最右端为当前位置
		for ; left[i] > 0 && heights[left[i] - 1] >= heights[i]; {  // 更新最左端
			left[i] = left[left[i] - 1]
		}
		for ; right[j] < length - 1 && heights[right[j] + 1] >= heights[j]; {  // 更新最左端
			right[j] = right[right[j] + 1]
		}
	}

	result := 0
	for i, height := range heights {
		result = max(result, (right[i] - left[i] + 1) * height)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
