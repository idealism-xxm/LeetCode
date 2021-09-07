# 链接：https://leetcode.com/problems/find-all-groups-of-farmland/
# 题意：给定一个 01 矩阵， 0 表示林地， 1 表示农田，
#       且所有相邻的林地和农田都会组成一个矩形，求所有农田左上角和右下角的下标？

# 数据限制：
#   m == land.length
#   n == land[i].length
#   1 <= m, n <= 300
#   land 由 0 和 1 组成
#   相邻的农田会组成一个矩形

# 输入： land = [[1,0,0],[0,1,1],[0,1,1]]
# 输出： [[0,0,0,0],[1,1,2,2]]
# 解释： 
#       100
#       011
#       011
#
#       第一组农田： 左上角 (0, 0) ，右下角 (0, 0)
#       第二组农田： 左上角 (1, 1) ，右下角 (2, 2)

# 输入： land = [[1,1],[1,1]]
# 输出： [[0,0,1,1]]
# 解释： 
#       11
#       11
#
#       第一组农田： 左上角 (0, 0) ，右下角 (1, 1)

# 输入： land = [[0]]
# 输出： []


# 思路： DFS
#
#       遍历矩阵，从坐上开始找到一个 1 ，这里就是一个农田矩阵的左上角，
#       然后 DFS 找到农田矩阵的右下角，并标记遍历过的所有点
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(m * n)


direction = [(-1, 0), (0, 1), (1, 0), (0, -1)]


class Solution:
    def findFarmland(self, land: List[List[int]]) -> List[List[int]]:
        ans = []
        # 从左上角开始遍历
        for r in range(len(land)):
            for c in range(len(land[0])):
                # 如果是 1 ，则 (r, c) 是矩阵的左上角
                if land[r][c] == 1:
                    # dfs 找到矩阵的右下角
                    er, ec = self.dfs(land, r, c)
                    # 放入结果数组中
                    ans.append([r, c, er, ec])
        return ans
    
    def dfs(self, land: List[List[int]], r: int, c: int) -> Tuple[int, int]:
        m, n = len(land), len(land[0])
        # 左上角也可能是右下角
        er, ec = r, c
        # 四方向遍历
        for dr, dc in direction:
            rr, cc = r + dr, c + dc
            # 如果该位置还在矩阵内 且 是农田，则递归调用
            if 0 <= rr < m and 0 <= cc < n and land[rr][cc] == 1:
                # 标记已访问过
                land[rr][cc] = -1
                # 递归找到矩阵的右下角，然后更新
                eer, eec = self.dfs(land, rr, cc)
                er = max(er, eer)
                ec = max(ec, eec)
        return er, ec
