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
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, l, r int) int {
	// 区间只有一个数，则其必定是众数
	if l == r {
		return nums[l]
	}

	// 递归找到左右区间的众数
	mid := (l + r) >> 1
	left := dfs(nums, l, mid)
	right := dfs(nums, mid+1, r)

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

// 思路2：位运算
//
//		针对每一个二进制：
//		    1. 若众数的该位为 1 ，那么这位为 1 的数字个数必定超过一半
//		    2. 若众数的该位为 0 ，那么这位为 1 的数字个数必定不超过一半
//
//      所以我们维护一个长度为 32 的数组 count ，
//      count[i] 表示所有数中第 i 个二进制位为 1 的数字个数。
//
//      那么最终统计完成后，我们维护 majority 表示众数，
//      遍历所有的二进制位 i ，如果 count[i] > nums.len() / 2 ，
//      则众数的第 i 位是 1 ，执行 majority |= 1 << i
//
//
//      假设 N 为 n 的最大值，这里是 2 ^ 31 - 1
//
//		时间复杂度： O(nlogN)
//          1. 需要遍历全部 O(n) 个数字
//          2. 每个数字都要遍历全部二进制位，可以看作 O(1) ，
//              但实际严格来讲应该是 O(logN)
//		空间复杂度： O(logN)
//          1. 实际上开辟的二进制位空间与 n 有关系，严格来说应该是 O(logN)

func majorityElement(nums []int) int {
	half := len(nums) >> 1
	majority := 0
	for i := 0; i < 32; i++ {
		bit := 1 << i
		count := 0
		for _, num := range nums {
			if (num & bit) != 0 {
				count++
			}
		}
		if count > half {
			majority |= bit
		}
	}
	// 传入的是 32 位整数，所以需要先转换成 32 位整数兼容负数的情况
	return int(int32(majority))
}

// 思路3： Boyer-Moore 投票算法
//
//		题解也提到了最优的 Boyer-Moore 投票算法
//		先指定众数 majority = random ，并且其出现的次数 count = 0
//		然后遍历整个数组：
//		    1. count == 0 时： 令 majority = num
//		    2. majority == num 时： count++
//		    3. majority != num 时： count--
//
//		算法正确性：
//		    1. 由于先判断 count == 0 时，令 majority = num ，
//              所以 count 必定是非负数
//		    2. 若 majority 就是众数，那么下一次 count 为 0 时，
//              必定抵消了相同数量的非众数，
//              剩余的数组中，众数还是占一半以上
//		    3. 若 majority 不是众数，那么下一次 count 为 0 时，
//              最多抵消了相同数量的众数，
//		    	剩余的数组中，众数还是占一半以上
//
//
//		时间复杂度： O(n)
//          1. 只需要遍历全部 O(n) 个数字一次
//		空间复杂度： O(1)
//          1. 只维护 2 个变量，所以空间复杂度为 O(1)

func majorityElement(nums []int) int {
	majority, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if majority == num {
			count++
		} else {
			count--
		}
	}
	return majority
}
