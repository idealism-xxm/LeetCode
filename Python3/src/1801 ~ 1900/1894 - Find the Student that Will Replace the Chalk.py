# 链接：https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
# 题意：给定一个整数数组 chalk, chalk[i] 表示当前学生回答问题需要的粉笔数，
#       老师会按顺序问每个学生，到达最后一个学生后，会再回到第 0 个学生，
#       现有 l 个粉笔，求第一个遇到粉笔不够的学生的下标？

# 数据限制：
#   chalk.length == n
#   1 <= n <= 10 ^ 5
#   1 <= chalk[i] <= 10 ^ 5
#   1 <= k <= 10 ^ 9

# 输入： chalk = [5,1,5], k = 22
# 输出： 0
# 解释： 学生 0: k = 22 - 5 = 17
#       学生 1: k = 17 - 5 = 16
#       学生 2: k = 16 - 5 = 11
#       学生 0: k = 11 - 5 = 6
#       学生 1: k = 6 - 5 = 5
#       学生 2: k = 5 - 5 = 0
#       学生 0: k = 0 - 5 = -5 < 0


# 输入： chalk = [3,4,1,2], k = 25
# 输出： 1
# 解释： 学生 0: k = 25 - 3 = 22
#       学生 1: k = 22 - 4 = 18
#       学生 2: k = 18 - 1 = 17
#       学生 3: k = 17 - 2 = 15
#       学生 0: k = 15 - 3 = 12
#       学生 1: k = 12 - 4 = 8
#       学生 2: k =  8 - 1 = 7
#       学生 3: k =  7 - 2 = 5
#       学生 0: k =  5 - 3 = 2
#       学生 1: k =  2 - 4 = -2 < 0

# 思路： 枚举
#
#       先求出所有学生回答完一轮所需的粉笔数 total,
#       然后直接跳过所有完整的轮， k %= total,
#       最后再从第一个学生枚举，找到第一个粉笔数不够的学生即可
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)


class Solution:
    def chalkReplacer(self, chalk: List[int], k: int) -> int:
        total = sum(chalk)
        k %= total
        for i in range(len(chalk)):
            k -= chalk[i]
            if k < 0:
                return i
        # 仅用于通过类型校验，永远不会走到这
        return -1
