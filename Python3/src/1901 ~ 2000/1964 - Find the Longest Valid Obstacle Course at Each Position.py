# 链接：https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
# 题意：给定一个整型数组，求以每个元素为结尾的最长上升子序列的长度？

# 数据限制：
#   n == obstacles.length
#   1 <= n <= 10 ^ 5
#   1 <= obstacles[i] <= 10 ^ 7

# 输入： obstacles = [1,2,3,2]
# 输出： [1,2,3,3]
# 解释： 
#       i = 0: [(1)] ，[1] 的长度为 1
#       i = 1: [(1),(2)] ，[1,2] 的长度为 2
#       i = 2: [(1),(2),(3)] ，[1,2,3] 的长度为 3
#       i = 3: [(1),(2),3,(2)] ，[1,2,2] 的长度为 3

# 输入： obstacles = [2,2,1]
# 输出： [1,2,1]
# 解释： 
#       i = 0: [(2)] ，[2] 的长度为 1
#       i = 1: [(2),(2)] ，[2,2] 的长度为 2
#       i = 2: [2,2,(1)] ，[1] 的长度为 1

# 输入： obstacles = [3,1,5,6,4,2]
# 输出： [1,1,2,3,2,2]
# 解释：
#       i = 0: [(3)] ，[3] 的长度为 1
#       i = 1: [3,(1)] ，[1] 的长度为 2
#       i = 2: [3,(1),(5)] ，[1,5] 的长度为 2
#       i = 3: [3,(1),(5),(6)] ，[1,5,6] 的长度为 3
#       i = 4: [3,(1),5,6,(4)] ，[1,4] 的长度为 2
#       i = 5: [3,(1),5,6,4,(2)] ，[1,2] 的长度为 2


# 思路： DP + 二分
#
#       这题就是最长上升子序列的模版题，
#       题目数据要求我们使用 O(nlogn) 的解法，
#       我们需要维护一个 mn_last 数组，
#       mn_last[j] 表示长度为 j 的最长上升子序列中最后一个数的最小值，
#       那么 mn_last 必定是单调递增的，
#           如果存在 m > l 使得 mn_last[m] < mn_last[l] ，
#           那么以 mn_last[m] 为结尾的最长上升子序列的第 l 个数 <= mn_last[m] ，
#           与我们假设矛盾，所以 mn_last 必定是单调递增的
#       此时我们可以对 mn_last 进行二分，
#       找到第一个大于 obstacles[i] 的值，
#       此时对应的长度 l 就是以 obstacles[i] 为结尾的最长上升子序列的长度
#       同时我们还要更新 mn_last[l] = min(mn_last[l], obstacles[i])
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)

class Solution:
    def longestObstacleCourseAtEachPosition(self, obstacles: List[int]) -> List[int]:
        n = len(obstacles)
        ans = [None] * n
        ans[0] = 1
        mn_last = [100000000] * (n + 1)
        mn_last[0] = 0
        mn_last[1] = obstacles[0]
        mx_len = 1
        for i in range(1, n):
            cur = obstacles[i]
            # 二分找到第一个大于 cur 的长度，
            # 那么这个长度就是以 cur 为结尾的最长上升子序列的长度
            l, r = 0, mx_len
            while l <= r:
                mid = (l + r) >> 1
                if mn_last[mid] <= cur:
                    l = mid + 1
                else:
                    r = mid - 1
            # 以 cur 为结尾的最长上升子序列的长度为 l
            ans[i] = l
            # 更新最长上升子序列的长度
            mx_len = max(mx_len, l)
            # 更新长度为 l 的最长上升子序列最后一个数的最小值
            mn_last[l] = min(mn_last[l], cur)
        return ans

