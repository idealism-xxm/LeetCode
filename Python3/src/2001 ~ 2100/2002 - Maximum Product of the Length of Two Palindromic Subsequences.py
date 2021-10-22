# 链接：https://leetcode.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/
# 题意：给定一个字符串，求不交的两个回文子序列的长度的最大乘积？

# 数据限制：
#   2 <= s.length <= 12
#   s 仅由英文小写字母组成

# 输入： s = "leetcodecom"
# 输出： 9
# 解释：
#   两个回文子序列 "ete" 和 "cdc" ，长度都是 3 ，最大乘积为 9 。

# 输入： s = "bb"
# 输出： 1
# 解释：
#   两个回文子序列 "b" 和 "b" ，长度都是 1 ，最大乘积为 1 。

# 输入： s = "accbcaxxcxx"
# 输出： 25
# 解释：
#   两个回文子序列 "accca" 和 "xxcxx" ，长度都是 5 ，最大乘积为 25 。


# 思路： 状压 DP
#
#       由于字符串长度很小，所以可以想到状态压缩 dp
#       设 cnt[i] 表示 i 中为 1 的下标对应的回文子序列的长度（如果不是回文，则为 0 ），
#       可以在 O(n * (2 ^ n)) 内求出所有结果
#
#       然后我们枚举所有的状态 i ，再枚举所有的状态 j ，
#       如果 dp[i] && dp[j] && (i & 1) > 0 ，则可以用来更新结果
#
#       时间复杂度： O(n * (2 ^ n) + 2 ^ (n + 1)) ≈ O(n * (2 ^ n) + 3 ^ n)
#       空间复杂度： O(2 ^ n)


class Solution:
    def maxProduct(self, s: str) -> int:
        mx = 1 << len(s)
        cnt = [0] * mx
        # 枚举所有状态
        for i in range(1, mx):
            # 子序列拼接成字符串
            cur = "".join(s[j] for j in range(len(s)) if i & (1 << j))
            # 如果是回文串，则记录长度
            if self.is_palindromic(cur):
                cnt[i] = len(cur)

        ans = 0
        # 枚举第一个回文子序列
        for i in range(1, mx):
            # 如果不是回文，则直接处理下一个
            if not cnt[i]:
                continue
            # 枚举第二个回文子序列
            for j in range(i + 1, mx):
                # 如果两者不相交，且第二个也是回文，则更新最大乘积
                if not(i & j) and cnt[j]:
                    ans = max(ans, cnt[i] * cnt[j])
        return ans
    
    def is_palindromic(self, s: str) -> bool:
        # 双指针判断是否为回文串
        l, r = 0, len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return False
            l += 1
            r -= 1
        return True
