# 链接：https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
# 题意：给定一个 01 串，求至少使用多少次操作 2 可以将其变为 01 交错的字符串？
#       操作 1: 将第一个字符放在末尾
#       操作 2: 将一个 0 变为 1 或将一个 1 变为 0

# 输入： s = "111000"
# 输出： 2
# 解释： "111000" -> "101000" -> "101010"

# 输入： s = "010"
# 输出： 0

# 输入： s = "1110"
# 输出： 1
# 解释： "1110" -> "1010"

# 思路： 滑动窗口
#
#       刚开始没有思路，靠猜测做错了两次，然后突然灵光一闪，
#       如果我们将字符串再拼接一次，每次只关注窗口大小为 n 的一部分，
#       这样就没有操作 1 了，统计所有窗口中使用操作 2 次数最小的即可
#
#       注意最终结果有两种情况： 0 开始的串和 1 开始的串
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def minFlips(self, s: str) -> int:
        n = len(s)
        # 翻转次数最多不超过 n // 2
        ans = n
        # 转换成以 0 和 1 开始的字符串的窗口中的操作 2 次数
        ans_0, ans_1 = 0, 0
        s += s
        for i in range(len(s)):
            # 先计算以 0 开始的字符串，当前位置所需的字符
            ch_0 = '0' if i % 2 == 0 else '1'
            # 如果当前字符与 ch_0 不同，那么 0 开始的字符串需要进行一次操作 2
            if ch_0 != s[i]:
                ans_0 += 1
            else:
                # 此时当前字符与 ch_0 相同，则与 ch_1 不同，
                # 那么 1 开始的字符串需要进行一次操作 2
                ans_1 += 1

            # 当超过窗口长度 n 时，需要去除被移除字符带来的操作
            if i >= n:
                # 先计算以 0 开始的字符串，移除字符位置所需的字符
                ch_0 = '0' if (i - n) % 2 == 0 else '1'
                # 如果移除字符与 ch_0 不同，那么 0 开始的字符串需要进行一次操作 2
                if ch_0 != s[i]:
                    ans_0 -= 1
                else:
                    # 此时当前字符与 ch_0 相同，则与 ch_1 不同，
                    # 那么 1 开始的字符串需要进行一次操作 2
                    ans_1 -= 1

                # 此时窗口长度为 n ，可以更新答案
                ans = min(ans, ans_0, ans_1)

        return ans
