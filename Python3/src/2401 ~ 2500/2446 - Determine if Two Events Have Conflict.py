# 链接：https://leetcode.com/problems/determine-if-two-events-have-conflict/
# 题意：给定同一天的两个事件的起止时间 event1 和 event2 。
#          event1 = [startTime1, endTime1]
#          event2 = [startTime2, endTime2]
#      事件的时间都是 24 小时制，格式为 HH:MM 。
#      判断两个事件是否存在有交叉的时间点？


# 数据限制：
#  evnet1.length == event2.length == 2
#  event1[i].length == event2[i].length == 5
#  startTime1 <= endTime1
#  startTime2 <= endTime2
#  所有时间的格式都是 HH:MM


# 输入： event1 = ["01:15","02:00"], event2 = ["02:00","03:00"]
# 输出： true
# 解释： 两个事件在 [02:00, 02:00] 有交叉

# 输入： event1 = ["01:00","02:00"], event2 = ["01:20","03:00"]
# 输出： true
# 解释： 两个事件在 [01:20, 02:00] 有交叉

# 输入： event1 = ["10:00","11:00"], event2 = ["14:00","15:00"]
# 输出： false
# 解释： 两个事件没有交叉的时间点


# 思路： 模拟
#
#      设 event1 的起止时间分别为 s1, e1 ， event2 的起止时间分别为 s2, e2 。
#      那么当且仅当 s1 <= e2 && s2 <= e1 时，两个事件存在交叉的时间点。
#
#
#      时间复杂度：O(1)
#          1. 只需要常数次布尔运算即可
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量


class Solution:
    def haveConflict(self, event1: List[str], event2: List[str]) -> bool:
        return event1[0] <= event2[1] and event2[0] <= event1[1]
