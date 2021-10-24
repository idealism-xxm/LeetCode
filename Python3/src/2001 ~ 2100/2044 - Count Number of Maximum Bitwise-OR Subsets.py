# 链接：https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
# 题意：给定一个整数数组 nums ，设所有非空子集合中按位或的最大值为 maximum ，
#       求所有非空子集合中按位或的值为 maximum 的子集合个数？

# 数据限制：
#   1 <= nums.length <= 16
#   1 <= nums[i] <= 10 ^ 5

# 输入： nums = [3,1]
# 输出： 2
# 解释： 
#   - [3]
#   - [3,1]

# 输入： nums = [2,2,2]
# 输出： 7
# 解释： 总共有 7 个非空子集合，他们的按位或都是 2

# 输入： nums = [3,2,1,5]
# 输出： 6
# 解释： 
#   - [3,5]
#   - [3,1,5]
#   - [3,2,5]
#   - [3,2,1,5]
#   - [2,5]
#   - [2,1,5]


# 思路： dfs
#
#       所有数字的按位或必定是最大可能的值，那么这个 target 就是我们需要求的的值，
#       然后 dfs 枚举每个数取或者不取，并记录选取的所有数字的按位或 total ，
#       每次遍历完所有数字后，如果 total == target ，则 ans += 1 。
#       时间复杂度为 O(n)
#
#       可以使用剪枝优化，即如果发现 total == target ，就不需要继续 dfs 了，
#       因为后续的数字取与不取都不影响 total 了，
#       那么剩余的数字对答案的贡献为 2 ^ (n - i)
#
#       时间复杂度： O(2 ^ n)
#       空间复杂度： O(n)


class Solution:
    def countMaxOrSubsets(self, nums: List[int]) -> int:
        n = len(nums)
        # 所有数的按位或必定是最大可能的值
        target = 0
        for num in nums:
            target |= num

        def dfs(i: int, total: int) -> int:
            # 如果现在 total 就是 target ，
            # 那么剩余的 n - i 个数怎么选都满足题意，
            # 总共有 2 ^ (n - i) 个组合
            if total == target:
                return 1 << (n - i)
            
            # 如果遍历完所有数都不满足题意，则返回 0
            if i == n:
                return 0
            
            # 当前数字有两种选择方式：不选当前数 和 选当前数，继续 dfs
            return dfs(i + 1, total) + dfs(i + 1, total | nums[i])

        return dfs(0, 0)
