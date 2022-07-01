# 链接：https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
# 题意：给定一个整数数组 nums ，求最少多少次操作可以将所有数字相等？
#      每次操作可以将一个数增加或者减小 1 。


# 数据限制：
#  n == nums.length
#  1 <= nums.length <= 10 ^ 5
#  -(10 ^ 9) <= nums[i] <= 10 ^ 9


# 输入： nums = [1,2,3]
# 输出： 2
# 解释： [1,2,3]  =>  [2,2,3]  =>  [2,2,2]

# 输入： nums = [1,10,2,9]
# 输出： 16
# 解释： 将所有数字都变成 2 ，
#       总共需要 (2 - 1) + (10 - 2) + (2 - 2) + (9 - 2)
#             = 1 + 8 + 0 + 7
#             = 16


# 思路： 排序 + DP
#
#      第一反应就是排序 + DP ，先按照大小升序排序，
#      然后递推每一个数作为最终的数计算所需操作数，维护所有所需操作数的最小值即可。
#
#      递推时利用相邻状态之间的关系，就能在 O(n) 内求出所有情况所需的操作数。
#      不过直接按照这种方式计算会有一点问题，某些所需操作数可能会超出 32 位整型的范围，
#      需要用到 64 位整型位数这些值，但后续可以继续优化。
#
#      设 cnt[i] 表示以 nums[i] 为最终数所需的操作数。
#      初始化： cnt[0] = sum(|num[0 ~ n] - nums[0]|)
#      状态转移： cnt[i] = cnt[i - 1] + i * (nums[i] - nums[i - 1]) - (n - i) * (nums[i] - nums[i - 1])
#                      = cnt[i - 1] - (n - 2 * i) * (nums[i] - nums[i - 1])
#
#      观察状态转移方程可以发现， cnt 的变化是可以确定的：
#          1. n - 2 * i >= 0 时，即 i <= n / 2 ， cnt 是递减的
#          2. n - 2 * i < 0 时，即 i > n / 2 ， cnt 是递增的
#
#      那么我们就无需 DP 了， cnt 必定在 i = n / 2 时取得最小值，
#      最终数必定是 nums[n / 2] ，直接计算对应的所需操作数即可。         
#
#
#      时间复杂度：O(nlogn)
#          1. 需要对 nums 进行排序，时间复杂度为 O(nlogn)
#          2. 需要遍历全部 O(n) 个数计算所需操作数之和
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量即可


class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        # 按照大小升序排序
        nums.sort()
        # nums 中点处的数，就是所有数最终要变成的数
        target: int = nums[len(nums) >> 1]
        # 计算把所有数变成 target 所需的操作数之和，
        # 把 num 变为 target 需要操作 abs(num - target) 次
        return sum(abs(num - target) for num in nums)
