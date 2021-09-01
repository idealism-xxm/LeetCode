# 链接：https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/
# 题意：给定一个整型数组 tasks ，表示需要执行的任务的耗时，
#       再给定一个整数，表示一次工作会话时长，
#       求按照如下要求执行任务，最少需要多少个会话？
#       要求：
#           1. 一个任务开始后，必须在相同会话中执行完毕
#           2. 一个任务结束后，可以立即开始下一个任务
#           3. 可以按照任意顺序执行任务

# 数据限制：
#   n == tasks.length
#   1 <= n <= 14
#   1 <= tasks[i] <= 10
#   max(tasks[i]) <= sessionTime <= 15

# 输入： tasks = [1,2,3], sessionTime = 3
# 输出： 2
# 解释： 
#       第一个会话：执行前 2 个任务，耗时 3 小时
#       第二个会话：执行第 3 个任务，耗时 3 小时

# 输入： tasks = [3,1,3,1,1], sessionTime = 8
# 输出： 2
# 解释： 
#       第一个会话：执行前 4 个任务，耗时 8 小时
#       第二个会话：执行第 5 个任务，耗时 1 小时

# 输入： tasks = [1,2,3,4,5], sessionTime = 15
# 输出： 1
# 解释： 
#       第一个会话：执行全部 5 个任务，耗时 15 小时

# 思路： 状压 DP
#
#       查看数据范围很小 (n < 20) ，可以想到是为了让 O(n * (2 ^ n)) 的时间复杂度通过，
#       那这个一般就是 状压 DP
#
#       定义状态 dp[i] = [cnt_i, remain_i] ， i 为压缩后的状态，
#           如果 i & (1 << j) != 0 ，则表示第 j 个任务已被执行，
#       cnt_i 表示执行完 i 中的所有任务后，最少使用的会话数，
#       remain_i 表示执行完 i 中的所有任务后，最少使用的会话数为 cnt_i 时，最后一个会话最多剩余的时间。
#
#       初始化： dp[i] = [15, 0], dp[0] = [0, 0] ，表示不执行任务时未使用会话，且无剩余时间
#       状态转移：
#           先枚举 0 ~ (1 << n) 的所有状态 i ，其使用的最小会话数为 cnt ，
#               最后一个会话最多剩余时间为 remain ，
#           再枚举下一个要执行的任务 j ，
#               如果 j 在 i 中已执行过，则直接处理下一个，
#               如果 j 在 i 中未执行过，则可以选择执行 j ，并可以更新状态 dp[i | (1 << j)]
#                   先计算该状态下执行 j 后使用的会话数 nxt_cnt 和最后一个会话剩余的时间 nxt_remain ，
#                   再更新状态 dp[i | (1 << j)] ，
#                       如果 nxt_cnt 更小，则更新 dp[i | (1 << j)] = [nxt_cnt, nxt_remain]
#                       如果 nxt_cnt 相等，但 nxt_remain 更大，则更新 dp[i | (1 << j)][1] = nxt_remain     
#
#       时间复杂度： O(n * (2 ^ n))
#       空间复杂度： O(2 ^ n)

class Solution:
    def __init__(self):
        self.dp = [[15, 0]  for _ in range(1 << 15)]

    def minSessions(self, tasks: List[int], sessionTime: int) -> int:
        dp = self.dp
        # 计算最大的状态，然后初始化
        mx = 1 << len(tasks)
        for i in range(mx):
            dp[i][0] = 15
            dp[i][1] = 0
        # 不执行任务时未使用会话，且无剩余时间
        dp[0][0] = 0
        
        # 遍历已经确定的状态，更新后续的状态
        for i in range(mx):
            # 获取当前状态下的最小会话数和最后一个会话最多剩余时间
            cnt, remain = dp[i]
            # 遍历该状态下要执行的任务
            for j, task in enumerate(tasks):
                # 如果这个任务不在状态 i 中，则可以选择执行当前任务
                if ~(i & (1 << j)):
                    # 如果剩余时间不足，则需要新开一个会话
                    if remain < task:
                        nxt_cnt, nxt_remain = cnt + 1, sessionTime - task
                    else:
                        # 如果剩余时间充足，则可以直接使用当前会话
                        nxt_cnt, nxt_remain = cnt, remain - task
                    
                    # 计算执行完当前任务后的状态
                    nxt = i | (1 << j)
                    # 如果从当前状态执行任务后使用的会话数更少，则全部更新
                    if nxt_cnt < dp[nxt][0]:
                        dp[nxt][0], dp[nxt][1] = nxt_cnt, nxt_remain
                    elif nxt_cnt == dp[nxt][0] and nxt_remain > dp[nxt][1]:
                        # 如果使用的会话数相等，但最后一个会话使用的时间更多，则仅更新剩余时间
                        dp[nxt][1] = nxt_remain
        # 返回执行完所有任务后最少使用的会话数
        return dp[mx - 1][0]
