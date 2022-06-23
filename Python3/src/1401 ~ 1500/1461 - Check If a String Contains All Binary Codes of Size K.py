# 链接：https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
# 题意：给定一个二进制字符串 s ，判断所有长度为 k 的二进制码是否都是 s 的子串？


# 数据限制：
#  1 <= s.length <= 5 * 10 ^ 5
#  s[i] 是 '0' 或 '1'
#  1 <= k <= 20


# 输入： s = "00110110", k = 2
# 输出： true
# 解释： "00", "01", "10" 和 "11" 都是 s 的子串，
#       对应的子串的下标为 0, 1, 3, 2 。

# 输入： s = "0110", k = 1
# 输出： true
# 解释： "0" 和 "1" 都是 s 的子串，对应的子串的下标为 0, 1 。

# 输入： s = "0110", k = 2
# 输出： false
# 解释： "00" 不是 s 的子串。


# 思路： 滑动窗口 + 滚动哈希
#
#      如果已知 s[i:i+k] 对应的值，
#      那么我们可以在 O(1) 内计算出 s[i+1:i+1+k] 对应的值。
#
#      因为 s[i:i+k] 和 s[i+1:i+1+k] 中间的 s[i+1:k] 是一样的，
#      不需要重复计算。
#
#      所以按照滚动哈希的方法即可进行 O(1) 的转移，
#      假设 s[i:i+k] 对应的值为 code ：
#          1. 去除 code 中 s[i] 对应的贡献，即 code -= (s[i] - '0') << (k - 1)
#          2. 向右移动滑动窗口一位，即 code <<= 1
#          3. 加上 s[i+k] 对应的贡献，即 code += (s[i+1:i+1+k] - '0')
#
#      将所有的这些 code 全部放入一个集合中，
#      最后只要集合的大小是 1 << k ，就说明所有长度为 k 的二进制码都是 s 的子串。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 s 中全部 O(n) 个字符
#      空间复杂度：O(2 ^ k)
#          1. 需要维护长度为 O(k) 的二进制码的集合，最差情况下有 O(2 ^ k) 个


class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        # 使用滑动窗口根据 code 计算相邻的下一个 code 的值，放入集合中
        codes: Set[int] = set()
        code: int = 0
        for r in range(len(s)):
            # 向右移动滑动窗口一位，加上 s[r] 对应的贡献
            code = (code << 1) | (ord(s[r]) - ord('0'))
            # 如果滑动窗口长度已经达到 k ，则将 code 加入集合中，
            # 然后移除 s[r - k + 1] 对应的贡献
            if r >= k - 1:
                codes.add(code)
                code -= (ord(s[r - k + 1]) - ord('0')) << (k - 1)

        # 如果最后集合中的 code 数量就是 1 << k ，
        # 那么 s 包含所有长度为 k 的二进制码
        return len(codes) == (1 << k)