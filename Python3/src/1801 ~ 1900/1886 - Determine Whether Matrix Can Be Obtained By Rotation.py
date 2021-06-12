# 链接：https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
# 题意：给定两个 n * n 的矩阵 mat 和 target ，判断 mat 是否可以通过旋转得到 target ？

# 输入： mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
# 输出： true
# 解释： 顺时针旋转 mat 90° 可以得到 target

# 输入： mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
# 输出： false

# 输入： mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
# 输出： -1
# 解释： 顺时针旋转 mat 180° 可以得到 target

# 思路： 枚举
#
#       枚举旋转的角度为 0°, 90°, 180°, 270° ，
#       如果又一次旋转后两个矩阵相等，则返回 true ，
#       否则，返回 false
#
#       【按层旋转】
#       最先想到的就是最朴素的一层一层旋转，
#       但是这样不太清楚位置关系的变换，
#       比较容易写错
#
#       【两次翻转】
#       推荐先按对角线翻转，再按中轴线翻转，这样也能达到顺时针旋转 90° 的目的
#       mat[i][j] ==主对角线翻转==> mat[j][i] ==中轴线翻转==> mat[j][n - i - 1]
#
#       以上两种方法的本质都是满足一下这个等式：
#           target[j][n - i - 1] = mat[i][j]
#
#       时间复杂度： O(n ^ 2)
#       空间复杂度： O(1)


class Solution:
    def findRotation(self, mat: List[List[int]], target: List[List[int]]) -> bool:
        for _ in range(4):
            if self.is_equal(mat, target):
                return True
            self.rotate(mat)
        return False

    def rotate(self, mat: List[List[int]]):
        """顺时针旋转 mat 90° """
        n = len(mat)
        # 1. 朴素的一层一层旋转
        # for i in range(n // 2):
        #     for j in range(i, n - i - 1):
        #         tmp = mat[i][j]
        #         mat[i][j] = mat[n - j - 1][i]
        #         mat[n - j - 1][i] = mat[n - i - 1][n - j - 1]
        #         mat[n - i - 1][n - j - 1] = mat[j][n - i - 1]
        #         mat[j][n - i - 1] = tmp

        # 2. 先按对角线翻转，再按中轴线翻转
        for i in range(n):
            for j in range(i + 1, n):
                mat[i][j], mat[j][i] = mat[j][i], mat[i][j]

        for i in range(n):
            for j in range(n // 2):
                mat[i][j], mat[i][n - j - 1] = mat[i][n - j - 1], mat[i][j]

    def is_equal(self, mat: List[List[int]], target: List[List[int]]) -> bool:
        """判断两个矩阵是否相同"""
        for i in range(len(mat)):
            for j in range(len(mat[0])):
                if mat[i][j] != target[i][j]:
                    return False
        return True
