// 链接：https://leetcode.com/problems/maximum-subarray/
// 题意：给定一个整数数组，求和最大的子数组的和？

// 输入：[-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：子数组 [4,-1,2,1] 的和最大，为 6

// 思路1：DP
//		dp[i] 表示以 nums[i] 为结尾的子数组和的最大值
//		1. 若 dp[i] < 0 ，则 dp[i] = nums[i]
//		2. 若 dp[i] >= 0 ， 则 dp[i] = dp[i - 1] + nums[i]
//		时间复杂度：O(n)
package main

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路2：分治
//		用分治的思想就需要考虑结果会从哪些方面得到范围内最大和：
//		1. 左半部分最大和
//		2. 右半部分最大和
//		3. 左半部分右起最大和 + 右半部分左起最大和
//		所以我们在返回当前范围内最大和时，还要同时返回辅助值：左起最大和、右起最大和、范围和
//		时间复杂度： O(n) ， 空间复杂度： O(n)

func maxSubArray(nums []int) int {
	_, _, _, maxSum := maxRange(nums, 0, len(nums)-1)
	return maxSum
}

func maxRange(nums []int, low, high int) (sum, leftMaxSum, rightMaxSum, maxSum int) {
	if low == high {
		return nums[low], nums[low], nums[low], nums[low]
	}

	mid := (low + high) >> 1
	lSum, lLeftMaxSum, lRightMaxSum, lMaxSum := maxRange(nums, low, mid)
	rSum, rLeftMaxSum, rRightMaxSum, rMaxSum := maxRange(nums, mid+1, high)

	// 范围和 = 左半部分范围和 + 右半部分范围和
	sum = lSum + rSum
	// 左起最大和 = max(左半部分左起最大和，左半部分范围和 + 右半部分左起最大和)
	leftMaxSum = max(lLeftMaxSum, lSum+rLeftMaxSum)
	// 右起最大和 = max(右半部分右起最大和，右半部分范围和 + 左半部分右起最大和)
	rightMaxSum = max(rRightMaxSum, rSum+lRightMaxSum)
	// 范围最大和 = max(左半部分范围最大和，右半部分范围最大和，左半部分右起最大和 + 右半部分左起最大和)
	maxSum = max(lMaxSum, rMaxSum, lRightMaxSum+rLeftMaxSum)
	return
}

func max(firstNum int, remainNums ...int) int {
	for _, num := range remainNums {
		if firstNum < num {
			firstNum = num
		}
	}
	return firstNum
}
