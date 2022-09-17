# 链接：https://leetcode.com/problems/find-original-array-from-doubled-array/
# 题意：有一个非负整型数组 original ，将 original 中的每个元素乘 2 再加入其中，
#      然后打乱所有元素的顺序，最终得到长度为 original 的两倍的整型数组 changed 。
#
#      现给定一个整型数组 changed ，求 original 是多少？
#      如果不存在 original ，则返回空数组。


# 数据限制：
#  1 <= changed.length <= 10 ^ 5
#  0 <= changed[i] <= 10 ^ 5


# 输入： changed = [1,3,4,2,6,8]
# 输出： [1,3,4]
# 解释： 一种可能的 original 为 [1,3,4]:
#       - 1 * 2 = 2
#       - 3 * 2 = 6
#       - 4 * 2 = 8
#       其他可能的 original 为 [3,1,4], [4,1,3] 等。

# 输入： changed = [6,3,0,1]
# 输出： []
# 解释： changed 无法通过任何 original 转化而来。

# 输入： changed = [1]
# 输出： []
# 解释： changed 无法通过任何 original 转化而来。


# 思路： 贪心 + 排序 + 双指针
#
#      由于 original 中的所有数字都是非负数，
#      那么 changed 最小的数字必须要在 original 中。
#
#      所以我们可以按照如下操作处理：
#          1. 找到 chenged 中最小的数字 num
#          2. 如果 num * 2 不在 changed 中，则 changed 无法通过任何 original 转化而来，
#              直接返回空数组即可
#          3. 如果 num * 2 在 changed 中，则将 num 和 num * 2 从 changed 中移除，
#              并将 num 放入 original 中，然后继续重复上述操作
#
#      最后还没有返回的话，那么 original 必定是满足题意的，直接返回即可。
#
#      但这样的时间复杂度为 O(n ^ 2) ，需要进一步优化才能通过。
#
#      可以发现我们每次只关心最小的数字，那么可以先对 changed 按升序排序，
#      时间复杂度为 O(nlogn) 。
#
#      然后使用双指针的方法寻找未使用的最小数字 changed[l] ，
#      以及 changed[l] 的两倍 changed[r] ，按照前面的操作循环处理即可，
#      时间复杂度为 O(n) 。   
#
#
#      时间复杂度：O(nlogn)
#          1. 需要对 changed 中全部 O(n) 个数进行排序，时间复杂度为 O(nlogn)
#          2. 需要遍历排序后的全部 O(n) 个数字两次
#      空间复杂度：O(n)
#          1. 需要维护原数组 original 中全部 O(n) 个数字
#          2. 需要标记 changed 中 O(n) 个数字是否已被使用
#              （本实现使用了 changed 自身进行标记，但要被视为使用了 O(n) 的额外空间）


class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        # 按照题目转化所得的 changed 长度必定是偶数
        if len(changed) & 1:
            return []

        # 按照数字升序排序
        changed.sort()
        # original 维护题目所求的原数组
        original: List[int] = []
        # r 指向下一个待使用的两倍数字
        r: int = 1
        # l 指向下一个需要放入 original 的数字
        for l in range(len(changed)):
            # 如果当前数字，已在前面用于转换，则直接处理下一个数字
            if changed[l] == -1:
                continue

            # l 有可能追上 r ，例如 [0,0,0,0]
            r = max(r, l + 1)
            # 此时当前数字必须放入 original ，需要找到其两倍的 changed[r]
            while r < len(changed) and changed[l] << 1 != changed[r]:
                r += 1

            # 如果找不到，则不满足题意，直接返回空数组
            if r == len(changed):
                return []

            # 将 changed[l] 放入原数组，并标记 changed[r] 已使用
            original.append(changed[l])
            changed[r] = -1

        # 此时 original 就是题目所求的原数组
        return original
