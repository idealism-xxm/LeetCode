# 链接：https://leetcode.com/problems/last-day-where-you-can-still-cross/
# 题意：有一个 row * col 的二维 01 数组（下标从 1 开始），
#       0 代表陆地， 1 代表水，第 0 天都是 0 ，
#       现在有一个长度为 row * col 的二维数组 cells ，
#       cells[i] = [r_i, c_i] （ r_i 和 c_i 都是从 1 开始），
#       表示第 i + 1 天会将 (r_i, c_i) 从陆地变成水，
#       求能从第 1 行四方向走陆地走到第 row 行的最后一天？

# 数据限制：
#   2 <= row, col <= 2 * 10 ^ 4
#   4 <= row * col <= 2 * 10 ^ 4
#   cells.length == row * col
#   1 <= r_i <= row
#   1 <= c_i <= col
#   cells[i] 各不相同

# 输入： row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
# 输出： 2
# 解释： 
#       天：   0     1     2     3
#       陆地：
#             00    10    10    11
#             00    00    10    10

# 输入： row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]
# 输出： 1
# 解释： 
#       天：   0     1     2
#       陆地：
#             00    10    11
#             00    00    00

# 输入： row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]
# 输出： 3
# 解释： 
#       天：   0     1     2     2     2
#       陆地：
#             000   010   010   010   010
#             000   000   100   100   110
#             000   000   000   001   001


# 思路1： 二分 + BFS
#
#       我们可以发现结果具有单调性，
#       即第 i 天不能成功，后续都不能成功，
#       第 i 天能成功，之前都能成功。
#
#       所以我们可以二分天数 l = 1, r = row * col ，
#       二分时间复杂度为 O(log(row * col)) ，
#
#       每一次二分都需要先对图进行标记，然后以第一行的陆地为源点进行多源 BFS 搜索，
#       时间复杂度为 O(row * col)
#       
#       时间复杂度： O(row * col * log(row * col))
#       空间复杂度： O(row * col)

import queue


d = [(-1, 0), (0, 1), (1, 0), (0, -1)]


class Solution:
    def latestDayToCross(self, row: int, col: int, cells: List[List[int]]) -> int:
        l, r = 1, len(cells)
        # 已访问矩阵（同时会将所有的水标记为 True）
        self.visited = [[False] * (col + 1) for _ in range(row + 1)]
        while l <= r:
            mid = (l + r) >> 1
            # 如果中点能成功，则下次二分 [mid + 1, r] 区间
            if self.is_ok(row, col, cells, mid):
                l = mid + 1
            else:
                # 如果中点不能成功，则下次二分 [l, mid - 1] 区间
                r = mid - 1
        # 考虑最终 l == r 时的情况
        # 1. 如果此时能成功，则 l 会变成 l + 1 ，但在前面的二分中已经证明 l + 1 时不可能成功的，
        #    所以此时 r 是最后一天能成功的
        # 2. 如果此时不能成功，则 r 会变成 r - 1 ，在前面有 l - 1 及其左边的区间都能成功，
        #    所以此时 r 还是最后一天能成功的
        return r
    
    def is_ok(self, row: int, col: int, cells: List[List[int]], num: int) -> bool:
        visited = self.visited
        # 先重置 visited 数组
        for i in range(1, row + 1):
            for j in range(1, col + 1):
                visited[i][j] = False
        # 将前 num 位置都标记为不可访问
        for i in range(num):
            r, c = cells[i]
            visited[r][c] = True
        
        # 将第一行的陆地放入队列中，进行多源 BFS
        q = queue.Queue()
        for j in range(1, col + 1):
            if not visited[1][j]:
                q.put((1, j))
        
        # 队列不为空，则可以继续走下一步
        while not q.empty():
            r, c = q.get()
            # 遍历四个方向的下标差
            for dr, dc in d:
                # 计算下一步的下标
                rr, cc = r + dr, c + dc
                # 如果该下标合法，且没有访问过
                if 1 <= rr <= row and 1 <= cc <= col and not visited[rr][cc]:
                    # 如果该位置已到最后一行，则能成功，直接返回 True
                    if rr == row:
                        return True
                    # 标记当前位置已访问，并放入队列中
                    visited[rr][cc] = True
                    q.put((rr, cc))
        # 循环中没有到达最后一行，则不能成功
        return False
