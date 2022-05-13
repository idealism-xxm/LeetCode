# 链接：https://leetcode.com/problems/count-sorted-vowel-strings/
# 题意：计算满足以下三个条件的字符串：
#          1. 长度为 n
#          2. 仅包含元音字母 (a, e, i, o, u)
#          3. 字符串内的字母是字典序升序的


# 数据限制：
#  1 <= n <= 50 


# 输入： n = 1
# 输出： 5
# 解释： 总共有 5 个满足题意的字符串：
#       "a", "e", "i", "o", "u"

# 输入： n = 2
# 输出： 15
# 解释： 总共有 15 个满足题意的字符串： 
#       "aa", "ae", "ai", "ao", "au", 
#       "ee", "ei", "eo", "eu", "ii", 
#       "io", "iu", "oo", "ou", "uu"


# 思路： DP
#
#      设 dp[i][j] 表示长度为 i 以第 j 个元音字母结尾，
#      且字典序升序的字符串的数量。
#
#      初始化： dp[1][0~5] = 1
#      状态转移：
#          1. 以 a 结尾，则倒数第二个字母也只能是 a ，即：
#              dp[i][0] = dp[i - 1][0]
#          2. 以 e 结尾，则倒数第二个字母可以是 a 和 e ，即：
#              dp[i][1] = dp[i - 1][0] + dp[i - 1][1]
#                       = dp[i][j - 1] + dp[i - 1][1]
#          3. 以 i 结尾，则倒数第二个字母可以是 a, e, i ，即：
#              dp[i][2] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2]
#                       = dp[i][j - 1] + dp[i - 1][2]
#          4. 以 o 结尾，则倒数第二个字母可以是 a, e, i, o ，即：
#              dp[i][3] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3]
#                       = dp[i][j - 1] + dp[i - 1][3]
#          5. 以 u 结尾，则倒数第二个字母可以是 a, e, i, o, u ，即：
#              dp[i][4] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4]
#                       = dp[i][j - 1] + dp[i - 1][4]
#
#      综上，状态转移方程可以简化为：
#          1. j == 0: dp[i][0] = dp[i - 1][0]
#          2. j != 0: dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
#
#       最终 sum(dp[i]) 就是所有满足题意的字符串数量。​
#
#      时间复杂度：O(n)
#          1. 需要遍历二维数组 dp 的全部 O(n) 个状态
#      空间复杂度：O(n)
#          1. 需要维护一个大小为 O(n) 的二维数组 dp 
#          2. （因为 dp[i][j] 仅与 dp[i - 1][j] 和 dp[i][j - 1] 有关，
#              所以可以使用一维数组优化为 O(1)）


class Solution:
    def countVowelStrings(self, n: int) -> int:
        # dp[i][j] 表示长度为 i 以第 j 个元音字母结尾，
        # 且字典序升序的字符串的数量。
        # 初始化 dp[1] = [1, 1, 1, 1, 1]
        dp = [[1, 1, 1, 1, 1] for _ in range(n + 1)]
        for i in range(2, n + 1):
            # 以 a 结尾，则倒数第二个字母也只能是 a
            dp[i][0] = dp[i - 1][0]
            # 以第 j 个元音字母结尾，则倒数第二个字母只能是第 0~j 个元音字母
            for j in range(1, 5):
                dp[i][j] = dp[i][j - 1] + dp[i - 1][j]

        # 最后返回长度为 n 且以不同元音字母结尾的字符串的数量之和
        return dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]


# 思路2： 排列组合
#
#      想到了高中的排列组合，特别是和挡板法非常相似。
#
#      可以按照类似的思路处理将题目转化为：
#      将一个长度为 n 的字符串分成 5 组，
#      每组字符分别由 a, e, i, o, u 组成，
#      且每一组都可以为空串。
#
#      这时和挡板法区别就在于：挡板法要求每组元素非空，
#      而本题中的每组可以为空，那么可以转换一下思路来处理。
#
#      先看这样的一个具体例子：
#      假设 n = 3 ，用 x 代表字符，用 | 代表挡板，
#      其中一种合法方式为： x||x|x| ，
#      这时不同组中的字符通过挡板已经确定，
#      这个合法方式代表字符串为 aio 。
#
#      可以发现，合法的方式必定有 n + 4 个位置，
#      且其中 n 个位置为字符， 4 个位置为挡板。
#
#      每一种合法方式都唯一对应一种合法的字符串，
#      那么合法的方式数量可以这样计算：
#      从 n + 4 个位置中，选取 4 个位置放置挡板，
#      对应的数量为：
#      C(n + 4, 4) = (n + 4)! / (n! * 4!)
#                  = (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
#
#
#      时间复杂度：O(1)
#          1. 直接按照公式计算即可，时间复杂度为 O(1)
#      空间复杂度：O(1)
#          1. 只需要常数个额外变量


class Solution:
    def countVowelStrings(self, n: int) -> int:
        # C(n + 4, 4) = (n + 4)! / (n! * 4!)
        #             = (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
        return (n + 4) * (n + 3) * (n + 2) * (n + 1) // 24
