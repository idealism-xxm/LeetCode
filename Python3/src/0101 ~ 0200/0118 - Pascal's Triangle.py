# 链接：https://leetcode.com/problems/pascals-triangle/
# 题意：给定一个整数 numRows ，返回杨辉三角的前 numRows 行。
#
#      在杨辉三角中，每一个数是它左上方和右上方的数之和。


# 数据限制：
#  1 <= numRows <= 30


# 输入： numRows = 5
# 输出： [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

# 输入： numRows = 1
# 输出： [[1]]


# 思路： DP
#
#      按照题意从第一层开始计算即可，
#      对于每一个位置 (i, j) 有 dp[i][j] = dp[i - 1][j - 1] + dp[i][j] 。
#
#      注意边界情况，当 j == 0 || j == i 时， dp[i][j] = 1 。
#
#
#      时间复杂度：O(n ^ 2)
#          1. 需要遍历计算全部 O(n ^ 2) 个位置的数
#      空间复杂度：O(n ^ 2)
#          1. 需要用 dp 维护全部 O(n ^ 2) 个位置的数


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        dp: List[List[int]] = []
        for i in range(numRows):
            # 每一行默认全都是 1 ，直接将边界情况设置好
            dp.append([1] * (i + 1))
            # 仅处理中间非边界情况的数，即它等于左上和右上的数之和
            for j in range(1, i):
                dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]

        return dp
