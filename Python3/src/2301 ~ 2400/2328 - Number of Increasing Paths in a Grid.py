# 链接：https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/
# 题意：给定一个 m * n 的矩阵 grid ，求所有严格上升的路径数量？
#
#      严格上升的路径可以从任意单元格开始，到单元格结束。
#      对于每个单元格，每次可以向上下左右四个方向移动，不能对角移动或越界。


# 数据限制：
#  m == grid.length
#  n == grid[i].length
#  1 <= m, n <= 1000
#  1 <= m * n <= 10 ^ 5
#  1 <= grid[i][j] <= 10 ^ 5


# 输入： grid = [[1,1],[3,4]]
# 输出： 8
# 解释： 严格上升的路径有：
#          长度为 1: [1], [1], [3], [4]
#          长度为 2: [1 -> 3], [1 -> 4], [3 -> 4]
#          长度为 3: [1 -> 3 -> 4]

# 输入： grid = [[1],[2]]
# 输出： 3
# 解释： 严格上升的路径有：
#          长度为 1: [1], [2]
#          长度为 2: [1 -> 2]


# 思路： DP + 排序
#
#      本思路与 LeetCode 329 一致，稍作修改就可以直接复用。
#
#      本题其实不容易想到 DP ，
#      因为常见的 DP 都可以直接按照题目给定数据的顺序处理，
#      例如：位置上从左到右，从上到下，从左上到右下等。
#
#      本题支持 4 个方向转移状态，所以不能直接用 DP 来解决，
#      但本题也加了其他限制，所以我们可以自定顺序，
#      即可以通过值从小到大的顺序处理，来进行状态转移。
#
#      设 dp[r][c] 表示以 (r, c) 为终点的严格上升路径的数量，
#      由于上升路径的中的值是严格单调递增的，
#      所以我们只要保证在处理 (r, c) 之前，
#      所有值小于 grid[r][c] 的单元格都已处理完成，
#      那么 dp[r][c] 就能通过相邻的单元格的 dp 值转移得到。
#
#      grid 不会改变，所以我们最开始就将 grid 转成单元格数组 cells ，
#      其中 cells[i] = (grid[r][c], r, c) ，
#      然后按照 grid[r][c] 升序排序。
#
#      初始化令所有 dp[r][c] = 1 ，即都能以自身为终点。
#
#      然后遍历 cells 进行状态转移，此时能保证在处理 (r, c) 之前，
#      所有值小于 grid[r][c] 的单元格都已处理完成。
#
#
#      时间复杂度：O(mn * log(mn))
#          1. 需要收集 grid 中全部 O(mn) 个元素到数组 cells 中
#          2. 需要对数组 cells 进行排序，时间复杂度为 O(mn * log(mn))
#          3. 需要遍历数组 cells 全部 O(mn) 个元素
#      空间复杂度：O(mn)
#          1. 需要收集 grid 中全部 O(mn) 个元素到数组 cells 中
#          2. 需要维护一个大小为 O(mn) 的二维数组 dp


# 每个方向的位置改变量
#  0: 向上走 1 步
#  1: 向右走 1 步
#  2: 向下走 1 步
#  3: 向左走 1 步
DIRS: List[Tuple[int, int]] = [(-1, 0), (0, 1), (1, 0), (0, -1)]
MOD = 1000000007


class Solution:
    def countPaths(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        # 将矩阵转换为单元格数组
        cells: List[Tuple(int, int, int)]= [
            (grid[r][c], r, c) 
            for r in range(m) 
            for c in range(n)
        ]
        # 按照单元格的值从小到大排序
        cells.sort()

        # ans 维护所有严格上升路径的数量
        ans: int = 0
        # dp[r][c] 表示以 (r, c) 为终点的严格上升路径的数量
        # 初始化都为 1 ，即都能以自身为终点
        dp: List[List[int]] = [[1] * n for _ in range(m)]
        # 按照单元格的值从小到大遍历，
        # 保证处理 (r, c) 时，值小于 v 的单元格已经处理过了
        for v, r, c in cells:
            for dr, dc in DIRS:
                # 计算相邻单元格的坐标
                rr: int = r + dr
                cc: int = c + dc
                if (
                    0 <= rr < m and 
                    0 <= cc < n and 
                    grid[rr][cc] < v
                ):
                    # 如果 (rr, cc) 的值比 (r, c) 的值小，
                    # 则 (r, c) 的前一个单元格可以时 (rr, cc) ，
                    # 即 dp[r][c] 可以由 dp[rr][cc] 转移而来
                    dp[r][c] = (dp[r][c] + dp[rr][cc]) % MOD

            # 更新 ans
            ans = (ans + dp[r][c]) % MOD

        return ans
