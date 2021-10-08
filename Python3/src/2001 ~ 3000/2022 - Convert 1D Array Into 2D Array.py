# 链接：https://leetcode.com/problems/convert-1d-array-into-2d-array/
# 题意：给定一个一维数组 original ，将其转换为 m * n 的二维数组，如果无法转换，则返回空数组。

# 数据限制：
#   1 <= original.length <= 5 * 10 ^ 4
#   1 <= original[i] <= 10 ^ 5
#   1 <= m, n <= 4 * 10 ^ 4


# 输入： original = [1,2,3,4], m = 2, n = 2
# 输出： [[1,2],[3,4]]
# 解释：
#   长度为 4 的数组可以转换成 2 * 2 的二维数组
#   前两个数变成第一行，后两个数变成第二行

# 输入： original = [1,2,3], m = 1, n = 3
# 输出： [[1,2,3]]
# 解释：
#   长度为 3 的数组可以转换成 1 * 2 的二维数组
#   三个数变成第一行

# 输入： original = [1,2], m = 1, n = 1
# 输出： []
# 解释：
#   长度为 2 的数组不能转换成 1 * 1 的二维数组

# 输入： original = [3], m = 1, n = 2
# 输出： []
# 解释：
#   长度为 1 的数组不能转换成 1 * 2 的二维数组


# 思路： 模拟
#
#       先判断数组的长度 len(original) 是否等于目标二维数组的大小 m * n ，
#       如果不等，则直接返回空数组，
#       如果相等，则按顺序将数组中的元素放入二维数组中，
#
#       时间复杂度： O(len(original))
#       空间复杂度： O(len(original))


class Solution:
    def construct2DArray(self, original: List[int], m: int, n: int) -> List[List[int]]:
        # 如果长度不等于目标二维数组的大小，则直接返回空数组
        if len(original) != m * n:
            return []

        # 定义二维数组
        ans = [None] * m
        # 初始化每一行的数组
        for i in range(m):
            ans[i] = original[i * n: (i + 1) * n]

        return ans
