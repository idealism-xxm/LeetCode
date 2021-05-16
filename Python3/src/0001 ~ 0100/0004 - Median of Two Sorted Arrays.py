# 链接：https://leetcode.com/problems/median-of-two-sorted-arrays/
# 题意：给定两个有序的数组，返回它们的中位数？

# 输入： nums1 = [1,3], nums2 = [2]
# 输出： 2.00000
# 解释： 合并后的数组为 [1,2,3], 中位数是 2

# 输入： nums1 = [1,2], nums2 = [3,4]
# 输出： 2.50000
# 解释： 合并后的数组为 [1,2,3,4], 中位数是 (2 + 3) / 2 = 2.5

# 输入： nums1 = [0,0], nums2 = [0,0]
# 输出： 0.00000
# 解释： 合并后的数组为 [0,0,0,0], 中位数是 (0 + 0) / 2 = 0

# 输入： nums1 = [], nums2 = [1]
# 输出： 1.00000
# 解释： 合并后的数组为 [1], 中位数是 1

# 思路： 二分
#
#       这题堪称劝退题，第一次做根本想不到 O(logn) 如何去解决，虽然再做时已经忘了具体解法，
#       但解决思路的关键点仍旧记得：充分利用已知信息——中位数，来确定数组的分割位置即可。
#
#       两个数组分别分割如下 nums1[:i], nums1[i:] 和 nums2[:j], nums2[j:] ，
#       我们假设中位数从 nums1[i - 1], nums1[i], nums2[j - 1], nums2[j] 中产生，
#       那么必有必有如下大小关系：nums1[i - 1] <= nums2[j] && nums2[j - 1] <= nums1[i]
#       且有如下相等关系：i + j == (m + n) / 2
#
#       1. 如果 m + n 是奇数，则有 i + j = (m - i) + (n - j) - 1 ，
#           即中位数为 min(nums1[i], nums2[j])
#       2. 如果 m + n 是偶数，则有 i + j = (m - i) + (n - j) ，
#           即中位数为 (max(nums1[i - 1], nums2[j - 1]) + min(nums1[i], nums2[j])) / 2
#
#       现在我们再考虑分割点不恰好再中位数时的情况，即前面提到的大小关系和相等关系有一个不满足，
#       那我们直接保证相等关系满足，这样就只有大小关系不满足着一种可能。
#       假设当前 nums1 被分成 nums1[:i] 和 nums1[i:] ，那么为了保证相等关系满足，
#       nums2 被分成 nums2[:j] 和 nums2[j:]，分割点 j = (m + n) / 2 - i 。
#
#       此时我们判断分割点附近四个数的大小，分别处理即可：
#       1. nums1[i - 1] <= nums2[j] && nums2[j - 1] <= nums1[i]:
#           则当前分割点以满足大小关系，根据前面提到的方式计算出中位数即可
#       2. nums1[i - 1] <= nums2[j] && nums2[j - 1] > nums1[i]:
#           则此时有： nums1[i - 1]  <= nums1[i] < nums2[j - 1] <= nums2[j] ，
#           按照正确的分割点，应该要保证 nums2[j - 1] <= nums1[i] ，
#           说明 nums1[i] 偏小了，正确分割点在 (i, m) 内
#       3. nums1[i - 1] > nums2[j] && nums2[j - 1] <= nums1[i]:
#           则此时有： nums2[j - 1] <= nums2[j] < nums1[i - 1] <= nums1[i] ，
#           按照正确的分割点，应该要保证 nums1[i - 1] <= nums2[i] ，
#           说明 nums1[i - 1] 偏大了，正确分割点在 [0, i) 内
#       4. nums1[i - 1] > nums2[j] && nums2[j - 1] > nums1[i]:
#           则此时有： nums1[i - 1] <= nums1[i] < nums2[j - 1] <= nums2[j] < nums1[i - 1]
#           由于存在严格小于，所以这种情况不存在，可以直接排除
#
#       时间复杂度： O(logn)
#       空间复杂度： O(1)

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        m, n = len(nums1), len(nums2)
        # 保证 nums1 长度更短，同时可以更好处理边界情况
        if m > n:
            nums1, nums2 = nums2, nums1
            m, n = n, m

        half = (m + n) // 2
        # 二分 nums1 的区间 [0, m)
        l, r = 0, m
        while l < r:
            # nums1 在 [l, r) 内取中间的位置为分割点
            i = (l + r) // 2
            # 计算出 nums2 的分割点
            j = half - i
            # 判断此时分割点数字的大小关系
            if j > 0 and nums2[j - 1] > nums1[i]:
                # 情况 2 ，正确分割点在 (i, m) 内
                l = i + 1
            else:
                # 情况 3 ，正确分割点在 [0, i) 内
                # 情况 1 ，即使是一个正确的分割点，但是区间大小没有收缩到 1 ，需要继续处理
                #   这样最终 nums1 正确的分割点就是 l
                r = i

            # 情况 4 不存在

        # 目前已找到正确分割点 l
        i, j = l, half - l
        # 找到右边部分两个数的较小值，注意处理好边界情况
        if i == m:
            right_min = nums2[j]
        elif j == n:
            right_min = nums1[i]
        else:
            right_min = min(nums1[i], nums2[j])

        # 如果是奇数，则中位数就是右边部分两个数的较小值
        if (m + n) & 1:
            return right_min

        # 找到左边部分两个数的较大值，注意处理好边界情况
        if i == 0:
            left_max = nums2[j - 1]
        elif j == 0:
            left_max = nums1[i - 1]
        else:
            left_max = max(nums1[i - 1], nums2[j - 1])

        # 偶数，中位数就是 (左边部分两个数的较大值 + 右边部分两个数的较小值) / 2
        return (left_max + right_min) / 2
