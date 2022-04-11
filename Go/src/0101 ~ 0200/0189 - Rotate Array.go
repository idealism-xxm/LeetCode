// 链接：https://leetcode.com/problems/rotate-array/
// 题意：给定一个整型数组，将它原地循环右移 k 次。
//      进阶：
//          1. 用至少 3 种不同的方法完成
//          2. 使用空间复杂度为 O(1) 的原地方法完成

// 数据限制：
//  1 <= nums.length <= 10 ^ 5
//  -(2 ^ 31) <= nums[i] <= 2 ^ 31 - 1
//  0 <= k <= 10 ^ 5

// 输入：nums = [1,2,3,4,5,6,7], k = 3
// 输出：[5,6,7,1,2,3,4]
// 解释：循环右移 1 次： [7,1,2,3,4,5,6]
//      循环右移 2 次： [6,7,1,2,3,4,5]
//      循环右移 3 次： [5,6,7,1,2,3,4]

// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释：循环右移 1 次： [99,-1,-100,3]
//      循环右移 2 次： [3,99,-1,-100]


// 思路2： GCD
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


// 思路3：三次翻转
//
//		可以发现循环右移 k 次后，
//      数组末尾的 k % n 个数字会移动到数组开始，
//      而数组开始的 n - k % n 个数字，则会向右移动 k % n 次。
//
//      如果我们想将数组 nums 整体翻转，
//      则可以使得末尾的 k % n 个数移动至数组开始，
//      不过此时 nums[:k % n] 和 nums[k % n:] 的顺序都是反的，
//      所以还需要分别对 nums[:k % n] 和 nums[k % n:] 再次翻转，
//      这样就能获得循环右移 k 次的结果。
//
//
//		时间复杂度： O(n)
//          1. 需要遍历数组 nums 中的全部 O(n) 个数字
//		空间复杂度： O(1)
//          1. 只需要使用常数个额外变量


func rotate(nums []int, k int)  {
	n := len(nums)
	// 计算最后有多少数字会被移动到数组开始
	k %= n
	// 翻转整个数组
	reverse(nums, 0, n - 1)
	// 翻转前 k 个数字
	reverse(nums, 0, k - 1)
	// 翻转后 n - k 个数字
	reverse(nums, k, n - 1)
}

func reverse(nums []int, l, r int) {
	// 双指针翻转 [l, r] 内的数字
	for l < r {
		// 交换 l 和 r 位置的数字
		nums[l], nums[r] = nums[r], nums[l]
		// l 向右移动一位
		l += 1
		// r 向左移动一位
		r -= 1
	}
}
