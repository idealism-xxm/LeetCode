# 链接：https://leetcode.com/problems/find-all-good-indices/
# 题意：给定一个长度为 n 的整型数组 nums 和一个正整数 k ，
#      返回所有满足以下条件的下标 i （按升序排序）。
#          1. k <= i < n - k
#          2. 子数组 nums[i-k : i] 是非递增的
#          2. 子数组 nums[i+1 : i+1+k] 是非递减的


# 数据限制：
#  n == nums.length
#  3 <= n <= 10 ^ 5
#  1 <= nums[i] <= 10 ^ 6
#  1 <= k <= n / 2


# 输入： nums = [2,1,1,1,3,4,1], k = 2
# 输出： [2,3]
# 解释： 有两个满足题意的下标 i ：
#       · 2: 子数组 [2,1] 是非递增的，子数组 [1,3] 是非递减的
#       · 3: 子数组 [1,1] 是非递增的，子数组 [3,4] 是非递减的
#       注意下标 4 不满足题意，因为子数组 [4,1] 是递减的

# 输入： nums = [2,1,1,2], k = 2
# 输出： []
# 解释： 没有满足题意的下标


# 思路： 前缀和
#
#      我们维护两个数组 lni (longest non-increasing) 和 lnd (longest non-desceasing) 。
#
#      lni[i] 表示以 i 为结束的子数组中，最长非递增的子数组的长度。
#      lnd[i] 表示以 i 为开始的子数组中，最长非递减的子数组的长度。
#
#      对于 lni ，我们可以初始化 lni[0] = 1 ，然后按照如下方式，顺序递推出所有值：
#          1. nums[i] <= nums[i - 1]: 则当前数加入前面的子数组后，仍旧能保持非递增。
#              令 lni[i] = lni[i] + 1
#          2. nums[i] >  nums[i - 1]: 则当前数只有自己作为子数组时，才是非递增的。
#              令 lni[i] = 1
#
#      对于 lnd ，我们可以初始化 lnd[n - 1] = 1 ，然后按照如下方式，倒序递推出所有值：
#          1. nums[i] <= nums[i + 1]: 则当前数加入后面的子数组后，仍旧能保持非递减。
#              令 lnd[i] = lnd[i] + 1
#          2. nums[i] >  nums[i + 1]: 则当前数只有自己作为子数组时，才是非递减的。
#              令 lnd[i] = 1
#
#      最后，我们可以遍历满足条件 1 的下标 i ，
#      如果其满足条件 2 (lni[i - 1] >= k) 和条件 3 (lnd[i + 1] >= k) ，
#      则满足题意，放入 ans 中。
#
#
#      时间复杂度：O(n)
#          1. 需要递推出 lni 和 lnd 中全部 O(n) 个状态
#          2. 需要遍历 nums 中全部 O(n) 个数
#      空间复杂度：O(n)
#          1. 需要维护 lni 和 lnd 中全部 O(n) 个状态
#          2. 需要维护 ans 中全部 O(n) 个结果


class Solution:
    def goodIndices(self, nums: List[int], k: int) -> List[int]:
        n: int = len(nums)
        # lni[i] 表示以 i 为结束的子数组中，最长非递增的子数组的长度。
        # 都初始化为 1 ，方便后续处理
        lni: List[int] = [1] * n
        # 顺序递推出所有值
        for i in range(1, n):
            # 当前数加入前面的子数组后，仍旧能保持非递增
            if nums[i] <= nums[i - 1]:
                lni[i] += lni[i - 1]

        # lnd[i] 表示以 i 为开始的子数组中，最长非递减的子数组的长度。
        # 都初始化为 1 ，方便后续处理
        lnd: List[int] = [1] * n
        # 倒序递推出所有值
        for i in range(n - 2, -1, -1):
            # 当前数加入后面的子数组后，仍旧能保持非递减
            if nums[i] <= nums[i + 1]:
                lnd[i] += lnd[i + 1]

        ans: List[int] = []
        # 遍历满足条件 1 的下标 i
        for i in range(k, n - k):
            # 如果其满足条件 2 和条件 3 ，则放入 ans 中
            if lni[i - 1] >= k and lnd[i + 1] >= k:
                ans.append(i)

        return ans
