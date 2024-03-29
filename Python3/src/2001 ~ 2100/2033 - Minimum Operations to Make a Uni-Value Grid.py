# 链接：https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
# 题意：给定一个 m * n 的二维数组和一个整数 x ，有一个操作：可以对任意一个值加上或减去 x ，
#       求至少多少次后，能让所有的数都相等？

# 数据限制：
#   m == grid.length
#   n == grid[i].length
#   1 <= m, n <= 10 ^ 5
#   1 <= m * n <= 10 ^ 5
#   1 <= x, grid[i][j] <= 10 ^ 4

# 输入： grid = [[2,4],[6,8]], x = 2
# 输出： 4
# 解释：
#   - 给 2 加上 x 一次
#   - 从 6 减去 x 一次
#   - 从 8 减去 x 两次
#   总共四次操作

# 输入： grid = [[1,5],[2,3]], x = 1
# 输出： 5
# 解释：
#   - 给 1 加上 x 两次
#   - 从 5 减去 x 两次
#   - 给 2 加上 x 一次
#   总共五次操作

# 输入： grid = [[1,2],[3,4]], x = 2
# 输出： -1

# 思路： 排序
#
#       想将所有数收集到一个数组中，然后排序，
#       那么所有的数都变成中位数 nums[n // 2] 时，需要的操作最少。
#
#       假设全部变成中位数需要的操作数为 total ，
#       1. 如果数组长度 n 是奇数：
#           那么中位数 nums[n // 2] 左右两边各 n // 2 个数字，
#           如果转换成所有数都变成 nums[n // 2] - x 或者 nums[n // 2] + x 时，
#           会减少 n // 2 次操作，但会增加 n // 2 + 1 次操作（多出来的这次是中位数转换的操作）
#
#       2. 如果数组长度 n 时偶数：
#           那么中位数 nums[n // 2] 左边有 n // 2 个数字，右边有 n // 2 - 1 个数字，
#           (1) 如果转换成所有数都变成 nums[n // 2] + x 时，
#               会减少 n // 2 - 1 次操作，但会增加 n // 2 + 1 次操作（多出来的这次是中位数转换的操作）
#           (2) 如果转换成所有数都变成 nums[n // 2] - x 时，
#               会减少 n // 2 次操作，但会增加 (n // 2 - 1) + 1 = n // 2 次操作，
#               这样的操作次数不会发生变化，直至转变为 nums[n // 2 - 1] ，
#               此时再将成所有数都变成 nums[n // 2 - 1] - x 时，
#               会减少 n // 2 - 1 次操作，但会增加 n // 2 + 1 次操作
#
#       综上：当将所有数都变为中位数时，所需操作数最少
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(m * n)


class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        # 将所有数收集成一个列表，并排序
        lst = sorted([num for row in grid for num in row])
        # 找到中位数
        median = lst[len(lst) >> 1]
        # 统计所需要的操作数
        total = 0
        # 枚举每个数
        for num in lst:
            # 计算差值
            diff = abs(num - median)
            # 如果差值不是 x 的整数倍，则不满足题意，直接返回 -1
            if diff % x != 0:
                return -1
            
            # 此时可以变成 median ，加上需要的操作数
            total += diff // x

        return total
