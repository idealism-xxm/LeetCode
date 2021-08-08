# 链接：https://leetcode.com/problems/remove-stones-to-minimize-the-total/
# 题意：给定一个整型数组，每次可以将其中一个数 nums[i] 变为 ceil(nums[i] / 2) ，
#       求 k 次这样的操作后， 所有数的和最小是多少？

# 数据限制：
#   1 <= piles.length <= 10 ^ 5
#   1 <= piles[i] <= 10 ^ 4
#   1 <= k <= 10 ^ 5

# 输入： piles = [5,4,9], k = 2
# 输出： 12
# 解释： [5,4,9] -> [5,4,5] -> [3,4,5]

# 输入： piles = [4,3,6,7], k = 3
# 输出： 12
# 解释： [4,3,6,7] -> [4,3,3,7] -> [4,3,3,4] -> [2,3,3,4]



# 思路： 贪心 + 堆
#
#       为了使最终的数最小，那么每次都尽量将当前最大的数执行这个操作，
#
#       我们维护一个最大堆，每次从堆中取当前最大的数 cur ，
#       然后将 ceil(cur / 2) 放回堆，这样操作 k 次后所有数的和就是结果
#
#       使用 queue.PriorityQueue 很容易就超时了，
#       因为它为了支持多线程安全，有一些额外的操作
#
#       时间复杂度： O(klogn)
#       空间复杂度： O(n)

import heapq

class Solution:
    def minStoneSum(self, piles: List[int], k: int) -> int:
        piles = [-pile for pile in piles]
        heapq.heapify(piles)
        while k > 0:
            k -= 1
            cur = -heapq.heappop(piles)
            heapq.heappush(piles, -((cur + 1) >> 1))
        return -sum(piles)
