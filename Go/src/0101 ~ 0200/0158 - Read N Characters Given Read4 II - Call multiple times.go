// 链接：https://leetcode.com/problems/read-n-characters-given-read4/
// 题意：给定 read4 函数，可以连续读取 4 个字节，
//		现在使用这个函数完成 read 函数，可以连续读取 n 个字节，
//		并返回读取的字节数（该函数会被调用多次）？

// File file("abc");
// Solution sol;
// // 假定 buf 已经被分配了内存，并且有足够的空间来存储文件中的所有字符。
// sol.read(buf, 1); // 当调用了您的 read 方法后，buf 需要包含 "a"。 一共读取 1 个字符，因此返回 1。
// sol.read(buf, 2); // 现在 buf 需要包含 "bc"。一共读取 2 个字符，因此返回 2。
// sol.read(buf, 1); // 由于已经到达了文件末尾，没有更多的字符可以读取，因此返回 0。

// File file("abc");
// Solution sol;
// sol.read(buf, 4);
// // 当调用了您的 read 方法后，buf 需要包含 "abc" 。
// // 一共只能读取 3 个字符，因此返回 3。
// sol.read(buf, 1);
// 由于已经到达了文件末尾，没有更多的字符可以读取，因此返回 0 。


// 思路： 模拟
//
//		直接模拟即可，多次调用 read4 并将字节拷贝至对应的位置，
//		直至读取长度超过 n 或者字符串全部读取完毕
//		由于需要调用多次，所以需要在闭包内定义三个变量：
//			1. 上次 read4 读取的 buf 数组
//			2. 上次 read4 读取的长度
//			2. buf 数组中已经读取的长度
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	// 创建一个长度为 4 的 preBuf 用于 read4 读取
	preBuf := make([]byte, 4)
	// 上次读取的 buf 长度
	preBufCount := 0
	// 其中已使用的长度
	usedLen := 0
	return func(buf []byte, n int) int {
		// 记录本次已经复制到 buf 中的字节数
		count := 0
		for {
			// 将上次读取中未使用的字节拷贝到 buf
			for ; count < n && usedLen < preBufCount; count, usedLen = count + 1, usedLen + 1 {
				buf[count] = preBuf[usedLen]
			}
			// 如果已经读够 n 个字节，则直接返回
			if count == n {
				return n
			}
			// 不够 n 个字节，则继续调用 read4 读取
			preBufCount = read4(preBuf)
			usedLen = 0
			// 如果文件没有更多字符了，则直接返回当前已读的数量
			if preBufCount == 0 {
				return count
			}
		}
	}
}
