# 链接：https://leetcode.com/problems/finding-pairs-with-a-certain-sum/
# 题意：给定两个整数数组 nums1 和 nums2 ，需要实现一个数据结构满足以下两个操作：
#       1. add(index, val): 给 nums2 下标为 index 的数加上 val ，
#                           即 nums2[index] += val
#       2. count(tot): 计算有多少对 (i, j) 满足 nums1[i] + num2[j] == tot

# 数据限制：
#   1 <= nums1.length <= 1000
#   1 <= nums2.length <= 10 ^ 5
#   1 <= nums1[i] <= 10 ^ 9
#   1 <= nums2[i] <= 10 ^ 5
#   0 <= index < nums2.length
#   1 <= val <= 10 ^ 5
#   1 <= tot <= 10 ^ 9
#   add 和 count 最多各调用 1000 次

# 输入： ["FindSumPairs", "count", "add", "count", "count", "add", "add", "count"]
#       [[[1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]], [7], [3, 2], [8], [4], [0, 1], [1, 1], [7]]
# 输出： [null, 8, null, 2, 1, null, null, 11]
# 解释： 
#   FindSumPairs findSumPairs = new FindSumPairs([1, 1, 2, 2, 2, 3], [1, 4, 5, 2, 5, 4]);
#   findSumPairs.count(7);  # 返回 8
#                           # (2,2), (3,2), (4,2), (2,4), (3,4), (4,4) 形成 2 + 5 = 7
#                           # (5,1), (5,5) 形成 3 + 4 = 7
#   findSumPairs.add(3, 2); # 现在 nums2 = [1,4,5,4,5,4]
#   findSumPairs.count(8);  # 返回 2
#                           # (5,2), (5,4) 形成 3 + 5 = 8
#   findSumPairs.count(4);  # 返回 1
#                           # (5,0) 形成 3 + 1 = 4
#   findSumPairs.add(0, 1); # 现在 nums2 = [2,4,5,4,5,4]
#   findSumPairs.add(1, 1); # 现在 nums2 = [2,5,5,4,5,4]
#   findSumPairs.count(7);  # 返回 11
#                           # (2,1), (2,2), (2,4), (3,1), (3,2), (3,4), (4,1), (4,2), (4,4) 形成 2 + 5  = 7
#                           # (5,3), (5,5) 形成 3 + 4 = 7


# 思路： map
#
#       我们注意到 nums1 长度只有 1000 ，且最多会调用 1000 次 count 操作，
#       那么我们每次 count 操作时枚举 nums1 里的数 num ，
#       然后找到 tot - num 在 nums2 中出现的次数 cnt ，
#       那么最终结果就需要加上 cnt
#
#       为了快速从 nums2 中获取一个数字出现的次数，我们可以维护一个 map ，
#       当使用 add 改变一个值 num 时，我们先从 map 中将 num 出现的次数减 1 ，
#       然后将 num + val 出现的次数加 1
#       
#       时间复杂度： O(len(nums1) * p + len(nums2) * q) ，其中 p 表示 count 调用次数， q 表示 add 调用次数 
#       空间复杂度： O(len(nums2) + q)


class FindSumPairs:

    def __init__(self, nums1: List[int], nums2: List[int]):
        self.nums1 = nums1
        self.nums2 = nums2
        # 统计 nums2 中每个数字出现的次数
        self.cnt = collections.Counter(nums2)

    def add(self, index: int, val: int) -> None:
        # 原数字出现次数减 1
        self.cnt[self.nums2[index]] -= 1
        # 更新当前数字
        self.nums2[index] += val
        # 现数字出现次数加 1
        self.cnt[self.nums2[index]] += 1

    def count(self, tot: int) -> int:
        ans = 0
        # 枚举 nums1 中的每个数字
        for num in self.nums1:
            # 统计 tot - num 在 nums2 中出现的次数
            ans += self.cnt[tot - num]
        return ans
        


# Your FindSumPairs object will be instantiated and called as such:
# obj = FindSumPairs(nums1, nums2)
# obj.add(index,val)
# param_2 = obj.count(tot)
