// 链接：https://leetcode.com/problems/maximum-product-subarray/
// 题意：给定一个整数数组，求所有子数组中，乘积最大的值是多少 ？

// 输入： [2,3,-2,4]
// 输出： 6
// 解释： [2,3] 的乘积最大，为 6

// 输入： [-2,0,-1]
// 输出： 0
// 解释： [2,3] 的乘积最大，为 6

// 输入： "a good   example"
// 输出： "example good a"

// 思路： DP
//
//		对第 i 个数维护两个值 dpMin[i] 和 dpMax[i] ，
//		分别表示以第 i 个数为结尾第子数组中乘积小值和乘积最大值
//		初始化： dpMin[0] = dpMax[0] = nums[0]
//		状态转移：
//			dpMin[i] = min(nums[i], dpMin[i] * nums[i], dpMax[i] * nums[i])
//			dpMax[i] = min(nums[i], dpMin[i] * nums[i], dpMax[i] * nums[i])
//
//		为了降低空间复杂度，我们可以在更新的时候直接用 dpMin 和 dpMax 表示当前结果
//		每次迭代时，先计算 dpMin[i] * nums[i] 和 dpMax[i] * nums[i] ，然后更新即可
//		注意同时更新结果最大值
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func maxProduct(nums []int) int {
	length := len(nums)
	// 初始化
	dpMin, dpMax := nums[0], nums[0]
	result := dpMax
	// 状态转移
	for i := 1; i < length; i++ {
		a, b := dpMin * nums[i], dpMax * nums[i]
		dpMin = min(nums[i], a, b)
		dpMax = max(nums[i], a, b)
		result = max(result, dpMax)
	}
	// result 即为所有 dpMax 中最大的
	return result
}

func min(firstNum int, remainNums ...int) int {
	for _, num := range remainNums {
		if firstNum > num {
			firstNum = num
		}
	}
	return firstNum
}

func max(firstNum int, remainNums ...int) int {
	for _, num := range remainNums {
		if firstNum < num {
			firstNum = num
		}
	}
	return firstNum
}

