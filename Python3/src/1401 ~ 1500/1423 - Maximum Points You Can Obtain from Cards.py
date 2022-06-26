# 链接：https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
# 题意：给定一个整数数组 cardPoints ，每次可以从左或者从右拿走 1 个数，
#       直至拿走 k 个数，求拿走数之和的最大值？


# 数据限制：
#   1 <= cardPoints.length <= 10 ^ 5
#   1 <= cardPoints[i] <= 10 ^ 4
#   1 <= k <= cardPoints.length


# 输入： cardPoints = [1,2,3,4,5,6,1], k = 3
# 输出： 12
# 解释： 拿走最右边的 3 个数，和为 1 + 6 + 5 = 12

# 输入： cardPoints = [2,2,2], k = 2
# 输出： 4
# 解释： 无论怎么拿，最终拿走的两个数的和都是 4

# 输入： cardPoints = [9,7,7,9,7,7,9], k = 7
# 输出： 55
# 解释： 拿走全部的 7 个数，和为 55


# 思路： 滑动窗口
#
#      如果一道题目需要在所有满足某种状态的连续子串/连续子数组中，
#      找到满足题意的一个，那么可以考虑滑动窗口。
#
#      本题类似 LeetCode 1658 ，可以采用类似的方法处理，
#      基本修改一下思路中的文字和代码就可以完美适配。
#
#      如果我们将 cardPoints 拼接一次，得到一个长度为 2 * len(cardPoints) 的新数组，
#      那么本题就转化为：在新数组中找到一个长度为 k 连续子数组的最大和。
#
#      可以使用滑动窗口的方法解决本题，这样的空间复杂度为 O(n) 。
#      不过可以特殊处理右边界 r 的情况，能让空间复杂度优化为 O(1) ，
#      但不便于理解且容易出错。
#
#      此时可以考虑原问题的镜像问题：求长度为 len(cardPoints) - k 的子数组的最小和 ans ，
#      那么 sum(cardPoints) - ans 就是原问题的答案。
#
#      设 target = sum(cardPoints) - x ，
#      我们使用滑动窗口 [l, r] 表示一个数字和小于等于 target 的连续子数组，
#      初始化为左边界 l = 0 ，右边界 r = target - 1 ，滑动窗口数字和 total = sum(cardPoints[:target]) 。
#
#      我们不断右移右边界 r ，将其纳入到滑动窗口中考虑， total += cardPoints[r] ，
#      同时右移左边界 l 一次，保持滑动窗口长度为 target ，total -= cardPoints[l] 。
#
#      获取所有滑动窗口数字和的最小值 ans ，最后返回 sum(cardPoints) - ans 即可。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历 cardPoints 中全部 O(n) 个数字
#      空间复杂度：O(1)
#          1. 只需用维护常数个额外变量


class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        total: int = sum(cardPoints)
        n: int = len(cardPoints)
        # 原问题可以转化为求长度为 target 的子数组的最小和
        target: int = n - k
        # 如果 target 等于 0 ，则原问题的解为 total - 0 = total
        if target == 0:
            return total

        # 初始化滑动窗口为 [0, target - 1] ，那么滑动窗口内数字和为 sum(cardPoints[:target]) 。
        # 这里使用 for 循环处理，是为了避免为切片分配空间
        amount: int = sum(cardPoints[i] for i in range(target))
        # 初始连续子数组的最小和为 amount
        ans = amount
        # 不断右移右边界 r
        for r in range(target, n):
            # 将 cardPoints[r] 纳入到滑动窗口中考虑
            amount += cardPoints[r]
            # 同时 cardPoints[r - target] 要移出滑动窗口，使滑动窗口长度保持为 target
            amount -= cardPoints[r - target]

            # 更新长度为 k 的连续子数组的最小和
            ans = min(ans, amount)

        return total - ans
