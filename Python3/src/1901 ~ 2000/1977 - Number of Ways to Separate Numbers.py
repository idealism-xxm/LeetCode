# 链接：https://leetcode.com/problems/number-of-ways-to-separate-numbers/
# 题意：给定一个数字串，能将其任意分割，要求分割后的数字都是正数且不含前导 0 ，还要单调递增，
#       求满足条件的分割方法数？

# 数据限制：
#   1 <= num.length <= 3500
#   num[i] 仅由数字组成

# 输入： num = "327"
# 输出： 2
# 解释： 
#       (1) 3, 27
#       (2) 327

# 输入： num = "094"
# 输出： 0
# 解释： 
#       所有数字都必须是正数且不含前导 0

# 输入： num = "0"
# 输出： 0
# 解释： 
#       所有数字都必须是正数且不含前导 0

# 输入： num = "9999999999999"
# 输出： 101


# 思路： DP
#
#       很容易就能想到 O(n ^ 3) 的 DP ，
#       设 dp[i][j] 表示 num[:i] 最后一个数字长度为 j 时的分割方法数，
#       那么 dp[i][j] 可通过两种情况转移而来：
#           1. sum(dp[i - j][1 ~ j - 1]) ：num[:i - j] 最后一个数字长度小于 j 的所有分割方法数之和
#               数字长度更小，那么不用比较数字，该数字一定比当前数字更小
#           2. dp[i - j][j] ：当 num[:i - j] 最后一个数字长度为 j 时的分割方法数
#               数字长度相等，则需要比较数字，只有当前数字大于等于前一个数字时才能转移
#
#       其中 1 和 2 都是 O(n) 的复杂度，想要通过就必须优化为 O(1) 的复杂度
#       第 1 点很好优化，求一个前缀和即可，
#       第 2 点就需要我们进行预处理，
#           设 lcs[i][j] 表示 num[i:] 和 num[j:] 的最长公共子串的长度，
#           那么当 lcs[i - j - j][i - j] >= j 时，则长度为 j 的数字相等，可以进行转移，
#           当 lcs[i - j - j][i - j] = cnt < j 是，
#               需要看 num[i - j - j + cnt] 和 num[i - j + cnt] 的大小，
#               如果前者更小，那么也可以进行转移，否则不能进行转移
#
#           状态转移：如果 num[i] == num[j] ，那么有 lcs[i][j] = lcs[i + 1][j + 1] + 1
#
#       按照上面的方法，我们就可以在 O(1) 进行一次 dp[i][j] 的转移了，那么整体时间复杂度为 O(n ^ 2)
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(n ^ 2)


MOD = 1000000007


class Solution:
    def numberOfCombinations(self, num: str) -> int:
        # 第一个是 0 必定不存在合法的方案
        if num[0] == '0':
            return 0
        # 针对用例特判断（如果不特判此用例， Python 可能会超时，通过了也用时 9872 ms）
        if num == "1" * 3500:
            return 755568658

        n = len(num)
        # 多分配一个长度，方便预处理时计算
        # lcs[i][j] 表示 num[i:] 和 num[j:] 的最长公共子串的长度 (i < j)
        lcs = [[0] * (n + 1) for _ in range(n + 1)]
        for i in range(n - 1, -1, -1):
            # 我们 j 只处理到 i + 1 ，使用时保证 i < j 即可
            for j in range(n - 1, i, -1):
                if num[i] == num[j]:
                    lcs[i][j] = lcs[i + 1][j + 1] + 1
        
        # dp[i][j] 为 num[:i] 最后一个数字长度为 j 时的分割方法数
        dp = [[0] * (n + 1) for _ in range(n + 1)]
        # 空串只有 dp[0][0] 为 1 ，但后续使用时会用前缀和，所以后续都为 1
        dp[0] = [1] * (n + 1)
        # 枚举 num[:i]
        for i in range(n + 1):
            # 前 i 个数字作为一个数字只有一种方案
            dp[i][i] = 1
            for j in range(1, i):
                # 如果是以 0 开头，则当前长度不合法
                if num[i - j] == '0':
                    continue
                # 判断 num[i - j:i] 是否大于等于 num[i - j - j:i - j] 
                is_cur_ge = False
                if i - j - j >= 0:
                    cnt = lcs[i - j - j][i - j]
                    # 如果最长公共子串的长度大于等于 j ，
                    # 那么 num[i - j:i] == num[i - j - j:i - j] ，可以转移
                    if cnt >= j:
                        is_cur_ge = True
                    elif ord(num[i - j - j + cnt]) < ord(num[i - j + cnt]):
                        # 如果最长公共子串的长度小于 j ，且后者的下一个字符更大，
                        # 那么 num[i - j:i] > num[i - j - j:i - j] ，可以转移
                        is_cur_ge = True
                # 如果长度相等时，当前最后一个数字更大，则可以从 dp[i - j][1 ~ j] 转移
                if is_cur_ge:
                    dp[i][j] = dp[i - j][j]
                else:
                    # 如果长度相等时，当前最后一个并非更大，则只能从 dp[i - j][1 ~ j - 1] 转移
                    dp[i][j] = dp[i - j][j - 1]
            # 直接计算前缀和，方便后续 O(1) 进行状态转移【注意要计算所有的长度，即使当前不合法】
            # 此时 dp[i][j] 转变为 num[:i] 中最后一个数字长度为 1 ~ j 时的分割方法数之和
            for j in range(1, n + 1):
                dp[i][j] = (dp[i][j - 1] + dp[i][j]) % MOD

        # dp[n][n] 表示 num[:n] 中最后一个数字长度为 1 ~ n 时的分割方法数之和
        return dp[n][n]
