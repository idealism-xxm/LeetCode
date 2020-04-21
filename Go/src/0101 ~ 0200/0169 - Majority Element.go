// 链接：https://leetcode.com/problems/majority-element/
// 题意：给定一个整数数组，有一个数出现的次数超过一半（向下取整），找出这个数？

// 输入： [3,2,3]
// 输出： 3

// 输入： [2,2,1,1,1,2,2]
// 输出： 2

// 思路1： 分治
//		这题很简单，很容易就能想到各种解法：
//			1. 排序后取中间的数：
//				时间复杂度： O(nlogn)
//				空间复杂度： O(1) | O(n) （无法原地排序时为 O(n)）
//			2. map 统计每个数的数字，然后取出现次数最多的即可
//				时间复杂度： O(n)
//				空间复杂度： O(n)
//
//		看了题解后，还知道可以使用分治，
//		其实很多算法都是使用这种思想，通过递归分治降低时间复杂度
//
//		分别处理左右两边的，获得对应的众数 left 和 right
//		1. left == right ：则左右两边都是同一个众数，直接返回即可
//		2. left != right ：则无法确定当前范围内的众数，需要遍历当前区间，
//			统计哪一个数更多即可 （由于题目保证众数出现次数超过一半，
//			所以不断二分的情况下必定每一层都有一个区间是真正的众数）
//
//		时间复杂度： O(nlogn)
//		空间复杂度： O(logn)

func majorityElement(nums []int) int {
	return dfs(nums, 0, len(nums) - 1)
}

func dfs(nums []int, l, r int) int {
	// 区间只有一个数，则其必定是众数
	if l == r {
		return nums[l]
	}

	// 递归找到左右区间的众数
	mid := (l + r) >> 1
	left := dfs(nums, l, mid)
	right := dfs(nums, mid + 1, r)

	// 如果左右区间的众数一样，则其必定是当前区间的众数
	if left == right {
		return left
	}

	leftCount, rightCount := count(nums, l, r, left, right)
	// 当前区间出现次数更多的是众数
	//（不用担心次数相等而出错，因为每一层一定有一个区间是真正的众数，不同时就会再次计算次数）
	if leftCount > rightCount {
		return left
	}
	return right
}

// 统计左右区间的众数出现的次数
func count(nums []int, l, r, left, right int) (leftCount, rightCount int) {
	for ; l <= r; l++ {
		if nums[l] == left {
			leftCount++
		}
		if nums[l] == right {
			rightCount++
		}
	}
	return leftCount, rightCount
}

// 思路2： 按位统计
//		看题解下面的评论还有提到按位统计，这样空间复杂度就可以优化为 O(1)
//		针对每一位看：
//		1. 若众数的该位为 1 ，那么这位为 1 的数字个数必定超过一半
//		2. 若众数的该位为 0 ，那么这位为 1 的数字个数必定不超过一半
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func majorityElement(nums []int) int {
	half := len(nums) >> 1
	result := 0
	for i := 0; i < 32; i++ {
		bit := 1 << i
		count := 0
		for _, num := range nums {
			if (num & bit) != 0 {
				count++
			}
		}
		if count > half {
			result |= bit
		}
	}
	// 传入的是 32 位整数，所以需要先转换成 32 位整数兼容负数的情况
	return int(int32(result))
}

// 思路3： Boyer-Moore 投票算法
//		题解也提到了最优的 Boyer-Moore 投票算法
//		先指定众数 result = random ，并且其出现的次数 count = 0
//		然后遍历整个数组：
//		1. count == 0 时： 令 result = num
//		2. result == num 时： count++
//		3. result != num 时： count--
//
//		算法正确性：
//		1. 由于先判断 count == 0 时，令 result = num ，所以 count 必定时非负数
//		2. 若 result 就是众数，那么下一次 count 为 0 时，必定抵消了相同数量的非众数，
//			剩余的数组中，众数还是占一半以上
//		3. 若 result 不是众数，那么下一次 count 为 0 时，最多抵消了相同数量的众数，
//			剩余的数组中，众数还是占一半以上
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func majorityElement(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if result == num {
			count++
		} else {
			count--
		}
	}
	return result
}
