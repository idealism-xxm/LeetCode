# 链接：https://leetcode.com/problems/happy-number/
# 题意：给定一个正整数，不断求每一位的平方和，判断最终是否能变为 1 ？


# 数据限制：
#  1 <= n <= 2 ^ 31 - 1


# 输入： 19
# 输出： true
# 解释： 1^2 + 9^2 = 82
#       8^2 + 2^2 = 68
#       6^2 + 8^2 = 100
#       1^2 + 0^2 + 0^2 = 1

# 输入： n = 2
# 输出： false


# 思路： Set
#
#      用 used 维护已出现的数字集合。
#
#      如果 n 不在 used 中，则将 n 放入 used 中，不断重复计算平方和的过程。
#
#      如果在循环中计算出的平方和为 1 ，那么直接返回 true 。
#
#      最后退出循环时，表明遇到了重复的数，则不能变为 1 ，返回 false 。
#
#
#      时间复杂度： O(logn)
#          1. 每次计算下一个数时，需要遍历全部 O(logn) 位。
#              最多会计算 O(810 + 1) 次下一个数： 32 位整型直接按照 10 个 9 计算，
#              则最终能产生的数为 81 * 10 = 810 。
#      空间复杂度： O(1)
#          1. 需要维护 used 中全部 O(810 + 1) 个数： 32 位整型直接按照 10 个 9 计算，
#              则最终能产生的数为 81 * 10 = 810 。


class Solution:
    def isHappy(self, n: int) -> bool:
        # used 维护已出现的数字集合
        used: Set[int] = set()
        # 当 n 未出现过时，继续计算下一个数
        while n not in used:
            # 标记 n 已出现过
            used.add(n)
            # 计算下一个数，即求 n 的每一位的平方和
            nxt: int = 0
            # 当 n 不为 0 时，则还可以计算平方和
            while n > 0:
                # 获取 n 剩余的最后一位
                digit: int = n % 10
                # 计算平方和加入 nxt
                nxt += digit * digit
                n //= 10

            # 如果下一个数字是 1 ，则直接返回 true
            if nxt == 1:
                return True

            # 准备计算下一个数
            n = nxt

        # 此时遇到了重复的数，说明不能变为 1
        return False
