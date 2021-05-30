# 链接：https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/
# 题意：给定两个长度为 n 的整型数组 nums1 和 nums2，我们可以重排 nums2 ，
#       求 sum(nums1[i] XOR nums2[i]) 的最小值？

# 输入： nums1 = [1,2], nums2 = [2,3]
# 输出： 2
# 解释： 重排 nums2 为 [3, 2] ，
#       则最小值为： (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2

# 输入： nums1 = [1,0,3], nums2 = [5,3,4]
# 输出： 8
# 解释： 重排 nums2 为 [5, 4, 3] ，
#       则最小值为： (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8

# 思路1： 记忆化搜索
#
#       设 count(i) 表示 i 的二进制中 1 的个数，
#       设 dp[i] 表示前 count(i) 个 nums1 与 nums2 中标记的数字已组合的情况下，
#       剩余数字能组合出来的最小 XOR 的和
#       最终 dp[0] 就是我们需要的答案
#
#       初始化 dp[i] = -1, dp[(1 << len(nums2)) - 1] = 0
#       对于每一层 dfs ，我们遍历需要用到的 nums2[j]
#       （保证 nums2[j] 未使用，即 used & (1 << j) 为 0），
#       然后将其与 nums1[count(used)] 求 XOR ，即： nums1[count(used)] ^ nums[j]
#       再递归计算对应的 dfs(used | (1 << j)) 这个状态下的最下 XOR 的和即可，
#       则当前层 dp[used] = min(nums1[count(used)] ^ nums[j] + dfs(used | (1 << j)))
#
#       时间复杂度： O(n * 2 ^ n)
#       空间复杂度： O(2 ^ n)

class Solution:
    def minimumXORSum(self, nums1: List[int], nums2: List[int]) -> int:
        # 初始化 dp
        dp = [-1] * (1 << len(nums2))
        dp[(1 << len(nums2)) - 1] = 0
        return self.dfs(dp, nums1, nums2, 0)

    def dfs(self, dp: List[int], nums1: List[int], nums2: List[int], used: int) -> int:
        # 如果以及计算过了，则直接返回
        if dp[used] != -1:
            return dp[used]

        # 通过 used ，找到当前层需要用到的 nums1 中的哪个数
        num1 = nums1[self.count(used)]
        # 计算所有情况下的最小值
        dp[used] = 0x3f3f3f3f3f3f3f3f
        for j in range(len(nums2)):
            # 只有 nums2[j] 没有用过，才能继续更新
            if not (used & (1 << j)):
                dp[used] = min(dp[used], (num1 ^ nums2[j]) + self.dfs(dp, nums1, nums2, used | (1 << j)))

        return dp[used]

    def count(self, num: int) -> int:
        cnt = 0
        while num > 0:
            cnt += num & 1
            num >>= 1
        return cnt


# 思路2： 状压 DP
#
#       设 count(i) 表示 i 的二进制中 1 的个数，
#       dp[i] 表示前 count(i) 个 nums1 与 nums2 中标记的数字已组合的情况下，
#       所有 XOR 的和的最小值
#
#       初始化 dp[i] = 0x3f3f3f3f3f3f3f3f, dp[0] = 0
#       则已知状态 i 的情况下，遍历下一个使用的 nums2[j] ，有状态转移方程：
#           dp[i | (1 << j)] = min(dp[i | (1 << j)], dp[i] + (num1 ^ nums2[j]))
#
#       时间复杂度： O(n * 2 ^ n)
#       空间复杂度： O(2 ^ n)

class Solution:
    def minimumXORSum(self, nums1: List[int], nums2: List[int]) -> int:
        # 初始化 dp
        mx = 1 << len(nums2)
        dp = [0x3f3f3f3f3f3f3f3f] * mx
        dp[0] = 0
        # 遍历状态 i ，最后全使用的状态 i 不会再进行更新，不需要遍历
        for i in range(mx - 1):
            # 通过状态 i 确定当前要使用的 num1
            num1 = nums1[self.count(i)]
            # 遍历将要使用的 nums2[j]
            for j in range(len(nums2)):
                # 如果这个数没有使用过，则可以用来更新状态
                if not (i & (1 << j)):
                    dp[i | (1 << j)] = min(dp[i | (1 << j)], dp[i] + (num1 ^ nums2[j]))
        return dp[mx - 1]

    def count(self, num: int) -> int:
        cnt = 0
        while num > 0:
            cnt += num & 1
            num >>= 1
        return cnt
