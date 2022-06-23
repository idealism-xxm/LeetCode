# 链接：https://leetcode.com/problems/delete-operation-for-two-strings/
# 题意：给定两个字符串 word1 和 word2 ，每次可以从 word1 或 word2 中删除一个字符，
#      求最少需要多少次删除操作，才能使 word1 和 word2 相同。


# 数据限制：
#  1 <= word1.length, word2.length <= 500
#  word1 和 word2 仅含有英文小写字母


# 输入： word1 = "sea", word2 = "eat"
# 输出： 2
# 解释： word1 删除 1 个字符 's' ，变为 "ea" ；
#       word2 删除 1 个字符 't' ，变为 "ea" 。

# 输入： word1 = "leetcode", word2 = "etco"
# 输出： 4
# 解释： word1 删除 4 个字符，分别是 'l', 'e', 'd', 'o' ，变为 "etco" ；
#       word2 不删除字符，仍为 "etco" 。


# 思路： DP
#
#      题目要求尽可能少执行删除操作，
#      那么这样处理后的形成的相同字符串必定是 word1 和 word2 的最长公共子序列。
#
#      所以我们只需要求出 word1 和 word2 的最长公共子序列的长度 lcs 即可，
#      令 m = word1.length, n = word2.length ，
#      则最少需要删除操作的次数就是 m - lcs + n - lcs = m + n - 2 * lcs 。
#
#      设 dp[i][j] 为 word1[0..i] 和 word2[0..j] 的最长公共子序列的长度，
#      初始化： dp[i][0] = dp[0][j] = 0 ，表示空串的最长公共子序列为 0
#      状态转移：
#          1. word1[i - 1] == word2[j - 1]: 最后一个字符相同，
#              则状态必定由 dp[i - 1][j - 1] 转移而来，
#              不用考虑 dp[i - 1][j] 和 dp[i][j - 1] ，因为结果不会更优
#          2. word1[i - 1] != word2[j - 1]: 最后一个字符不相同，
#              则状态可由 dp[i - 1][j] 和 dp[i][j - 1] 转移而来
#  
#      那么 dp[m][n] 就是 word1 和 word2 的最长公共子序列的长度，
#      则最少需要操作的次数就是 m + n - 2 * dp[m][n] 。
#
#      当然空间复杂度可以优化为 O(n) ，一般可采用以下方法：
#          1. 滚动数组：如果每一行的状态都是由当前行与上一行的状态转移而来时，
#              就可以使用滚动数组优化。
#              这种方法对状态转移方程限制较小，是常用的优化方法。
#          2. 一维数组 + 临时变量：如果状态 dp[i][j] 仅由 dp[i - 1] 和 dp[i] 中常数个状态转移而来时，
#              可以使用临时变量优先保存上一行的所需状态，以免使用当前行的状态进行转移
#          3. 一维数组 + 倒序转移：如果状态 dp[i][j] 仅由 dp[i - 1][..=j] 中的状态转移而来时，
#              可以使用倒序转移，从后往前计算，以免使用当前行的状态进行转移
#
#      本题只能采用前两种方法中的任意一种进行优化，但题目没有相关要求，
#      为了代码清晰易懂，实际处理时就不使用优化了。
#
#      本题不能使用第 3 种方法的原因是：
#          1. dp[i][j] 依赖上一行的状态 dp[i - 1][j - 1] 和 dp[i - 1][j] ，
#              这要求行状态必须从小到大进行更新，不能倒序
#          2. dp[i][j] 也依赖当前行的状态 dp[i][j - 1] ，
#              这要求当前行的状态也必须从小到大更新，不能倒序
#
#
#      时间复杂度：O(mn)
#          1. 需要遍历全部 O(mn) 个状态
#      空间复杂度：O(mn)
#          1. 需用维护全部 O(mn) 个状态


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m, n = len(word1), len(word2)
        # dp[i][j] 表示 word1[..i] 和 word2[..j] 的最长公共子序列的长度，
        # 初始化 dp[i][0] 和 dp[0][j] 为 0 ，表示空串的最长公共子序列为 0
        dp: List[int] = [[0] * (n + 1) for i in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if word1[i - 1] == word2[j - 1]:
                    # 最后一个字母相同，则状态必定由 dp[i - 1][j - 1] 转移而来，
                    # 不用考虑 dp[i - 1][j] 和 dp[i][j - 1] ，因为结果不会更优
                    dp[i][j] = dp[i - 1][j - 1] + 1
                else:
                    # 最后一个字母不同，则状态可由 dp[i - 1][j] 和 dp[i][j - 1] 转移而来
                    dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

        return m + n - 2 * dp[m][n]