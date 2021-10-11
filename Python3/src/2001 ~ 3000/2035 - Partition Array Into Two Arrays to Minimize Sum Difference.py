# 链接：https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/
# 题意：给定一个长度为 2 * n 的整数数组，现在需要将其分成两个长度为 n 的数组，
#       求这两个数组的和的差的最小值？

# 数据限制：
#   1 <= n <= 15
#   nums.length == 2 * n
#   -10 ^ 7 <= nums[i] <= 10 ^ 7

# 输入： nums = [3,9,7,3]
# 输出： 2
# 解释：
#   分成的两个数组为 [3,9] 和 [7,3] ，两个数组的和的差为 2

# 输入： nums = [-36,36]
# 输出： 72
# 解释：
#   分成的两个数组为 [-36] 和 [36] ，两个数组的和的差为 72

# 输入： nums = [2,-1,0,4,-2,-9]
# 输出： 0
# 解释：
#   分成的两个数组为 [2,4,-9] 和 [-1,0,-2] ，两个数组的和的差为 0

# 思路： Meet in the middle + DP + 二分
#
#       （ Meet in the middle 一般可以用于 25 <= n <= 40 的情况）
#
#       直接枚举每个数取与不取，时间复杂度是 O(2 ^ (2 * n)) ，会 TLE ，
#       所以需要使用 Meet in the middle 分成两部分求解，
#       即先分别求 nums[:n] 和 nums[n:] 中分别选取 cnt 个数的和集合 left_sums 和 right_sums ，
#           left_sums[cnt] 表示从 nums[:n] 中选取 cnt 个数的和的集合，
#           right_sums[cnt] 表示从 nums[n:] 中选取 cnt 个数的和的集合。
#       使用 01 背包的方法可以在 O(n * 2 ^ n) 内求得所有的集合，
#       而且空间复杂度为 O(2 ^ n) ，
#       且最多的那个集合的大小为 C(n, n // 2) ，最大为 C(15, 7)  ≈ 6435 。
#
#       假设所有数的和 total = sum(nums) ，其一半为 half = total // 2
#
#       然后我们枚举 nums[n:] 选取 cnt 个数，那么其所有可能的和为 right_sums[cnt] ，
#       那么可以确定 nums[:n] 中会选取 n - cnt 个数，其所有可能的和为 left_sums[n - cnt] 。
#
#       此时对 left_sums[n - cnt] 排序后得到 lsums 列表，
#       我们再枚举 right_sums[cnt] 中的数 rsum ，那么要想使得两个数组的和的差最小，
#       那么就是令 |lsums[i] + rsum - half| 最小，
#       假设这个值可以取得 0 ，那么可得当 lsums[i] == half - rsum 时最小，
#       那么只要 lsums[i] 与 target = half - rsum 差距越小，就越能取得 0 ，
#       所以我们二分 lsums 列表，找到第一个大于等于 target 的数的下标 lindex ，
#       而且最后一个小于 target 的数的下标 lindex - 1 也可能更新结果。
#
#       所以 lsums 中 lindex 和 lindex - 1 这两个下标可能会更新结果。
#
#       最开始求 left_sums 和 right_sums 的时间复杂度为 O(n * 2 ^ n)
#       后续使用排序 + 二分的时间复杂度为 O(log(2 ^ n) * 2 ^ n) = O(n * 2 ^ n)
#       综上：时间复杂度为 O(n * 2 ^ n)
#
#       时间复杂度： O(n * 2 ^ n)
#       空间复杂度： O(2 ^ n)


class Solution:
    def minimumDifference(self, nums: List[int]) -> int:
        # n 是 nums 长度的一般
        n = len(nums) >> 1
        # 初始化 left_sums 和 right_sums
        left_sums = [set() for i in range(n + 1)]
        right_sums = [set() for i in range(n + 1)]
        # 不取任何数时， sum 为 0
        left_sums[0].add(0)
        right_sums[0].add(0)
        # 枚举要选取的数
        for i in range(n):
            # 枚举左半边选取的数字个数 cnt
            # 【注意】 01 背包倒着枚举防止重复选取
            for cnt in range(n, 0, -1):
                # 枚举左半边枚举取 cnt - 1 个数字的和
                for sm in left_sums[cnt - 1]:
                    left_sums[cnt].add(sm + nums[i])
            # 枚举右半边选取的数字个数 cnt
            # 【注意】 01 背包倒着枚举防止重复选取
            for cnt in range(n, 0, -1):
                # 枚举右半边枚举取 cnt - 1 个数字的和
                for sm in right_sums[cnt - 1]:
                    right_sums[cnt].add(sm + nums[n + i])

        # 求出总和以及和的一半
        total = sum(nums)
        half = total >> 1
        # 答案初始化为左半边全选，右半边全不选这种情况
        ans = abs(sum(nums[:n]) - sum(nums[n:]))
        # 枚举左半边选取的数字个数 lcnt
        for lcnt in range(1, n + 1):
            # 获取左半边选取 lcnt 个数字时的和列表，并排序
            lsums = sorted(left_sums[lcnt])
            # 枚举右半边选取 n - lcnt 个数字时的和
            for rsum in right_sums[n - lcnt]:
                # 现在 rsum 已确定，需要从 lsum 中找到 lsum ，
                # 使得 |lsum + rsum - half| 最小
                # 我们假设最小值为 0 ，那么目标 lsum 应该是 half - rsum
                target = half - rsum
                # 二分查找第一个大于等于 target 的 lsum 的下标 lindex
                # 那么 lindex - 1 就是最后一个小于 target 的 lsum 的下标
                # 两者对应的和都可能更新 ans
                lindex = bisect.bisect_left(lsums, target)
                for i in (lindex - 1, lindex):
                    if 0 <= i < len(lsums):
                        # 求出其中一个数组的和 sm
                        # 则另一个数组的和为 total - sm * 2
                        sm = lsums[i] + rsum
                        # 更新 ans
                        ans = min(ans, abs(total - (sm << 1)))
        
        return ans
