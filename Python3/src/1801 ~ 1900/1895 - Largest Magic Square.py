# 链接：https://leetcode.com/problems/largest-magic-square/
# 题意：一个 k 阶幻方是指每一行的和、每一列的和、主对角线的和、副对角线的和都相等的 k * k 矩阵。
#       现给定一个 m * n 的矩阵，求其中最大的幻方的阶数？

# 数据限制：
#   m == grid.length
#   n == grid[i].length
#   1 <= m, n <= 50
#   1 <= grid[i][j] <= 10 ^ 6

# 输入： grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
# 输出： 3
# 解释： 以 (1, 1) 为左上角的 3 * 3 矩阵是 3 阶幻方，和为 12


# 输入： grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
# 输出： 2
# 解释： 以 (1, 1) 为左上角的 2 * 2 矩阵是 2 阶幻方，和为 6

# 思路： 枚举
#
#       直接枚举矩阵的左上角位置 (i, j)，
#       然后枚举其阶数 k + 1 ，然后判断是否满足 k + 1 阶幻方的条件
#
#       最开始肯定会想到使用前缀和降低每次计算的时间复杂度，
#       将时间复杂度变为 O(m ^ 2 * n ^ 2)
#
#       如果直接求前缀和，那么会占用 O(m * n) 的空间，
#       但由于在判断是否是幻方的时候，还要进行 O(m + n) 的遍历，
#       所以可以不用提前求出前缀和，每次动态计算即可，
#       那么空间复杂度降低为 O(m + n)
#
#       时间复杂度： O(m ^ 2 * n ^ 2)
#       空间复杂度： O(m + n)


class Solution:
    def largestMagicSquare(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        row_sum = [0] * m
        col_sum = [0] * n
        ans = 0
        # 枚举矩阵的左上角 (i, j)
        for i in range(m):
            for j in range(n):
                # 初始化每一行的和、每一列的和、主对角线的和
                row_sum[0] = 0
                col_sum[0] = 0
                main_sum = 0
                # 枚举 k + 1 阶矩阵
                for k in range(m):
                    # 如果不够 k + 1 阶，则直接处理下一个位置
                    if i + k >= m or j + k >= n:
                        break

                    # 主对角线的和需要加上新纳入的右下角的数
                    main_sum += grid[i + k][j + k]
                    # 副对角线的和需要重新算，初始化为右上角的数
                    sub_sum = grid[i][j + k]
                    # 新加入一行和一列，初始化为右下角的数
                    row_sum[k] = grid[i + k][j + k]
                    col_sum[k] = grid[i + k][j + k]
                    for l in range(k):
                        # 副对角线的从左下角开始加
                        sub_sum += grid[i + k - l][j + l]
                        # 已经记录的每一行都加上该行新增的数
                        row_sum[l] += grid[i + l][j + k]
                        # 已经记录的每一列都加上该列新增的数
                        col_sum[l] += grid[i + k][j + l]
                        # 新增的一行加上前面的每一列的数
                        row_sum[k] += grid[i + k][j + l]
                        # 新增的一列加上前面的每一行的数
                        col_sum[k] += grid[i + l][j + k]

                    # 最后把所有数放到一个集合中
                    st = set(row_sum[:k + 1]) | set(col_sum[:k + 1])
                    st.add(main_sum)
                    st.add(sub_sum)
                    # 如果集合只有一个数，则是 k + 1 阶幻方，更新最大阶数
                    if len(st) == 1:
                        ans = max(ans, k + 1)
        return ans
