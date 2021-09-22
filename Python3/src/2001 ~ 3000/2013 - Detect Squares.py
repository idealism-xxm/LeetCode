# 链接：https://leetcode.com/problems/detect-squares/
# 题意：有一个二维平面，设计一个算法满足以下操作：
#       1. add(point): 添加一个点到这个平面中，允许存在重复的点
#       2. count(point): 查询能和 point 组成的边平行于坐标轴的正方形的数量
#           （任意一个点不同，计为不同的一个）

# 数据限制：
#   point.length == 2
#   0 <= x, y <= 1000
#   add 和 count 总共调用不超过 5000 次

# 输入： ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
#       [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
# 输出： [null, null, null, null, 1, 0, null, 2]
# 解释：
#   DetectSquares detectSquares = new DetectSquares();
#   detectSquares.add([3, 10]);
#   detectSquares.add([11, 2]);
#   detectSquares.add([3, 2]);
#   detectSquares.count([11, 10]); // 返回 1 ，可以选择：
#                                  //   第 1, 2, 3 个点
#   detectSquares.count([14, 8]);  // 返回 0 ，目前的点不足以形成满足题意的正方形
#   detectSquares.add([11, 2]);    // 允许添加重复的点
#   detectSquares.count([11, 10]); // 返回 2 ，可以选择：
#                                  //   第 1, 2, 3 个点
#                                  //   第 1, 3, 4 个点


# 思路： Map
#
#       比赛时枚举了对角线上的坐标值，但其实直接枚举出现的点更快一点
#
#       使用 map 统计每个点出现的次数。
#
#       然后在需要 count 时，有 x, y = point ，
#       我们枚举 map 中每个点 cur_x, cur_y = cur_point，让该点和 point 形成正方形的对角，
#       则有： abs(x - cur_x) == abs(y - cur_y) ，即横坐标差和纵坐标差相等。
#       如果此时 (x, cur_y) 和 (cur_x, y) 均在 map 中有，则可以形成正方形，
#       正方形数量为 map[(cur_x, cur_y)] * map[(x, cur_y)] * map[(cur_x, y)]
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class DetectSquares:

    def __init__(self):
        self.cnt = defaultdict(int)

    def add(self, point: List[int]) -> None:
        # 统计 point 的出现次数
        x, y = point[0], point[1]
        self.cnt[(x, y)] += 1

    def count(self, point: List[int]) -> int:
        ans = 0
        x, y = point
        # 枚举对角线的另一个点
        for cur_x, cur_y in self.cnt:
            # 如果要形成正方形，则面积不能为 0 ，即两个点不在一个地方
            # 如果横坐标差和纵坐标差相等，则两个点是正方形对角线上的点
            if x != cur_x and abs(x - cur_x) == abs(y - cur_y):
                # 如果 (x, cur_y) 和 (cur_x, y) 均在 map 中有，则可以形成正方形
                # 此处先判断 (x, cur_y) 和 (cur_x, y) 是否存在，是为了避免未出现的点
                if (x, cur_y) in self.cnt and (cur_x, y) in self.cnt:
                    ans += self.cnt[(cur_x, cur_y)] * self.cnt[(x, cur_y)] * self.cnt[(cur_x, y)]
        return ans
        


# Your DetectSquares object will be instantiated and called as such:
# obj = DetectSquares()
# obj.add(point)
# param_2 = obj.count(point)
