// 链接：https://leetcode.com/problems/container-with-most-water/
// 题意：给定 n 个非负整数 a1...an，表示有端点分别在 (i, 0), (i, ai) 的线段，求任意两条线段组成的矩形的最大面积

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49

// 思路：双指针，维护两个指针 i, j，分别初始化为 0，n - 1，则其面积为：area = (j - i) * min(height[i], height[j])
//		(1) 如果 height[i] <= height[j]，则只能向右移动 i
//			假设向左移动 j，则面积必定变小，因为：
//			   (j - 1 - i) * min(height[i], height[j - 1])
//			<= (j - 1 - i) * height[i]
//			<  (j - i) * height[i] = area
//		(2) 如果 height[i] >  height[j]，则只能向左移动 j
//			假设向右移动 i，则面积必定变小，因为：
//			   (j - (i + 1)) * min(height[i + 1], height[j])
//			<= (j - (i + 1)) * height[j]
//			<  (j - i) * height[j] = area
//		时间复杂度：O(n)，空间复杂度：O(1)

// 这题求的结果和另一题（https://leetcode.com/problems/largest-rectangle-in-histogram/）类似
// 刚开始一直以为条件都一样，就用了单调栈，结果样例就跑不过（其实想的时候在大脑中跑就感觉不对，有很多反例）
// 后来仔细一对比，结果发现本题的结果只与选择的两条线段有关
// 而另一题是要求连续的高度都大于等于矩形的高度才行
// 还是需要仔细审题，不能凭借经验胡乱猜测题意
func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	result := 0
	for ; l < r;  {
		result = max(result, (r - l) * min(height[l], height[r])) // 更新面积最大值
		if height[l] <= height[r] {
			l += 1
		} else {
			r -= 1
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