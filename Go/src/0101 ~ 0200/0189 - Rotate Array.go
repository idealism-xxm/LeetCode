// 链接：https://leetcode.com/problems/rotate-array/
// 题意：给定一个数组，返回其向右移动 k 步后的数组？

// 输入： nums = [1,2,3,4,5,6,7], k = 3
// 输出： [5,6,7,1,2,3,4]
// 解释： 向右移动 1 步： [7,1,2,3,4,5,6]
//		 向右移动 2 步： [6,7,1,2,3,4,5]
//		 向右移动 3 步： [5,6,7,1,2,3,4]

// 输入： nums = [-1,-100,3,99], k = 2
// 输出： [3,99,-1,-100]
// 解释： 向右移动 1 步： [99,-1,-100,3]
//		 向右移动 2 步： [3,99,-1,-100]

// 思路1： 循环移动
//
//		首先将 k 对数组长度取余，获得最少的移动次数，
//		再对 [0, gcd(k, length)) 范围内的数依次右移动 k 步，
//
//		刚开始取的范围是 [0, k) ，发现无法获得正确的结果，
//		仔细一想是最后一个满足 index % k == i 的数字不一定会落在 i 处，
//		应该是所有 (index + x * length) % k == i 的数字形成了一个循环，
//		循环长度可以这么思考：我们从 i 开始，那么循环结束必须要再次回到 i ，
//			即： (i + y * k) % length == i
//			即： i + y * k 能整除 length ，且 y 是最小的那个整数
//			可以知道 gcd(k, length) 一定包含在 length 中，
//			所以只需要 y = length / gcd(k, length) 即可满足 y * k 能整除 length
//		由于只向有移动了 y 次，所以循环长度为 length / gcd(k, length) ，
//		那么我们就需要将范围变为 [0, gcd(k, length)) ，使得整体移动次数为 length
//
//		对于当前 i 来说，是所有 (index + x * length) % k == i 的数字整体循环向右移 k 步，
//		（这些数字在同一个右移循环内）
//		我们定义两个变量 pre 和 j ，分别表示 k 步之前的数字和当前要处理的下标，
//		初始化 pre, j := nums[i], i + k
//		在循环内，我们每次交换 pre 和 nums[j] 的值，
//		保证前一个数右移到当前位置，并且当前位置的数能继续右移动到下一个位置，
//		并不断令 j = (j + k) % length ，直到 j == i
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func rotate(nums []int, k int)  {
	length := len(nums)
	if length <= 1 {
		return
	}

	// 计算最小移动长度
	k %= length
	for i := gcd(k, length) - 1; i >= 0 ; i-- {
		// 对 (index + x * length) % k == i 的所有数字整体循环右移 k 步
		// （这些数字在同一个右移循环内）
		pre, j := nums[i], i + k
		for ; j != i; {
			pre, nums[j] = nums[j], pre
			// j = (j + k) % length
			j += k
			if j >= length {
				j -= length
			}
		}
		// 最后再将数组中最后一个满足 (index + x * length) % k == i 的数字移动到 i 处
		nums[i] = pre
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

// 思路2： 翻转
//
//		看了题解发现还有一种复杂度也是最优，但更简单的方法
//		我们先翻转整个数组，
//		此时 [0, k) 和 [k, length) 内的数都是结果的倒序，
//		所以再对这两个区间进行翻转即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func rotate(nums []int, k int)  {
	length := len(nums)
	if length <= 1 {
		return
	}

	// 计算最小移动长度
	k %= length
	// 翻转整个数组
	reverse(nums, 0, length - 1)
	// 翻转 [0, k)
	reverse(nums, 0, k - 1)
	// 翻转 [k, length)
	reverse(nums, k, length - 1)
}

func reverse(nums []int, l, r int) {
	for ;l < r; l, r = l + 1, r - 1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
