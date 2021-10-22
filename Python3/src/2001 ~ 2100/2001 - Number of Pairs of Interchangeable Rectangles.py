# 链接：https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/
# 题意：给定一组矩形 rectangles ， rectangles[i] = [width_i, height_i] 表示第 i 个矩形的宽和高，
#       如果两个矩形 i 和 j (i < j) 满足 width_i // height_i == width_j // height_j ，
#       那么它们是可互换的，求这些矩形中有多少对可互换的？

# 数据限制：
#   n == rectangles.length
#   1 <= n <= 10 ^ 5
#   rectangles[i].length == 2
#   1 <= width_i, height_i <= 10 ^ 5

# 输入： rectangles = [[4,8],[3,6],[10,20],[15,30]]
# 输出： 6
# 解释：
#   矩形 0 和 1: 4/8 == 3/6
#   矩形 0 和 2: 4/8 == 10/20
#   矩形 0 和 3: 4/8 == 15/30
#   矩形 1 和 2: 3/6 == 10/20
#   矩形 1 和 3: 3/6 == 15/30
#   矩形 2 和 3: 10/20 == 15/30

# 输入： rectangles = [[4,5],[7,8]]
# 输出： 0


# 思路： 数学
#
#       先求出每个矩形的宽高比的最简分数，分子分母都除以最大公约数即可，
#       然后统计每个最简分数的个数，
#       最后就是组合即可，即从每个最简分数中选择两个
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def interchangeableRectangles(self, rectangles: List[List[int]]) -> int:
        # 统计每个最简分数的个数
        cnt = defaultdict(int)
        for w, h in rectangles:
            # 求最大公约数
            g = gcd(w, h)
            # 统计最简分数
            cnt[(w // g, h // g)] += 1
        
        # 计算所有最简分数的组合
        ans = 0
        for value in cnt.values():
            ans += value * (value - 1) // 2
        return ans
