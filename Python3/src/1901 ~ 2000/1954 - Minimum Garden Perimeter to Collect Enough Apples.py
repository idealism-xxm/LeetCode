# 链接：https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/
# 题意：有一个以 (0, 0) 为中心的正方形果园，在每个整数坐标处都收集苹果，
#       在 (r, c) 处可以收集 |r| + |c| 个苹果，
#       求能至少收集 neededApples 个苹果的果园边长？


# 数据限制：
#   1 <= neededApples <= 10 ^ 15

# 输入： neededApples = 1
# 输出： 8
# 解释： 边长为 2 的 果园可以收集 12 个苹果
#       (1, 0) 可以收集 1 个苹果
#       (1, 1) 可以收集 2 个苹果
#       (0, 1) 可以收集 1 个苹果
#       (-1, 1) 可以收集 2 个苹果
#       (-1, 0) 可以收集 1 个苹果
#       (-1, -1) 可以收集 2 个苹果
#       (0, -1) 可以收集 1 个苹果
#       (1, -1) 可以收集 2 个苹果

# 输入： neededApples = 13
# 输出： 16

# 输入： neededApples = 1000000000
# 输出： 5040


# 思路： 数学 + 二分
#
#       我们可以计算第 n 圈可以收集的苹果数量，先计算其中 1/4 的部分，
#       从 (n, 0) 到 (n, n) 再到 (1, n) 这部分的苹果数量和，
#       苹果数量和 = n + (n + 1) + (n + 2) + ... + (n + n) + (n + n - 1) + ... + (n + 1)
#                = 2 * [(n + 1) + (n + 2) + ... + (n + n)] - n
#                = 2 * [(n + 1 + n + n) * n] / 2 - n
#                = 3 * (n ^ 2) + n - n
#                = 3 * (n ^ 2)
#       所以第 n 圈可以收集的苹果数量总和为 12 * (n ^ 2)
#       那么前 n 圈可以收集的
#       苹果数量总和 = 12 * sum(i ^ 2) 
#                  = 12 * [n * (n + 1) * (2n + 1)] / 6
#                  = 2n * (n + 1) * (2n + 1)
#
#       现在我们就可以使用二分，找到第一个使苹果数量总和 >= neededApples 的圈数 n ，
#       那么边长就是 8n
#
#       时间复杂度： O(log(n ^ (1/3)))
#       空间复杂度： O(1)

class Solution:
    def minimumPerimeter(self, neededApples: int) -> int:
        l, r = 1, 100000
        while l <= r:
            mid = (l + r) >> 1
            total = 2 * mid * (mid + 1) * (2 * mid + 1)
            # 如果苹果树不够 neededApples ，则需要更大的正方形
            if total < neededApples:
                l = mid + 1
            else:
                r = mid - 1
        return l * 8
