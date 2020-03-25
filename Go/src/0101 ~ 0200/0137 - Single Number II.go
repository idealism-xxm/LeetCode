// 链接：https://leetcode.com/problems/single-number-ii/
// 题意：给一个整数数组，有一个数出现一次，其他数均出现两次，找出出现一次的数？
//
// 输入： [2,2,1]
// 输出： 1

// 输入： [4,1,2,1,2]
// 输出： 4

// 思路1： 按位统计
//
//		感觉上还是 0136 的延伸扩展，这题没法使用亦或，但是可以按位进行统计
//		如果某一位上的 1 出现的次数不是 3 的倍数，那么出现一次的数的这一位必定是 1
//		计算每一位即可
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < 32; i++ {
		bit := 1 << i
		count := 0
		for _, num := range nums {
			// 如果这一位不为 0 ，则计数值 + 1
			if num & bit != 0 {
				count++
			}
		}
		// 如果这一位为 1 的数字个数不是 3 的倍数，那么出现一次的数的这一位必定是 1
		if count % 3 != 0 {
			result |= bit
		}
	}

	// 由于输入都是 32 位数，所以要先看成 32 位
	return int(int32(result))
}

// 思路2： 异或
//
//		这题没法直接使用异或是因为无法区分出现 1 次和出现 3 次的情况
//		所以只要能够区分这两种情况，就可以用类似的方法找到答案
//		对于某一位来说，如果我们用 seenOnce, seenTwice 标记这一位出现 1 的次数分别为 1, 2 ：
//			若该位出现 0 次，那么 seenOnce = 0, seenTwice = 0
//			若该位出现 1 次，那么 seenOnce = 1, seenTwice = 0
//			若该位出现 2 次，那么 seenOnce = 0, seenTwice = 1
//		由于我们想排除掉出现 3 次的位，所以出现 3 次时的情况，和出现 0 次的情况一致
//		现在我们就可区分出现 1 次和出现 3 次的情况了
//			当 seenTwice = 0 时，进行 seenOnce  ^ 1 ，这样第 1 次出现时就会标记，而第 2 次出现时就会删除
//			当 seenOnce  = 0 时，进行 seenTwice ^ 1 ，这样第 2 次出现时就会标记，而第 3 次出现时就会删除
//
//		当我们综合所有位考虑的时候，对于每一个数 num 有：
//			seenOnce = ~seenTwice & (seenOnce ^ num)
//			seenTwice = ~seenOnce & (seenTwice ^ num)
//
//		这种方法还可以以此类推，若出现的此处分别为 1 次和 k 次，那么用 k - 1 个遍历标记即可
//		seenI 表示该位第 i 次出现，那么该位第 i 位出现时进行标记，第 i + 1 次出现时删除
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

func singleNumber(nums []int) int {
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}

	return seenOnce
}
