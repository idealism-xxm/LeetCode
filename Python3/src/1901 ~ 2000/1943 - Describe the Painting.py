# 链接：https://leetcode.com/problems/describe-the-painting/
# 题意：给定一个线段数组 segments ， segment[i] = [start_i, end_i, color_i]
#       表示将 [start_i, end_i) 的范围内涂上 color_i ，
#       保证所有颜色均不同，现在输出每个有被涂了颜色的线段的起始点和终点，及其被涂的所有颜色，
#       为了方便，用被涂的所有颜色和表示被涂的所有颜色。


# 数据限制：
#   1 <= segments.length <= 2 * 10 ^ 4
#   segments[i].length == 3
#   1 <= start_i < end_i <= 10 ^ 5
#   1 <= color_i <= 10 ^ 9
#   所有颜色均不同


# 输入： segments = [[1,4,5],[4,7,7],[1,7,9]]
# 输出： [[1,4,14],[4,7,16]]
# 解释： 
#       [1,4) 被涂了颜色 {5,9} ，颜色和为 14
#       [4,7) 被涂了颜色 {7,9} ，颜色和为 16

# 输入： segments = [[1,7,9],[6,8,15],[8,10,7]]
# 输出： [[1,6,9],[6,7,24],[7,8,15],[8,10,7]]
# 解释： 
#       [1,6) 被涂了颜色 {9} ，颜色和为 9
#       [6,7) 被涂了颜色 {9,15} ，颜色和为 24
#       [7,8) 被涂了颜色 {15} ，颜色和为 15
#       [8,10) 被涂了颜色 {7} ，颜色和为 7

# 输入： segments = [[1,4,5],[1,4,7],[4,7,1],[4,7,11]]
# 输出： [[1,4,12],[4,7,12]]
# 解释： 
#       [1,4) 被涂了颜色 {5,7} ，颜色和为 12
#       [4,7) 被涂了颜色 {1,11} ，颜色和为 12


# 思路1： 树状数组
#
#       比赛时想到了可以用线段树，但我们的主要操作是区间更新，单点查询，
#       所以可以用树状数组，但没注意到区分边界点，就换了一种思路，结果写挫了。
#
#       区分边界点很容易，我们直接用 set 记录所有的边界点即可，
#       最后直接遍历边界点，因为非边界点的颜色绝对不会改变
#
#       我们用树状数组来存储，同时进行区间更新，
#       最后我们枚举边界点，判断是否和前一个边界点的颜色相同，
#       如果颜色相同 或者 当前点是边界点，
#           1. 如果前一段有颜色，则需要将前一段的区间放入答案数组中
#           2. 无论前一段是否有眼色，都需要更新 前一段区间的开始点 和 颜色
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class BinaryIndexTree:
    def __init__(self, n: int):
        """初始化树状数组"""
        self.n = n
        self.tree = [0] * (n + 1)
    
    def add(self, i: int, val: int) -> None:
        """对 i 加上 val"""
        while i <= self.n:
            self.tree[i] += val
            i += i & -i
    
    def query(self, i: int) -> int:
        """求 [1, i] 的前缀和"""
        s = 0
        while i > 0:
            s += self.tree[i]
            i -= i & -i
        return s

class Solution:
    def splitPainting(self, segments: List[List[int]]) -> List[List[int]]:
        # 找到所有的分界点
        points = set()
        mx = 0
        for start, end, _ in segments:
            points.add(start)
            points.add(end)
            mx = max(mx, end)

        # 树状数组区间更新
        tree = BinaryIndexTree(mx)
        for start, end, color in segments:
            tree.add(start, color)
            tree.add(end, -color)
        
        # 单点查询，放入答案数组中
        ans = []
        start, pre_color = 0, 0
        # 从小到大枚举边界点
        for i in sorted(points):
            # 获取当前点的颜色
            cur_color = tree.query(i)
            # 如果前一个边界点有颜色，则可以放入答案数组中
            if pre_color:
                ans.append([start, i, pre_color])
            # 更新区间开始点和颜色
            start = i
            pre_color = cur_color
            
        return ans


# 思路2： 差分数组
#
#       赛后看到大佬的思路，可以使用差分数组计算，
#       这样就不需要用树状数组进行区间更新和单点查询了，常数减小，
#
#       为什么可以使用差分数组呢？
#       因为是先进行更新，结束后才需要查询，不用实时查询，所以可以使用差分数组。
#
#       差分数组是什么？
#           d[0] = a[0]
#           d[i] = a[i] - a[i - 1]
#       差分数组维护了当前一个数与前一个数的差，
#       我们区间更新时只需要维护这个差分数组即可。
#
#       由于我们随后可以按顺序遍历，所以一起计算出当前点的颜色
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(n)


class Solution:
    def splitPainting(self, segments: List[List[int]]) -> List[List[int]]:
        # 初始化差分数组
        d = defaultdict(int)
        for start, end, color in segments:
            d[start] += color
            d[end] -= color
        
        ans = []
        start, pre_color = 0, 0
        # 从小到大枚举边界点
        for i in sorted(d):
            # 获取当前点的颜色
            cur_color = pre_color + d[i]
            # 如果前一个边界点有颜色，则可以放入答案数组中
            if pre_color:
                ans.append([start, i, pre_color])
            # 更新区间开始点和颜色
            start = i
            pre_color = cur_color
        return ans
