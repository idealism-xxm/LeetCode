# 链接：https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/
# 题意：给定一个字符串数组 nums 和一个目标字符串 target ，
#       计算有多少对 (i, j) ，使得 nums[i] + nums[j] == target ？

# 数据限制：
#   2 <= nums.length <= 100
#   1 <= nums[i].length <= 100
#   2 <= target.length <= 100
#   nums[i] 和 target 仅由数字组成
#   nums[i] 和 target 没有前导零


# 输入： nums = ["777","7","77","77"], target = "7777"
# 输出： 4
# 解释：
#   - (0, 1): "777" + "7"
#   - (1, 0): "7" + "777"
#   - (2, 3): "77" + "77"
#   - (3, 2): "77" + "77"

# 输入： original = [1,2,3], m = 1, n = 3
# 输出： [[1,2,3]]
# 解释：
#   长度为 3 的数组可以转换成 1 * 2 的二维数组
#   三个数变成第一行

# 输入： nums = ["123","4","12","34"], target = "1234"
# 输出： 2
# 解释：
#   - (0, 1): "123" + "4"
#   - (2, 3): "12" + "34"

# 输入： nums = ["1","1","1"], target = "11"
# 输出： 6
# 解释：
#   - (0, 1): "1" + "1"
#   - (1, 0): "1" + "1"
#   - (0, 2): "1" + "1"
#   - (2, 0): "1" + "1"
#   - (1, 2): "1" + "1"
#   - (2, 1): "1" + "1"


# 思路： 枚举
#
#       我们先维护两个数组 is_prefix 和 is_suffix ，
#           is_prefix[i] 表示 nums[i] 是否是 target 的前缀，
#           is_suffix[i] 表示 nums[i] 是否是 target 的后缀。
#
#       我们可以在 O(n * m) 的时间内计算出 is_prefix 和 is_suffix 的值。
#
#       接下来我们枚举下标 i 和 j ，
#       如果 is_prefix[i] 和 is_suffix[j] 都是 True ，
#           且 len(nums[i]) + len(nums[j]) == len(target) ，那么这一对下标满足题意 ，
#
#       时间复杂度： O(n * m + n ^ 2) ，其中 n 是 nums 的长度， m 是 target 的长度。
#       空间复杂度： O(n)


class Solution:
    def numOfPairs(self, nums: List[str], target: str) -> int:
        n = len(nums)
        m = len(target)

        # 初始化 is_prefix 和 is_suffix
        is_prefix = [False] * n
        is_suffix = [False] * n
        for i in range(n):
            is_prefix[i] = target.startswith(nums[i])
            is_suffix[i] = target.endswith(nums[i])      

        # 统计结果
        ans = 0
        # 枚举 i 和 j
        for i in range(n):
            for j in range(i + 1, n):
                # 如果长度加起来等于 m ，才有可能组成 target
                if len(nums[i]) + len(nums[j]) == m:
                    # 如果 i 是前缀且 j 是后缀，则满足题意
                    if is_prefix[i] and is_suffix[j]:
                        ans += 1
                    # 如果 i 是后缀且 j 是前缀，则满足题意
                    if is_prefix[j] and is_suffix[i]:
                        ans += 1

        return ans
