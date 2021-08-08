# 链接：https://leetcode.com/problems/check-if-move-is-legal/
# 题意：给定一个 8 * 8 的二维数组 board ，其中 board[i][j] 是 '.', 'B', 'W' 其中之一，
#       '.' 表示未上色， 'B' 表示已上白色， 'W' 表示已上黑色。
#       现在将 (r, c) 染成 color ，
#       判断以 (r, c) 为起点的 8 个方向中，是否有一段长度大于 2 的连续格子，
#       使得该段末尾颜色是 color ，其他部分都是已上色的另一种颜色。

# 数据限制：
#   board.length == board[r].length == 8
#   0 <= rMove, cMove < 8
#   board[rMove][cMove] == '.'
#   color 是 'B' 或 'W'

# 输入： board = [[".",".",".","B",".",".",".","."],[".",".",".","W",".",".",".","."],[".",".",".","W",".",".",".","."],[".",".",".","W",".",".",".","."],["W","B","B",".","W","W","W","B"],[".",".",".","B",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","W",".",".",".","."]], rMove = 4, cMove = 3, color = "B"
# 输出： true
# 解释： 
#       ...B....
#       ...W....
#       ...W....
#       ...W....
#       WBB?WWWB
#       ...B....
#       ...B....
#       ...W....
#
#       问号处会填入 B ，那么以其为起点向上和想有都有满足题意的连续格子

# 输入： board = [[".",".",".",".",".",".",".","."],[".","B",".",".","W",".",".","."],[".",".","W",".",".",".",".","."],[".",".",".","W","B",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".","B","W",".","."],[".",".",".",".",".",".","W","."],[".",".",".",".",".",".",".","B"]], rMove = 4, cMove = 4, color = "W"
# 输出： false
# 解释： 
#       ........
#       .B..W...
#       ..W.....
#       ...WB...
#       ....?...
#       ....BW..
#       ......W.
#       .......B
#
#       问号处会填入 W ，以其为起点的 8 个方向中，没有满足题意的连续格子


# 思路： 模拟
#
#       按照题意遍历 8 个方向，
#       每个方向都不断往前走，
#       1. 如果当前格子是 '.' ，那么直接返回 False ，表示当前方向不满足题意
#       2. 如果当前格子是 'B' ，那么继续往前走
#       3. 如果当前格子是 'W' ，那么已经到达连续格子的末尾：
#           (1) 此时长度大于等于 3 ，返回 True ，表示当前方向满足题意
#           (2) 此时长度小于 3 ，返回 False ，表示当前方向不满足题意
#       最后遍历完成后，直接返回 False ，表示当前方向不满足题意
#
#       时间复杂度： O(1)
#       空间复杂度： O(1)

class Solution:
    def checkMove(self, board: List[List[str]], r: int, c: int, color: str) -> bool:
        dr = [-1, -1, 0, 1, 1, 1, 0, -1]
        dc = [0, 1, 1, 1, 0, -1, -1, -1]
        # 遍历 8 个方向，如果任何一个方向满足题意，则返回 True
        for ddr, ddc in zip(dr, dc):
            if self.is_legal(board, r, c, color, ddr, ddc):
                return True
        return False
    
    def is_legal(self, board: List[List[str]], r: int, c: int, color: str, dr: int, dc: int) -> bool:
        # 算上 r, c 为起点的格子
        cnt = 1
        m, n = len(board), len(board[0])
        # 先走一步
        r, c = r + dr, c + dc
        # 如果还在范围内，则可以一直往前走
        while 0 <= r < m and 0 <= c < n:
            # 格子数 +1
            cnt += 1
            # 如果是空白格子，则不满足题意，直接返回 False
            if board[r][c] == '.':
                return False
            # 如果是 color ，则已到结尾处，需要判断长度是否大于等于 3
            if board[r][c] == color:
                return cnt >= 3
            r, c = r + dr, c + dc
        
        return False
