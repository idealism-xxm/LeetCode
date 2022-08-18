# 链接：https://leetcode.com/problems/reduce-array-size-to-the-half/
# 题意：给定一个整数数组 arr ，可以选择一个数字集合，
#      并移除 arr 中所有在这个集合中的数字。
#
#      求移除 arr 中至少一半数字所需集合的最小大小？


# 数据限制：
#  2 <= arr.length <= 10 ^ 5
#  arr.length 是偶数
#  1 <= arr[i] <= 10 ^ 5


# 输入： arr = [3,3,3,3,5,5,5,2,2,7]
# 输出： 2
# 解释： arr 长度为 10 ，所以至少移除 5 个数字即可。
#       数字集合选为 {3, 7} ，那么会移除 3,3,3,3,7 共五个数字。
#       其他长度为 2 的合法的数字集合有： {3,5},{3,2},{5,2}


# 输入： arr = [7,7,7,7,7,7]
# 输出： 1
# 解释： arr 只有一种数字，所以数字集合只能选择 {7} ，会移除全部数字。


# 思路： 贪心 + 排序
#
#      题目的意思其实就是让我们选择尽可能少的不同数字，移除 arr 中尽可能多的数字。
#
#      那么只要我们不断贪心地移除出现次数最多的数字，直至移除的数字个数超过数组的一半即可。
#
#      所以我们可以先统计每个数字的出现次数到 num_to_cnt 中，然后按照出现次数降序排序。
#      再按照前面的贪心逻辑找到数字集合的大小即可。
#
#
#      时间复杂度：O(nlogn)
#          1. 需要遍历 arr 中全部 O(n) 个数字
#          2. 需要对全部不同的数字进行排序，最差情况下有 O(n) 个不同数字，
#              对应的排序时间复杂度为 O(nlogn)
#          3. 需要遍历一半不同的数字，最差情况下有 O(n) 个不同数字
#      空间复杂度：O(n)
#          1. 需要维护 num_to_cnt 中全部不同数字的出现次数，最差情况下有 O(n) 个不同数字
#          2. 需要用一个数组维护全部不同的数字，并用于排序，最差情况下有 O(n) 个不同数字


class Solution:
    def minSetSize(self, arr: List[int]) -> int:
        n: int = len(arr)
        # num_to_cnt[num] 表示 num 的出现次数
        num_to_cnt: Counter = Counter(arr)

        # 转换成一个数组，并进行排序。
        # 先按 cnt 降序排序，再按 num 升序排序
        cnt_with_nums: List[Tuple[int, int]] = sorted((-cnt, num) for num, cnt in num_to_cnt.items())

        # ans 表示集合的大小
        ans: int = 0
        # total 表示已经移除的数字的个数
        total: int = 0
        for cnt, num in cnt_with_nums:
            # 计入当前数字及其出现次数
            ans += 1
            total += -cnt
            # 如果现在已经移除的数字的个数大于等于数组的一半，
            # 则 ans 就是满足题意的结果，直接返回
            if (total << 1) >= n:
                return ans

        # 必定会在循环中返回，不会走到这
