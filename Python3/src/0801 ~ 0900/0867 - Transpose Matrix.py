# 链接：https://leetcode.com/problems/transpose-matrix/
# 题意：给定一个二维数组 matrix ，返回它的转置矩阵。


# 数据限制：
#  m == matrix.length
#  n == matrix[i].length
#  1 <= m, n <= 1000
#  1 <= m * n <= 10 ^ 5
#  -(10 ^ 9) <= matrix[i][j] <= 10 ^ 9


# 输入： matrix = [[1,2,3],[4,5,6],[7,8,9]]
# 输出： [[1,4,7],[2,5,8],[3,6,9]]

# 输入： matrix = [[1,2,3],[4,5,6]]
# 输出： [[1,4],[2,5],[3,6]]


# 思路： 模拟
#
#      直接按照题意模拟即可。
#
#      矩阵 matrix 的大小为 m * n ，那么其转置矩阵 ans 的大小为 n * m ，
#      且对于每一个合法的 i, j ，都有 ans[j][i] = matrix[i][j] 。
#
#
#      时间复杂度：O(m * n)
#          1. 需要遍历 matrix 中全部 O(m * n) 个数字
#      空间复杂度：O(m * n)
#          1. 需要用维护转置矩阵全部 O(m * n) 个数字


class Solution:
    def transpose(self, matrix: List[List[int]]) -> List[List[int]]:
        m, n = len(matrix), len(matrix[0])
        # 定义 n * m 的转置矩阵 ans
        ans = [[0] * m for _ in range(n)]
        for i in range(m):
            for j in range(n):
                # 交换行列下标给转置矩阵 ans 赋值即可
                ans[j][i] = matrix[i][j]

        return ans
