# 链接：https://leetcode.com/problems/next-greater-numerically-balanced-number/
# 题意：如果一个数 x 的中出现的每个一个数位 d 出现的次数都是 d ，那么这个数 x 就是平衡数。
#       给定一个数 n ，求比 n 大的下一个平衡数？

# 数据限制：
#   0 <= n <= 10 ^ 6

# 输入： n = 1
# 输出： 22
# 解释： 
#   22 是一个平衡数，因为 2 出现的次数就是 2

# 输入： n = 1000
# 输出： 1333
# 解释： 
#   1333 是一个平衡数，因为 1 出现的次数就是 1 ，3 出现的次数就是 3
#   注意： 1022 不是一个平衡数，因为 0 没有出现 0 次

# 输入： n = 3000
# 输出： 3133
# 解释： 
#   3133 是一个平衡数，因为 1 出现的次数就是 1 ，3 出现的次数就是 3


# 思路： 枚举
#
#       最开始想到的就是暴力，不过时间复杂度需要自己再探索估计一下。
#       写完了就找大于 1000000 的下一个平衡数，可以发现这个数是 1224444 ，
#       如果每次都从 1 开始枚举到 12244444 ，那么时间复杂度也是 O(n) ，
#       在可接受的范围内
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def nextBeautifulNumber(self, n: int) -> int:
        # 因为要比 n 大，所以 n 先自加
        n += 1
        # 如果当前 n 不是平衡数，则继续自加
        while not self.is_beautiful(n):
            n += 1
        return n
    
    def is_beautiful(self, n: int) -> bool:
        # 统计每个数位出现的次数
        count = defaultdict(int)
        # 如果还不为 0 ，则还有需要统计的数位
        while n > 0:
            # 统计当前最低位
            count[n % 10] += 1
            # 十进制右移一位
            n //= 10
        # 遍历每个数位及其出现的次数
        for digit, cnt in count.items():
            # 如果数位出现的次数不等于他自己，则 n 不是平衡数
            if digit != cnt:
                return False
        # 如果每个数位出现的次数都是他自己，则 n 是平衡数
        return True
