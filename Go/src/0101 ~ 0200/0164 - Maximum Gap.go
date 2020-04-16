// 链接：https://leetcode.com/problems/maximum-gap/
// 题意：给定一个整数数组，求排序后相邻两个数最大间距是多少 ？

// 输入： [3,6,9,1]
// 输出： 3
// 解释： 3 和 6 ， 6 和 9 间距均最大，为 3

// 输入： [10]
// 输出： 0
// 解释： 只有一个数，所以是 0

// 思路1： 基数排序
//
//		想了半天，还是认为必须要对比相邻的元素，
//		那么就必须排序，所以就思考 O(n) 的排序方法
//		自然而然就想到了基数排序
//
//		时间复杂度： O(n + d), d = 10
//		空间复杂度： O(n + d), d = 10

func maximumGap(nums []int) int {
	// 基数排序
	radixSort(nums)
	// 遍历对比即可
	result := 0
	for i := len(nums) - 1; i > 0; i-- {
		result = max(result, nums[i] - nums[i - 1])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func radixSort(nums []int) {
	// numbersList[j] 表示当前第 i 位为 j 时的所有数
	numbersList := make([][]int, 10)
	// 默认待排序数组是第一个，方便后续处理
	numbersList[0] = nums
	for i, weight := 0, 1; i < 10; i, weight = i + 1, weight * 10 {
		nextNumbersList := make([][]int, 10)
		// 从上一次中基数最小的开始遍历
		for _, numbers := range numbersList {
			for _, number := range numbers {
				radix := (number / weight) % 10
				nextNumbersList[radix] = append(nextNumbersList[radix], number)
			}
		}
		numbersList = nextNumbersList
	}
	// 收集排序后的结果到 nums
	for i, radix := 0, 0; radix < 10; radix++ {
		for _, number := range numbersList[radix] {
			nums[i] = number
			i++
		}
	}
}

// 思路2： 桶 + 鸽笼原理
//
//		看了题解才发现又有可以继续深挖的部分
//		考虑最差情况： n 个数间距等分，令 diff = max(nums) - min(nums)
//		那么排序后相邻数的间距为 b = ceil(diff + / (n - 1))
//		那么我们让每个桶容纳的范围为 [x, x + b) ，
//		从最小值开始分割成 diff / b + 1 个桶
//		（由于是前闭后开的范围，所以当能够整除的时候，需要多一个桶放入最大数），
//		那么这种情况下每个桶都只有一个数字，
//		且相邻数间的最大间距就转换成了相邻桶中最大值|最小值的间距
//		此时在考虑间距非等分的情况，那么必定有些桶有不止一个值，而有些桶没有值，
//		那么相邻数间的最大间距仍旧能转换成相邻桶中最大值|最小值的间距
//
//		时间复杂度： O(n + b)
//		空间复杂度： O(b)

import "math"

func maximumGap(nums []int) int {
	// 找出最小值和最大值
	minNum, maxNum := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}
	// 如果最小值 >= 最大值，则直接返回（没有数字时取大于）
	if minNum >= maxNum {
		return 0
	}
	// 计算最差情况下相邻数的间距和桶的数量
	length := len(nums)
	diff := maxNum - minNum
	// 区间大小向上取整
	interval := diff / (length - 1)
	if diff % (length - 1) != 0 {
		interval++
	}
	bucketNum := diff / interval + 1
	// 把所有的数放入对应的桶内
	//（直接计算每个桶中的最大值和最小值即可，不用真放入）
	bucketsMinNum := make([]int, bucketNum)
	bucketsMaxNum := make([]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketsMinNum[i] = math.MaxInt32
		bucketsMaxNum[i] = math.MinInt32
	}
	for _, num := range nums {
		index := (num - minNum) / interval
		bucketsMinNum[index] = min(bucketsMinNum[index], num)
		bucketsMaxNum[index] = max(bucketsMaxNum[index], num)
	}
	// 统计桶间的最大间距
	result := 0
	preMaxNum := bucketsMaxNum[0]
	for i := 1; i < bucketNum; i++ {
		// 只有最小值 <= 最大值时，当前桶才有数字
		if bucketsMinNum[i] <= bucketsMaxNum[i] {
			result = max(result, bucketsMinNum[i] - preMaxNum)
			// 更新上一个桶的最大值
			preMaxNum = bucketsMaxNum[i]
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
