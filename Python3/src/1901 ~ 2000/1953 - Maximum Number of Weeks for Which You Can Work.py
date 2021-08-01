# 链接：https://leetcode.com/problems/maximum-number-of-weeks-for-which-you-can-work/
# 题意：给定 n 个不同的项目的里程碑数量列表 milestones ，每周可以完成一个里程碑，
#       连续两周不能完成同一个项目的里程碑，求最多能完成多少个里程碑？


# 数据限制：
#   n == milestones.length
#   1 <= n <= 10 ^ 5
#   1 <= milestones[i] <= 10 ^ 9

# 输入： milestones = [1,2,3]
# 输出： 6
# 解释： 
#       第 1 周完成第 0 个项目里程碑，
#       第 2 周完成第 2 个项目里程碑，
#       第 3 周完成第 1 个项目里程碑，
#       第 4 周完成第 2 个项目里程碑，
#       第 5 周完成第 1 个项目里程碑，
#       第 6 周完成第 2 个项目里程碑。

# 输入： milestones = [5,2,1]
# 输出： 7
# 解释：
#       第 1 周完成第 0 个项目里程碑，
#       第 2 周完成第 1 个项目里程碑，
#       第 3 周完成第 0 个项目里程碑，
#       第 4 周完成第 1 个项目里程碑，
#       第 5 周完成第 0 个项目里程碑，
#       第 6 周完成第 2 个项目里程碑，
#       第 7 周完成第 0 个项目里程碑。


# 思路： 贪心
#
#       统计所有里程碑数量 sm 和最大里程碑数量 mx ，
#       计算剩余里程碑数量 remain = sm - mx 。
#
#       我们优先安排最多的里程碑，然后后面加一个剩余的一个里程碑，
#       再安排最多的里程碑，然后后面再加一个剩余的一个里程碑，
#       以此类推……
#       最终有两种情况：
#           1. 最多的里程碑安排完成，那么剩余的里程碑每个都可以安排在任意两个里程碑之间，
#               所以所有里程碑都可以完成
#           2. 最多的里程碑没有安排完成，那么最多的里程碑只能安排 remain + 1 个，剩余的里程碑全部安排完成
#
#       综上： min(sm, 2 * remain + 1)
#
#       时间复杂度： O(n)
#       空间复杂度： O(1)

class Solution:
    def numberOfWeeks(self, milestones: List[int]) -> int:
        # 统计里程碑总数 和 里程碑最多的项目的里程碑数
        sm, mx = 0, 0
        for mileston in milestones:
            sm += mileston
            mx = max(mx, mileston)
        
        # 除去最多的里程碑，剩余的里程碑数量
        remain = sm - mx

        return min(sm, 2 * remain + 1)
