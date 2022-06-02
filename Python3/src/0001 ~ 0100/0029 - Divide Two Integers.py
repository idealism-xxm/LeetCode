# 链接：https://leetcode.com/problems/divide-two-integers/
# 题意：给定除数 dividend 和被除数 divisor ，不使用除法、乘法与模，
#      求整数除法 dividend / divisor 的整数商 quotient 。
#
#      quotient 采用截断方式，即保留整数部分，小数部分舍去。
#      如果 quotient 小于 -(2 ^ 31) ，返回 -(2 ^ 31) 。
#      如果 quotient 大于 2 ^ 31 - 1 ，返回 2 ^ 31 - 1 。


# 数据限制：
#  -(2 ^ 31) <= dividend, divisor <= 2 ^ 31 - 1
#  divisor != 0


# 输入： dividend = 10, divisor = 3
# 输出： 3
# 解释： 10/3 = 3.33333... ，截断为 3

# 输入： dividend = 7, divisor = -3
# 输出： -2
# 解释： 7/-3 = -2.33333... ，截断为 -2


# 思路： 位运算
#
#      不能使用除法、乘法与模，那么就从整数除法的定义上来寻找解法。
#
#      对于整数除法 dividend / divisor = quotient + remainder ，
#      表示 dividend 可以将 divisor 最多减去 quotient 次，
#      如果再多减去 1 次，那么 dividend 的符号就会改变，即不够减法。
#
#      因此我们可以使用减法来模拟整数除法，但商可能非常大，
#      不能单纯地只做减法，否则必定超时。
#
#      我们可以每次减去 divisor 的 cnt 倍，由于不能使用乘法，
#      所以要使用位运算替代乘法，那么 cnt 的二进制位就只能有一个 1 。
#
#      那么枚举为 1 的二进制位 i （从 31 开始到 0 ），
#      如果当前剩余的 dividend 能减去 divisor << i ，则执行减法并统计。
#
#
#      设 C 为整数的位数，本题为 32 。
#
#      时间复杂度：O(C)
#          1. 需要遍历全部 O(C) 位
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量


class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        # 只有这种情况会出现溢出，
        # 因为 32 位有符号整数的范围是 [-(2 ^ 31), 2 ^ 31 - 1] ，
        # 此时结果为 2 ^ 31 ，超过了 32 位有妇好整数的最大值，
        # 需要返回 2 ^ 31 - 1 。
        if divisor == -1 and dividend == -2147483648:
            return 2147483647

        # 当除数与被除数符号相同时，结果为正数
        is_positive: bool = (dividend < 0) == (divisor < 0)
        # 除数与被除数取绝对值，方便后续统一处理
        dividend = abs(dividend)
        divisor = abs(divisor)

        # 除法的结果 dividend 减去了 divisor 的次数，用 ans 维护
        ans: int = 0
        # dividend 从大到小减去 divisor << i
        for i in range(31, -1, -1):
            # 如果 dividend 大于等于 divisor << i ，则执行减法并统计。
            # 注意此处会有三个关键点（仅针对 dividend 是有符号 32 位整型）：
            #    1. 使用 dividend >> i 替代 divisor << i ，防止溢出
            #    2. 使用 (dividend >> i) - divisor >= 0 
            #       替代 (dividend >> i) >= divisor ，
            #       这样是兼容 dividend 是 i32::MIN 的情况。
            #       因为 abs(i32::MIN) 还是 i32::MIN ，
            #       但由于存在减法下溢，所以不影响结果，
            #       即 i32::MIN - num = |i32::MIN| - num 。
            #    3. 先将 dividend 转成无符号 32 位整型，
            #       再进行右移，实现带符号右移
            if (dividend >> i) >= divisor:
                dividend -= divisor << i
                # 当前减去了 1 << i 次 divisor ，计入 ans 中
                ans += 1 << i

        if is_positive:
            # 正数直接返回
            return ans

        # 负数返回相反数
        return -ans
