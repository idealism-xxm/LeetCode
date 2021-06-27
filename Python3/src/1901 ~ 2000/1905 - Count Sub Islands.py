# 链接：https://leetcode.com/problems/count-sub-islands/
# 题意：给定两个 m * n 的二维整型数组 grid1, grid2 ，其中每个数是 0 或 1 ，
#       0 表示水， 1 表示陆地，一组被水包围的陆地算作一个岛，数组之外的地方都是水，
#       现在求 grid2 中子岛数量？
#       子岛定义： grid2 的岛在 grid1 中对应位置都是陆地。

# 数据限制：
#   m == grid1.length == grid2.length
#   n == grid1[i].length == grid2[i].length
#   1 <= m, n <= 500
#   grid1[i][j] 和 grid2[i][j] 是 0 或 1

# 输入： grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
# 输出： 3

# 输入： grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
# 输出： 2

# 思路： DFS
#
#       我们照常使用 DFS 遍历 grid2 中的岛即可，
#       dfs(grid1, grid2, r, c) 返回 grid2[r][c] 是不是子岛
#
#       在 dfs 时判断： grid1 中对应位置是否为陆地，
#       只有当一个岛中对应在 grid1 中的位置都是陆地是才返回 True,
#       否则返回 False
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(m * n)


dr = [-1, 0, 1, 0]
dc = [0, 1, 0, -1]


class Solution:
    def countSubIslands(self, grid1: List[List[int]], grid2: List[List[int]]) -> int:
        ans = 0
        for i in range(len(grid2)):
            for j in range(len(grid2[0])):
                # 遍历 grid2 的每块陆地，如果是子岛，则 ans += 1
                if grid2[i][j] == 1 and self.dfs(grid1, grid2, i, j):
                    ans += 1
        return ans

    def dfs(self, grid1: List[List[int]], grid2: List[List[int]], r: int, c: int) -> bool:
        """判断 grid2[r][c] 开始的岛是不是子岛"""
        # 标记当前陆地已访问过
        grid2[r][c] = 0
        m, n = len(grid2), len(grid2[0])

        # 当前位置对应的 grid1 中如果不是陆地，则不是子岛，但还需要完全遍历
        res = grid1[r][c] == 1
        for ddr, ddc in zip(dr, dc):
            rr, cc = r + ddr, c + ddc
            if self.is_ok(m, n, rr, cc) and grid2[rr][cc] == 1 and not self.dfs(grid1, grid2, rr, cc):
                # 如果继续遍历的结果发现不是子岛，则当前也不是子岛
                res = False

        return res

    def is_ok(self, m, n, r, c) -> bool:
        return 0 <= r < m and 0 <= c < n
