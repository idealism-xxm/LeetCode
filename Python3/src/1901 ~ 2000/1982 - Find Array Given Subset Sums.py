# 链接：https://leetcode.com/problems/find-array-given-subset-sums/
# 题意：有一个长度为 2 ^ n 的数组，表示一个长度为 n 的数组的所有子集的和，
#       求这个数组？

# 数据限制：
#   1 <= n <= 15
#   sums.length == 2 ^ n
#   -10 ^ 4 <= sums[i] <= 10 ^ 4

# 输入： n = 3, sums = [-3,-2,-1,0,0,1,2,3]
# 输出： [1,2,-3]
# 解释： 
#       []: sum is 0
#       [1]: sum is 1
#       [2]: sum is 2
#       [1,2]: sum is 3
#       [-3]: sum is -3
#       [1,-3]: sum is -2
#       [2,-3]: sum is -1
#       [1,2,-3]: sum is 0
#
#       [-1,-2,3] 也是一种合法的结果

# 输入： n = 2, sums = [0,0,0,0]
# 输出： [0,0]

# 输入： n = 4, sums = [0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]
# 输出： [0,-1,4,5]
# 解释： 


# 思路： 贪心 + 双指针
#
#       我们按照升序排序数组 sums ，那么 sums[0] 必定是所有负数的和，
#       sums[1] 必定是 sums[0] 对应的集合中减去一个绝对值最小的数 x 。
#       （如果 sums[0] == 0 ，那么 sums[1] 就是最小的正数，包含在上述讨论中）
#
#       然后我们可以将 sums 拆分成两个长度相同的数组 S 和 T ，
#       且 S 中的每一个元素加上 x 后都在 T 中。
#
#       如果 0 在 S 中，那么这个数就是正数，将 x 放入到结果中，
#           且 S 是除去 x 后剩余数字的所有子集的和，可以递归按照相同方式继续求解下一个数。
#       如果 0 在 T 中，那么这个数就是负数，将 -x 放入到结果中。
#           且 T 是除去 x 后剩余数字的所有子集的和，可以递归按照相同方式继续求解下一个数。
#
#       排序时间复杂度： O(log(2 ^ n) * 2 ^ n) = O(n * 2 ^ n)
#       每一次选择一个数时，都要遍历前面一半的数，那么这部分时间复杂度是 O(2 ^ n)
#       总时间复杂度： O(n * 2 ^ n)
#
#
#       时间复杂度： O(n * 2 ^ n)
#       空间复杂度： O(2 ^ n)

class Solution:
    def recoverArray(self, n: int, sums: List[int]) -> List[int]:
        length = len(sums)
        ans = []
        sums.sort()
        while n > 0:
            n -= 1
            # 获取绝对值最小的数
            x = sums[1] - sums[0]
            # 划分成两个大小相同的集合 s 和 t ，且元素两两相差 x
            s, t = [], []
            # 表明这个数是否已经用过
            used = set()
            # 利用双指针遍历数组， 
            # left 表示下一个可以放到 s 中的数的下标，
            # right 表示下一个可以放到 t 中的数下标
            left, right = 0, 0
            while len(s) < (1 << n):
                # 如果 left 指向的数已被使用，则继续处理下一个
                while left in used:
                    left += 1
                # 选择第一个可用的数放到 s 中
                used.add(left)
                s.append(sums[left])
                # 如果 right 指向的数已被使用，或者与本次选择到 s 中的数相差不为 x
                # 则继续处理下一个
                while right in used or sums[right] - sums[left] != x:
                    right += 1
                # 选择第一个与本次选择到 s 中的数相差为 x 的数放到 t 中
                used.add(right)
                t.append(sums[right])
            
            # 如果 0 在 s 中，那么可以通过 s 获取剩余的数
            if 0 in s:
                ans.append(x)
                sums = s
            else:
                # 0 在 t 中，那么可以通过 t 获取剩余的数
                ans.append(-x)
                sums = t

        return ans
