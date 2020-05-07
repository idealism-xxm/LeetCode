// 链接：https://leetcode.com/problems/reverse-bits/
// 题意：给定一个 32 位无符号整数，对其二进制位翻转后，返回对应的 32 位无符号整数？

// 输入： 00000010100101000001111010011100
// 输出： 00111001011110000010100101000000
// 解释： 43261596 -> 964176192

// 输入： 11111111111111111111111111111101
// 输出： 10111111111111111111111111111111
// 解释： 4294967293 -> 3221225471

// 思路1： 双指针 + 异或
//
//		和翻转字符串类似，我们维护两个指针 l 和 r ，
//		每次交换第 l 位和第 r 位，
//		由于二进制位只有两种取值，且只有不同的时候才需要交换，
//		所以当该两位不同时，直接改变该两位二进制位即可
//		（使用异或可以很方便将一个二进制位取反）
//
//		时间复杂度： O(1)
//		空间复杂度： O(1)

func reverseBits(num uint32) uint32 {
	// 从两边开始向内处理
	for l, r := 0, 31; l < r; l, r = l + 1, r - 1 {
		// 如果这两位不同，则需要交换（直接改变对应的二进制位即可）
		if ((num >> l) & 1) != ((num >> r) & 1) {
			num ^= 1 << l
			num ^= 1 << r
		}
	}
	return num
}


// 思路2： 分治 + 位运算
//
//		题目有一个进一步的加强，就是这个方法会掉很多次，
//		其实就是想尽可能降低这个方法的时间复杂度
//
//		看了题解才发现可以分治翻转（刚开始还以为是递归分治）
//		思路和递归分治差不多，不过我们能对数字进行位运算，
//		所以每一层的所有翻转都可以批量一起处理
//		从下往上，每一层的需要交换的块逐渐变大（1 > 2 > 4 > 8 > 16）
//		块为 1 时：需要将奇数位和偶数位交换，由于偶数位在奇数位左侧 (1-based)，
//			所以只需要将所有偶数位右移 1 位，再将所有奇数位左移 1 位，
//			再或起来即可
//			即： num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
//			其中： 0xa = 1010
//				  0x5 = 0101
//		块为 2 时：
//			同理可得： num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
//			其中： 0xc = 1100
//				  0x3 = 0011
//		块为 4 时：
//			同理可得： num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
//			其中： 0xf0 = 11110000,
//				  0x0f = 00001111
//		块为 8 时：
//			同理可得： num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
//			其中： 0xff00 = 1111111100000000
//				  0x00ff = 0000000011111111
//		块为 16 时：
//			同理可得： num = ((num & 0xffff0000) >> 16) | ((num & 0x0000ffff) << 16)
//			其中： 0xffff0000 = 11111111111111110000000000000000
//				  0x0000ffff = 00000000000000001111111111111111
//			当然：由于只用处理 32 位无符号整数，所以最后这个不用 mask 也可
//
//		时间复杂度： O(1)
//		空间复杂度： O(1)

func reverseBits(num uint32) uint32 {
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xffff0000) >> 16) | ((num & 0x0000ffff) << 16)
	return num
}
