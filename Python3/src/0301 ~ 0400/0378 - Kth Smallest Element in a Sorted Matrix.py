# 链接：https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
# 题意：给定一个 n * n 的矩阵，每行从左到右递增，每列从上到下递增，
#      返回第 k 小的数字。
#
#      进阶：使用空间复杂度为 O(1) 的算法。
#
#      超进阶：使用时间复杂度为 O(n) 的算法。
#      （在面试中过于高深，可以查看论文 http://www.cse.yorku.ca/~andy/pubs/X+Y.pdf ）


# 数据限制：
#  0 <= s.length <= 1000
#  t.length == s.length + 1
#  s 和 t 只包含小写字母


# 输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
# 输出：13
# 解释：矩阵中的元素如下： [1,5,9,10,11,12,13,(13),15] ，第 8 小的数字是 13 。

# 输入：matrix = [[-5]], k = 1
# 输出：-5


# 思路：二分
#
#      我们可以使用二分进行处理，二分区间为 [l, r] = [matrix[i][0], matrix[i][n - 1]] ，
#      即矩阵中的最小值和最大值。
#
#      二分最后会找到 [l, r] 中最小的数 x ，使得矩阵中小于等于 x 的数字个数恰好大于等于 k 。
#
#      二分找到的 x 必定在矩阵中，因为如果 x 不在矩阵中，
#      那么矩阵中小于等于 x - 1 的数字个数也大于等于 k ，与我们二分的寻找的结果不符。
#
#
#      每次二分时，我们先计算区间中点 mid ，然后统计矩阵中小于等于 mid 的数字个数。
#
#      由于矩阵中每一行的数字是递增的，所以我们可以遍历每一行，
#      然后使用二分找到每一行中小于等于 mid 的数字个数（即 upper_bound ），
#      这样就能在 O(nlogn) 内求出矩阵中小于等于 mid 的数字个数。
#
#          1. count < k: 说明第 k 小的数字在区间右边，二分区间变为 [mid + 1, r]
#          2. count >= k: 说明第 k 小的数字在区间左边，二分区间变为 [l, mid - 1] 。
#              如果恰好 mid 就是所求之数，那么最后区间长度为 1 时，
#              区间必定为 [mid - 1, mid - 1] ，且必有 l = mid - 1 + 1 = mid ，
#              即 l 仍是二分的结果。
#
#
#      设 C 为矩阵中最大值 减去 最小值的差值。
#
#      时间复杂度：O(n * logn * logC)
#          1. 二分区间为 [min, max] ，时间复杂度为 O(logC)
#          2. 每次二分时都需要求矩阵中小于等于 mid 的数，
#              这时需要遍历全部 O(n) 行，每一行都需要用 O(logn) 的二分。
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n: int = len(matrix)
        # 取出矩阵中最小和最大的数字，分别作为二分区间的左右边界
        l: int = matrix[0][0]
        r: int = matrix[n - 1][n - 1]
        # 运用二分找到最小的数 x ，使得矩阵中小于等于 x 的数字个数恰好大于等于 k
        while l <= r:
            # 找到区间中点，计算矩阵中小于等于 mid 的数字个数
            mid: int = l + (r - l) // 2
            count: int = 0
            # 统计每一行中小于等于 mid 的数字个数
            for i in range(n):
                # 由于每一行的数字是递增的，所以可以使用二分（即 upper_bound ），
                # 这样就能在 O(logn) 内求出该行小于等于 mid 的数字个数
                count += bisect.bisect_right(matrix[i], mid)

            if count < k:
                # 如果矩阵中小于 mid 的数字个数小于 k ，
                # 则说明第 k 小的数字在区间右边，
                # 二分区间变为 [mid + 1, r]
                l = mid + 1
            else:
                # 如果矩阵中小于 mid 的数字个数大于等于 k ，
                # 则说明第 k 小的数字在区间左边，
                # 二分区间变为 [l, mid - 1]
                # 
                # （如果恰好 mid 就是所求之数，那么最后区间长度为 1 时，
                #  区间必为 [mid - 1, mid - 1] ，且必有 l = mid - 1 + 1 = mid ，
                #  即 l 仍是二分的结果）
                r = mid - 1

        return l
