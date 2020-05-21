// 链接：https://leetcode.com/problems/read-n-characters-given-read4/
// 题意：给定 read4 函数，可以连续读取 4 个字节，
//		现在使用这个函数完成 read 函数，可以连续读取 n 个字节，
//		并返回读取的字节数（该函数只会被调用一次）？

// 输入： file = "abc", n = 4
// 输出： 3
// 解释： 当执行你的 read 方法后，buf 需要包含 "abc" 。
//		 文件一共 3 个字符，因此返回 3 。
//		 注意 "abc" 是文件的内容，不是 buf 的内容，
//		 buf 是你需要写入结果的目标缓存区。

// 输入： file = "abcde", n = 5
// 输出： 5
// 解释： 当执行你的 rand 方法后，buf 需要包含 "abcde" 。
//		 文件共 5 个字符，因此返回 5 。

// 输入： file = "abcdABCD1234", n = 12
// 输出： 12
// 解释： 当执行你的 rand 方法后，buf 需要包含 "abcdABCD1234" 。
//		 文件一共 12 个字符，因此返回 12 。

// 输入： file = "leetcode", n = 5
// 输出： 5
// 解释： 当执行你的 rand 方法后，buf 需要包含 "leetc" 。
//		 文件中一共 5 个字符，因此返回 5 。

// 思路： 模拟
//
//		直接模拟即可，多次调用 read4 并将字节拷贝至对应的位置，
//		直至读取长度超过 n 或者字符串全部读取完毕
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
	return func(buf []byte, n int) int {
		// 创建一个长度为 4 的 curBuf 用于 read4 读取
		curBuf := make([]byte, 4)
		for i := 0; i < n; i += 4 {
			// 读取接下来的字符，并记录读取长度
			curBufCount := read4(curBuf)
			// 将读取的字符拷贝到 buf 指定位置
			for j := 0; j < curBufCount && i + j < n; j++ {
				buf[i + j] = curBuf[j]
			}
			// 如果本次已将字符串全部读取完毕，则返回读取的长度
			if curBufCount < 4 {
				return min(i + curBufCount, n)
			}
		}
		// 此时达到缓冲区长度限制
		return n
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
