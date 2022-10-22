# 链接：https://leetcode.com/problems/contains-duplicate-ii/
# 题意：给定一个数组 nums 和一个整数 k ，
#      判断是否存在一对重复的数的下标差最多为 k ？


# 数据限制：
#  1 <= nums.length <= 10 ^ 5
#  -(10 ^ 9) <= nums[i] <= 10 ^ 9
#  0 <= k <= 10 ^ 5


# 输入： nums = [1,2,3,1], k = 3
# 输出： true
# 解释： nums[0] = nums[3] ，下标差为 3

# 输入： nums = [1,0,1,1], k = 1
# 输出： true
# 解释： nums[2] = nums[3] ，下标差为 1

# 输入： nums = [1,2,3,1,2,3], k = 2
# 输出： false
# 解释： 所有相同的数的下标差都为 3 ，超过了 2


# 思路： Map
#
#      本题是 LeetCode 217 的加强版，限制了重复的两个数的下标之差不超过 k 。
#
#      我们可以用一个 map 维护每个数最后一次出现的下标。
#
#      在遍历数组 nums 的每个数 nums[i] 时，每次先获取其最后一次出现的下标 j 。
#      若存在，且 i - j <= k ，则直接返回 true 。
#      否则将 nums[i] 最后一次出现的下标设置为 i 。
#
#      最后在循环中没有返回，则所有数都不满足题意，直接返回 false
#
#
#      时间复杂度： O(n)
#          1. 需要遍历 nums 中全部 O(n) 个数
#      空间复杂度： O(n)
#          1. 需要维护全部不同数最后一次出现的下标，最差情况下有 O(n) 个


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        # 维护每个数最后一次出现的下标
        num_to_last_index: Dict[int, int] = {}
        for i, num in enumerate(nums):
            j: Optional[int] = num_to_last_index.get(num)
            # 若 num 存在，且 i - j <= k ，则满足题意
            if j is not None and i - j <= k:
                return True

            # num 最后一次出现的下标设置为 i
            num_to_last_index[num] = i
        
        # 在循环中没有返回，则所有数都不满足题意
        return False
