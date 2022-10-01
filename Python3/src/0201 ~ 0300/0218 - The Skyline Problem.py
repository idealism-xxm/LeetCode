# 链接：https://leetcode.com/problems/the-skyline-problem/
# 题意：给定第一象限的一些矩形（底边在 x 轴上），按左边升序排序，
#      用 [l, r, h] 表示左边 x 值，右边 x 值和 y 值，
#      求这些矩形形成的轮廓线横线上左侧的点列表？


# 数据限制：
#  1 <= buildings.length <= 10 ^ 4
#  0 <= left_i < right_i <= 2 ^ 31 - 1
#  1 <= height_i <= 2 ^ 31 - 1
#  buildings 按照 left_i 升序排序


# 输入： buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
# 输出： [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

# 输入： buildings = [[0,2,3],[2,5,3]]
# 输出： [[0,3],[5,0]]


# 思路2： 排序 + 优先队列/堆
#
#      可以发现转折点只会在以下两种情况下产生：
#          1. 当前 x 处恰好好进入一个更高的矩形的左侧，那么转折点就是该矩形的左上点坐标
#          2. 当前 x 处恰好离开当前最高矩形的右侧，那么转折点的横坐标就是 x ，
#              纵坐标为此后最高矩形的高度
#
#      所以我们可以将一个矩形的坐标信息 (l, r, h) 拆成左侧 (l, -h) 和右侧 (r, h) ，
#      以 (x, height) 存储在 heights 数组中。
#
#      这里左侧的高度是用负数表示，一是要和矩形右侧区分开来，二是方便后续处理 x 相同时的情况。
#
#      然后我们就可以对 heights 按照 x 升序排序，再按 height 升序排序。
#      这样后续在遍历时， x 相同时，高的矩形先进入后离开，就无需按 x 分组遍历。
#
#      然后用一个最大堆 cur_heights 维护当前 x 处所有矩形的高度。
#      并用名为 height_count 的 map 维护当前 x 处所有矩形的不同高度的出现次数，
#      标记哪些高度是无效的，方便后续移除。
#
#      此时遍历 heights 中的每个元素 (x, height) ，根据 height 的正负进行处理：
#          1. height < 0: 矩形左侧，将当前高度放入 cur_heights ，并增加出现次数
#          2. height > 0: 矩形右侧，减小出现次数即可。
#              （优先队列/堆 无法删除指定元素，所以等实际取的时候再从 cur_heights 中删除）
#
#      当前 x 处已经考虑所有的进出情况后，移除无效的最大高度，找到有效的最大高度 max_height 。
#      若 max_height 不等于前一处转折点的高度 pre_height ，则出现了转折点，
#      将 (x, max_height) 放入结果列表中，再更新 pre_height 为 max_height 即可。
#
#
#      时间复杂度： O(nlogn)
#          1. 需要遍历 buildings 全部 O(n) 个元素
#          2. 需要对 heights 全部 O(n) 个元素进行排序，时间复杂度为 O(nlogn)
#          3. 需要将 heights 全部 O(n) 个元素放入堆一次，再从堆中取出一次，
#              每次时间复杂度为 O(logn)
#      空间复杂度： O(n)
#          1. 需要维护 heights, cur_heights, height_count 中全部 O(n) 个元素
#          2. 需要维护结果数组 ans 中全部 O(n) 个元素


class Solution:
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        # heights[i] = (x_i, height_i) 表示 x_i 处的矩形高度的变化信息
        #      1. height_i < 0: 表示此处刚进入一个高度为 -height_i 的矩形左侧
        #      2. height_i > 0: 表示此处刚离开一个高度为  height_i 的矩形右侧
        heights: List[Tuple[int, int]] = []
        # 遍历每个矩形，将高度变化信息放入 height 中
        for l, r, h in buildings:
            heights.append((l, -h))
            heights.append((r, h))
        # 先按 x 升序排序，再按 height 升序排序。
        # 这样后续在遍历时， x 相同时，高的矩形先进入后离开，
        # 就无需按 x 分组遍历
        heights.sort()

        # ans 收集所有转折点的坐标
        ans: List[List[int]] = []
        # cur_heights 维护当前 x 处所有矩形的高度
        cur_heights: List[int] = []
        # height_count 维护当前 x 处所有矩形的不同高度的出现次数
        height_count: Dict[int, int] = defaultdict(int)
        # 初始存在地面高度 0 ，方便后面处理不存在任何矩形的情况
        heapq.heappush(cur_heights, 0)
        height_count[0] = 1
        # pre_height 维护上次转折点的高度，初始化为地面高度 0
        pre_height: int = 0
        # 遍历每一个 x 及对应的高度列表
        for x, height in heights:
            if height < 0:
                # 矩形左侧，将当前高度放入
                height_count[-height] += 1
                # heapq 默认是最小堆，这里放入负数转换成最大堆
                heapq.heappush(cur_heights, height)
            else:
                # 矩形右侧，将当前高度移除
                height_count[height] -= 1
                # 优先队列/堆 无法删除指定元素，所以等实际取的时候再从 cur_heights 中删除

            # 当前 x 处已经考虑所有的进出情况后，移除无效的最大高度，找到有效的最大高度
            while height_count[-cur_heights[0]] == 0:
                heapq.heappop(cur_heights)
            max_height: int = -cur_heights[0]
            # 如果当前最大高度不等于前一个最大高度，则出现了转折点，放入结果列表中
            if max_height != pre_height:
                # 当前转折点放入结果列表
                ans.append([x, max_height])
                # 更新上次转折点高度
                pre_height = max_height

        return ans
