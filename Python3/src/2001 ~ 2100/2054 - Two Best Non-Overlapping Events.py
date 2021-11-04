# 链接：https://leetcode.com/problems/two-best-non-overlapping-events/
# 题意：给定一个二维数组 events ，其中 events[i] = [startTime_i, endTime_i, value_i] ，
#       其中 startTime_i 和 endTime_i 是活动 i 的开始和结束时间，
#       如果参加了活动 i 则可以获得值 value_i ，
#       你可以最多参加两个不重叠的活动（仅起止时间相同也不行）。
#       求最多能获得的最大值？

# 数据限制：
#   2 <= events.length <= 10 ^ 5
#   events[i].length == 3
#   1 <= startTimei <= endTimei <= 10 ^ 9
#   1 <= valuei <= 10 ^ 6

# 输入： events = [[1,3,2],[4,5,2],[2,4,3]]
# 输出： 4
# 解释： 
#   参加活动 0 和活动 1 ，获得的最大值为 2 + 2 = 4

# 输入： events = [[1,3,2],[4,5,2],[1,5,5]]
# 输出： 5
# 解释： 
#   参加活动 2 ，获得的最大值为 5

# 输入： events = [[1,5,3],[1,5,1],[6,6,5]]
# 输出： 8
# 解释： 
#   参加活动 0 和活动 2 ，获得的最大值为 3 + 5 = 8


# 思路1： 二分
#
#       我们对 events 按照开始时间升序排序，
#       然后用 max_values 数组记录后缀最大值，
#       max_values[i] 表示 events[i:] 中的最大值，
#       最后从左到右遍历 events[i] = _, end_time, value ，
#       表示我们一定要参与活动 i ，
#       并找到开始时间大于 end_time 中的值最大的活动 target ，
#           1. 如果 target 不存在，则只参加活动 i ，获得的值为 value
#           2. 如果 target 存在，则参加活动 i 和 target ，
#               获得的值为 value + max_values[target]
#       不断更新所有这些值的最大值即可
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def maxTwoEvents(self, events: List[List[int]]) -> int:
        n = len(events)
        # 按照开始时间升序排序
        events.sort()
        # 初始化后缀最大值
        # max_values[i] 表示 events[i:] 中的最大值
        max_values = [0] * n
        # 最后一个本身就是最大值
        max_values[-1] = events[-1][2]
        # 从倒数第二个倒序更新
        for i in range(n - 2, -1, -1):
            # max_values[i] 的最大值来源：
            #   1. 活动 i 的值
            #   2. events[i + 1:] 中的最大值
            max_values[i] = max(events[i][2], max_values[i + 1])

        # 初始化答案为 0
        ans = 0
        for _, end_time, value in events:
            # 找到第一个开始时间比当前活动结束时间 end_time 大的活动下标
            target = bisect.bisect_left(events, [end_time + 1, 0])
            # 如果找不到，则只能参加当前活动
            if target == n:
                ans = max(ans, value)
            else:
                # 如果能找到，则两个活动都可以参加
                ans = max(ans, value + max_values[target])
        return ans


# 思路2： 堆
#
#       思路1 中我们选择定要参加的活动 i ，
#           寻找它之后可参加的活动中值最大的活动 target 。
#       在思路 2 中我们则选定要参加的活动 i ，
#           寻找它之前可参加的活动中最大值 max_value 。
#       由于这个最大值已经在我们前面遍历过了，
#           且只有结束时间小于当前活动开始时间的活动可以被计入。
#       所以我们可以用结束时间的最小堆维护，
#           每次将最小堆中结束时间小于当前活动刚开始时间的活动弹出，
#           更新 max_value 即可
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def maxTwoEvents(self, events: List[List[int]]) -> int:
        n = len(events)
        # 按照开始时间升序排序
        events.sort()
        # 初始化可参加活动中的最大值
        max_value = 0
        # 初始化结束时间的最小堆，值为 (end_time, value)
        pending_events = []

        # 初始化答案为 0
        ans = 0
        for start_time, end_time, value in events:
            # 如果最小堆为空，且堆顶活动的结束时间小于当前活动开始时间，
            # 则该活动可以参加，弹出它并更新 max_value
            while len(pending_events) > 0 and pending_events[0][0] < start_time:
                # 弹出最小堆中的活动
                _, pending_value = heapq.heappop(pending_events)
                # 更新 max_value
                max_value = max(max_value, pending_value)

            # 更新答案
            ans = max(ans, value + max_value)
            # 将当前活动插入最小堆
            heapq.heappush(pending_events, (end_time, value))
        return ans
