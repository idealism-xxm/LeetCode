# 链接：https://leetcode.com/problems/eliminate-maximum-number-of-monsters/
# 题意：有 n 个怪物在在进攻城市，每个怪物距离城市为 dist[i] ，速度为 speed[i] ，
#       防御者每秒可以打死一只怪物（包括第 0 分钟），问怪物到达前最多能打死多少只怪物？

# 数据限制：
#   n == dist.length == speed.length
#   1 <= n <= 10 ^ 5
#   1 <= dist[i], speed[i] <= 10 ^ 5

# 输入： dist = [1,3,4], speed = [1,1,1]
# 输出： 3
# 解释：
#      第 0 分钟：怪物距离为 [1,3,4] ，打死第 1 只怪物
#      第 1 分钟：怪物距离为 [X,3,4] ，打死第 2 只怪物
#      第 2 分钟：怪物距离为 [X,X,4] ，打死第 3 只怪物
#      第 3 分钟：所有怪物均被打死

# 输入： dist = [1,1,2,3], speed = [1,1,1,1]
# 输出： 1
# 解释：
#      第 0 分钟：怪物距离为 [1,1,2,3] ，打死第 1 只怪物
#      第 1 分钟：怪物距离为 [X,0,1,2] ，怪物抵达城市

# 输入： dist = [3,2,4], speed = [5,3,2]
# 输出： 1
# 解释：
#      第 0 分钟：怪物距离为 [3,2,4] ，打死第 1 只怪物
#      第 1 分钟：怪物距离为 [X,0,2] ，怪物抵达城市

# 思路： 贪心
#
#       按照到达城市的分钟数升序排序，每一分钟优先打死最先会到达城市的怪物，
#       如果当前分钟大于等于怪物达到的分钟，则统计结束
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def eliminateMaximum(self, dist: List[int], speed: List[int]) -> int:
        # 将怪物数据组装在一起
        monsters = [
            (d, s)
            for d, s in zip(dist, speed)
        ]

        # 比较函数，使用乘法代替除法
        def cmp(a, b):
            return a[0] * b[1] - a[1] * b[0]

        # 按到达时间从小到大的顺序排序
        monsters = sorted(monsters, key=cmp_to_key(cmp))
        # 每一分钟都打死一直怪物，所以打死的怪物数 = 经过分钟数
        minute = 0
        for d, s in monsters:
            # 如果怪物已经到了，则跳出循环
            if minute * s >= d:
                break

            minute += 1
        return minute
