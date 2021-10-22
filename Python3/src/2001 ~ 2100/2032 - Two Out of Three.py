# 链接：https://leetcode.com/problems/two-out-of-three/
# 题意：给定三个整数数组 nums1, nums2 和 nums3 ，求至少在其中两个数组中的数的集合？

# 数据限制：
#   1 <= nums1.length, nums2.length, nums3.length <= 100
#   1 <= nums1[i], nums2[j], nums3[k] <= 100

# 输入： nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
# 输出： [3,2]
# 解释：
#   - 3 在所有数组中都有
#   - 2 在 nums1 和 nums2 中

# 输入： nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
# 输出： [2,3,1]
# 解释：
#   - 2 在 nums2 和 nums3 中
#   - 3 在 nums1 和 nums2 中
#   - 1 在 nums1 和 nums3 中

# 输入： nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
# 输出： []

# 思路： 统计
#
#       我们对每个数组去重后统计每个数字出现的次数，
#       然后将出现次数大于等于 2 的数字收集成列表即可
#
#       时间复杂度： O(len(nums1) + len(nums2) + len(nums3))
#       空间复杂度： O(len(nums1) + len(nums2) + len(nums3))


class Solution:
    def twoOutOfThree(self, nums1: List[int], nums2: List[int], nums3: List[int]) -> List[int]:
        # 统计所有数字出现在不同数组的次数（需要去重）
        num_to_cnt = defaultdict(int)
        for num in set(nums1):
            num_to_cnt[num] += 1
        for num in set(nums2):
            num_to_cnt[num] += 1
        for num in set(nums3):
            num_to_cnt[num] += 1
        
        # 将所有出现次数大于等于 2 的数收集成列表
        ans = [num for num, cnt in num_to_cnt.items() if cnt >= 2]
        return ans
