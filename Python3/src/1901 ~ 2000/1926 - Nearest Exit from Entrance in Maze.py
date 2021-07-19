# 链接：https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
# 题意：给定一个 m * n 的迷宫，初始在 (row, col) 处，
#       判断是否能走到出口（其他的迷宫边界处）？

# 数据限制：
#   maze.length == m
#   maze[i].length == n
#   1 <= m, n <= 100
#   maze[i][j] 是 '.' 或 '+'
#   entrance.length == 2
#   0 <= entrance_row < m
#   0 <= entrance_col < n
#   entrance 所在位置一定是 '.'

# 输入： maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
# 输出： 1
# 解释： 有三个出口 [1,0], [0,2], [2,3]
#       ++.+
#       ..E+
#       +++.

# 输入： maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
# 输出： 2
# 解释： 有一个出口 [1,2]
#       +++
#       E..
#       +++

# 输入： maze = [[".","+"]], entrance = [0,0]
# 输出： -1
# 解释： 没有出口
#       E+

# 思路： BFS
#
#       用 BFS 即可，找到的第一个出口就是最短路径
#       主要要标记入口不能算作出口
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(m * n)


import queue

dr = [-1, 0, 1, 0]
dc = [0, 1, 0, -1]

class Solution:
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        m, n = len(maze), len(maze[0])
        er, ec = entrance
        q = queue.Queue()
        q.put((er, ec))
        # 标记入口已走过
        maze[er][ec] = 0
        while not q.empty():
            r, c = q.get()
            for ddr, ddc in zip(dr, dc):
                rr, cc = r + ddr, c + ddc
                # 如果下一个位置是空地，则可以走
                if self.is_ok(m, n, rr, cc) and maze[rr][cc] == '.':
                    # 如果是出口，则直接返回
                    if rr in (0, m - 1) or cc in (0, n - 1):
                        return maze[r][c] + 1
                    maze[rr][cc] = maze[r][c] + 1
                    q.put((rr, cc))
        return -1

    def is_ok(self, m, n, r, c):
        return 0 <= r < m and 0 <= c < n
