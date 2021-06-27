# 链接：https://leetcode.com/problems/the-number-of-full-rounds-you-have-played/
# 题意：给定两个形如 HH:MM 的时间，求它们之间有多少个完整的时刻段？
#       完整的时刻段有以下四种： 00 ~ 15, 15 ~ 30, 30 ~ 45, 45 ~ 00

# 数据限制：
#   startTime 和 finishTime 的形式都是 HH:MM
#   00 <= HH <= 23
#   00 <= MM <= 59
#   startTime 和 finishTime 不等

# 输入： startTime = "12:01", finishTime = "12:44"
# 输出： 1
# 解释： 12:15 ~ 12:30

# 输入： startTime = "20:00", finishTime = "06:00"
# 输出： 40
# 解释： 总共度过了 10 个完整的小时， 10 * 4 = 40

# 输入： startTime = "00:00", finishTime = "23:59"
# 输出： 95
# 解释： 总共度过了 23 个完整的小时，最后一个小时有 3 个完整的时刻段，
#       23 * 4 + 3 = 95

# 思路： 数学
#
#       我们先拆分出小时和分钟，
#       startTime -> start_h, start_m
#       finishTime -> end_h, end_m
#
#       当 finishTime 在 startTime 之前时，说明到了第二天，
#       为了方便处理，我们对 startTime 的 end_h 加上 24 进行统一
#
#       ans = 0
#       1. 不考虑分钟，只看完整的小时（存在多算的可能）： ans += (end_h - start_h + 1) * 4
#       2. 扣除 startTime 占据的时刻段： ans -= (start_m + 14) // 15
#       3. 扣除 finishTime 占据的时刻段： ans -= (60 - end_m + 14) // 15
#
#       时间复杂度： O(1)
#       空间复杂度： O(1)


class Solution:
    def numberOfRounds(self, startTime: str, finishTime: str) -> int:
        start_h, start_m = int(startTime[:2]), int(startTime[3:])
        end_h, end_m = int(finishTime[:2]), int(finishTime[3:])
        # 统一结束时间是第二天的情况
        if start_h > end_h or (start_h == end_h and start_m > end_m):
            end_h += 24
        # 1. 不考虑分钟，只看完整的小时（存在多算的可能）
        ans = (end_h - start_h + 1) * 4
        # 2. 扣除 startTime 占据的时刻段
        ans -= (start_m + 14) // 15
        # 3. 扣除 finishTime 占据的时刻段
        ans -= (60 - end_m + 14) // 15
        return ans
