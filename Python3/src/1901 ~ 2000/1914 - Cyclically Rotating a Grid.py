# 链接：https://leetcode.com/problems/cyclically-rotating-a-grid/
# 题意：给定一个二维数组，将每一层的数字逆时针旋转 k 位。

# 数据限制：
#   m == grid.length
#   n == grid[i].length
#   2 <= m, n <= 50
#   m 和 n 都是偶数
#   1 <= grid[i][j] <= 5000
#   1 <= k <= 10 ^ 9

# 输入： grid = [[40,10],[30,20]], k = 1
# 输出： [[10,20],[40,30]]

# 输入： grid = [[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]], k = 2
# 输出： [[3,4,8,12],[2,11,10,16],[1,7,6,15],[5,9,13,14]]

# 思路： 模拟
#
#       对每一层的数，模拟旋转 k 次即可，
#       先找到每一层的第一个位置 (sr, sc) 对应旋转 k 次后的位置 (er, ec)，
#       然后每次移动到下一个位置 (sr, sc) 和其对应的位置 (er, ec)，
#       将该层所有的数复制到结果数组对应位置即可
#
#       注意要先用 k 对这一层的长度取模
#
#
#       【另一种更简单的模拟方式】
#       结束后还看到其他模拟的方式，实现起来更简单一点，
#       就是先将每一层的每个数位置信息 (r * n + c) 按照顺序放到一个数组 pos 中，
#       则移动次数变为： c = k % len(pos) ，
#       然后枚举 pos 数组的下标 i ，可以计算出对应的位置为 j = (j - c + len(pos)) % len(pos)
#       则 grid[i // n][i % n] 会移动到 ans[j // n][j % n]
#
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(m * n)


dr = [1, 0, -1, 0]
dc = [0, 1, 0, -1]


class Solution:
    def rotateGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        m, n = len(grid), len(grid[0])
        # 初始化结果二维数组
        result = [None] * m
        for i in range(m):
            result[i] = [0] * n
        # 枚举每一层，注意只要行/列有一个遍历完了，就所有层都结束了
        for layer in range(min(m, n) // 2):
            # 计算该层的坐标范围
            min_r, max_r, min_c, max_c = layer, m - layer - 1, layer, n - layer - 1
            # 计算该层数的长度
            length = 2 * (max_r - min_r + 1 + max_c - min_c + 1) - 4
            # 原位置和初始方向
            sr, sc, sdir = layer, layer, 0
            # 计算目标位置和初始方向
            er, ec, edir = self.get_end_position(sr, sc, sdir, min_r, max_r, min_c, max_c, k % length)
            # 将所有的位置复制到对应的结果数组中
            for i in range(length):
                result[er][ec] = grid[sr][sc]
                # 计算原数组中下一个位置及方向
                sr, sc, sdir = self.get_next_position(sr, sc, sdir, min_r, max_r, min_c, max_c)
                # 计算目标数组中下一个位置及方向
                er, ec, edir = self.get_next_position(er, ec, edir, min_r, max_r, min_c, max_c)
        return result

    def get_end_position(self, r, c, direction, min_r, max_r, min_c, max_c, k):
        # 走 k 次之后即可达到目标位置
        while k > 0:
            r, c, direction = self.get_next_position(r, c, direction, min_r, max_r, min_c, max_c)
            k -= 1
        return r, c, direction

    def get_next_position(self, r, c, direction, min_r, max_r, min_c, max_c):
        # 先以当前方向走到下一个位置
        rr = r + dr[direction]
        cc = c + dc[direction]
        # 如果当前位置合法，则直接返回
        if self.is_ok(rr, cc, min_r, max_r, min_c, max_c):
            return rr, cc, direction
        # 如果当前位置不合法，则需要改变到下一个方向
        return self.get_next_position(r, c, (direction + 1) % 4, min_r, max_r, min_c, max_c)

    def is_ok(self, r, c, min_r, max_r, min_c, max_c):
        return min_r <= r <= max_r and min_c <= c <= max_c