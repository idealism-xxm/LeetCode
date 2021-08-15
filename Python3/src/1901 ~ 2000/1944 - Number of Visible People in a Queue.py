# 链接：https://leetcode.com/problems/number-of-visible-people-in-a-queue/
# 题意：给定一个数字各不相同的整数数组 heights 代表不同人的高度，
#       第 i 个人能其右侧的人（第 j 个人，即 i < j）
#       当且仅当 min(heights[i], heights[j]) > max(heights[i + 1], ..., heights[j - 1]) ，
#       求每个人能看到右侧的人的数量？


# 数据限制：
#   n == heights.length
#   1 <= n <= 10 ^ 5
#   1 <= heights[i] <= 10 ^ 5
#   heights 中所有的数各不相同

# 输入： heights = [10,6,8,5,11,9]
# 输出： [3,1,2,1,1,0]
# 解释： 
#   第 0 个人能看见第 1, 2, 4 个人
#   第 1 个人能看见第 2 个人
#   第 2 个人能看见第 3, 4 个人
#   第 3 个人能看见第 4 个人
#   第 4 个人能看见第 5 个人
#   第 5 个人不能看见任何人

# 输入： [5,1,2,3,10]
# 输出： [4,1,1,1,0]


# 思路： 单调栈
#
#       第 i 个人必定能看见第 i + 1 个人，
#       1. 如果 heights[i] < heights[i + 1] ，那么就不能看到后面的任何人了
#       2. 如果 heights[i] > heights[i + 1] ，那么可以再看到右边第一个比 i + 1 高的人，
#           不断重复上面的步骤，直到遇到比 i 高的人 j
#
#           因为 heights[j] > heights[i] ，
#           所以无论如何都满足不了题目中 min() > max() 的这个条件
#
#       可以发现，第 i 个人能看到的人是以第 i + 1 个人为首的单调递增子序列，那么可以使用单调栈来解决
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)

class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        cnt = defaultdict(int)
        for ch in s:
            cnt[ch] += 1
        return len(set(cnt.values())) == 1
