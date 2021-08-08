# 链接：https://leetcode.com/problems/minimum-total-space-wasted-with-k-resizing-operations/
# 题意：给定一个长度为整型数组 nums ，表示第 i 个物体的大小，
#       现在有一个可以装东西的包,初始大小可以随意指定。
#       按顺序开始装这些物体，总共能将包改变至任意大小 k 次，求最后的总体空间浪费和？
#       第 i 个物体浪费的空间 = 此时包的大小 - nums[i]

# 数据限制：
#   1 <= nums.length <= 200
#   1 <= nums[i] <= 10 ^ 6
#   0 <= k <= nums.length - 1

# 输入： nums = [10,20], k = 0
# 输出： 10
# 解释： size = [20,20]
#       初始指定大小为 20 ，然后不变，
#       总体空间浪费为 (20 - 10) + (20 - 20) = 10

# 输入： nums = [10,20,30], k = 1
# 输出： 10
# 解释： size = [20,20,30]
#       初始指定大小为 20 ，然后在最后一个物体时改变为 30 ，
#       总体空间浪费为 (20 - 10) + (20 - 20) + (30 - 30) = 10 + 0 + 0 = 10

# 输入： nums = [10,20,15,30,20], k = 2
# 输出： 15
# 解释： size = [10,20,20,30,30]
#       初始指定大小为 20 ，然后在第 1 个改变为 20 ，在第 3 个改变么为 30 ，
#       总体空间浪费为 (10 - 10) + (20 - 20) + (20 - 15) + (30 - 30) + (30 - 20) = 0 + 0 + 5 + 0 + 10 = 15


# 思路1： DP
#
#       设 dp[l][i][j] 表示成功装下前 l 个物体，当前包大小为 i ，改变 j 次包大小时的最小总体空间浪费，
#       由于包大小可能很大，但最多只有 len(nums) 个，所以可以离散化处理一下
#
#       初始化 dp[l][i][j] = INF ， dp[-1][i][0] = 0
#           表示其他结果还需要计算，但初始包大小可以为任意值
#       状态转移： szes[i] >= nums[l] 时才进行转移，否则当前包大小装不下
#           1. 装第 l 个物体时不改变包大小，则从 dp[l - 1][i][j] 转移而来，
#               浪费空间为 dp[l - 1][i][j] + szes[i] - nums[l]
#           2. 装第 l 个物体时改变包大小，则从 min(dp[l - 1][1 ~ len(szes)][j - 1]) 转移而来，
#               浪费空间为 min(dp[l - 1][1 ~ len(szes)][j - 1]) + szes[i] - nums[l]
#
#       由于现在已经时 O((n ^ 2) * k) 的复杂度了，所以在 2 中不能直接遍历转移，
#       我们可以在用一个 mn[l][j] 表示前 l 个物体，改变 j 次包大小时的最小总体空间浪费，
#       那么 2 中的状态转移变为 mn[l - 1][j - 1] + szes[i] - nums[l]
#       总体时间复杂度均为 O((n ^ 2) * k)
#
#       最终答案为 min(mn)
#
#       最后可以使用滚动数组降低空间复杂度
#
#       时间复杂度： O((n ^ 2) * k)
#       空间复杂度： O(n * k)

INF = 1000000000

class Solution:
    def minSpaceWastedKResizing(self, nums: List[int], k: int) -> int:
        # 离散化包的大小
        szes = sorted(set(nums))
        n = len(szes)
        # 初始化 dp 数组
        dp = [None] * n
        for i in range(n):
            dp[i] = [INF] * (k + 1)
            # 初始大小可以任意指定，所以切换 0 次包大小时，浪费空间均为 0
            dp[i][0] = 0
        # 初始化 mn_dp 数组
        mn = [INF] * (k + 1)
        # 切换 0 次包大小时，浪费空间均为 0
        mn[0] = 0
        
        for num in nums:
            nxt_dp = [[INF] * (k + 1)  for _ in range(n)]
            nxt_mn = [INF] * (k + 1)
            for i, sze in enumerate(szes):
                # 如果当前大小不能装下 num ，则直接处理下一个
                if sze < num:
                    continue
                # 当前浪费空间为 diff
                diff = sze - num
                # 改变 0 次则只能一直不改变大小
                nxt_dp[i][0] = dp[i][0] + diff
                nxt_mn[0] = min(nxt_mn[0], nxt_dp[i][0])
                # 剩下的要么不变，要么从 j - 1 变一次到 sze
                for j in range(1, k + 1): 
                    nxt_dp[i][j] = min(dp[i][j], mn[j - 1]) + diff
                    # 同时更新改变 j 次浪费的最小空间和
                    nxt_mn[j] = min(nxt_mn[j], nxt_dp[i][j])
            # 滚动更新
            dp = nxt_dp
            mn = nxt_mn
        return min(mn)


# 思路2： DP
#
#       赛后看到另一种思路，进行了一定程度的转化融合，
#
#       设 dp[i][j] 表示成功装下前 i （从 1 开始）个物体，改变 j 次时的最小总体空间浪费，
#       我们思考这样一种方式：
#       当前已知 dp[i][j] ，那么我们可以再改变 1 次包大小，然后成功装下 nums[i + 1:nxt] 这些物体，
#       那么改变后的包大小为 sze = max(nums[i + 1:nxt]) ，
#       这一段区间的最小总体空间浪费为 sze * (nxt - i) - sum(nums[i + 1:nxt])
#       那么 dp[nxt][j + 1] = min(dp[nxt][j + 1], dp[i][j] + sze * (nxt - i) - sum(nums[i + 1:nxt]))
#
#       初始化 dp[i][j] = INF ， dp[0][0] = 0
#           表示其他结果还需要计算，但不装任何物体且不改变时，浪费空间为 0
#       状态转移：
#           dp[nxt][j + 1] = min(dp[nxt][j + 1], dp[i][j] + sze * (nxt - i) - sum(nums[i + 1:nxt]))
#           枚举 nxt ，表示要成功装下 nums[i + 1:nxt] 内的物体，
#           此时我们可以动态更新 max(nums[i + 1:nxt]) 和 sum(nums[i + 1:nxt])
#       因此时间复杂度为 O((n ^ 2) * k)
#
#       第一次可以任意指定大小，但不计入改变次数，所以我们可以增加一次改变次数，
#       那么最终结果就是 dp[n][k + 1]
#
#       时间复杂度： O((n ^ 2) * k)
#       空间复杂度： O(n * k)

INF = 1000000000

class Solution:
    def minSpaceWastedKResizing(self, nums: List[int], k: int) -> int:
        n = len(nums)
        # 初始化 dp 数组
        dp = [[INF] * (k + 2) for _ in range(n + 1)]
        # 不装任何物体且不切换包大小时，浪费空间为 0
        dp[0][0] = 0

        for i in range(n):
            for j in range(k + 1):
                # 从 dp[i][j] 转移至 dp[nxt][j + 1]
                # sze 表示 max(nums[i + 1:nxt]) 
                # sm 表示 sum(nums[i + 1:nxt])
                sze, sm = 0, 0
                for nxt in range(i + 1, n + 1):
                    # 当前包的大小为这段区间内的最大物体大小
                    sze = max(sze, nums[nxt - 1])
                    # 这段区间内的物体大小总和
                    sm += nums[nxt - 1]
                    # 那么这段区间的浪费空间为 sze * (nxt - i) - sm
                    dp[nxt][j + 1] = min(dp[nxt][j + 1], dp[i][j] + sze * (nxt - i) - sm)
        return dp[n][k + 1]
