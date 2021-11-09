# 链接：https://leetcode.com/problems/number-of-valid-move-combinations-on-chessboard/
# 题意：在一个 8 * 8 的国际象棋棋盘上，用长度为 n 的数组 pieces 表示棋子类型列表，
#       其中 pieces[i] 表示第 i 个棋子的类型，
#       用长度为 n 的数组 positions 表示棋子的位置列表，
#       其中 positions[i] = [r_i, c_i] 表示第 i 个棋子的位置。
#       总共有三种棋子：Rook, Queen, Bishop ，
#       它们的移动方式如下：
#           1. Rook: (r, c) -> (r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)
#           2. Queen: (r, c) -> (r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1), (r + 1, c + 1), (r - 1, c + 1), (r + 1, c - 1), (r - 1, c - 1)
#           3. Bishop: (r, c) -> (r + 1, c + 1), (r - 1, c + 1), (r + 1, c - 1), (r - 1, c - 1)
#       现在你可以指定每个棋子按照一个固定的方向移动任意次，求有多少种合法的移动方式？
#       每秒只会按照选定的方向移动一格，任意时刻没有两个棋子在同一位置就是合法的。

# 数据限制：
#   n == pieces.length
#   n == positions.length
#   1 <= n <= 4
#   pieces 只含有字符串 "rook", "queen", 和 "bishop".
#   棋盘上最多只有一个 queen
#   1 <= x_i, y_i <= 8
#   每个 positions[i] 都是唯一的

# 输入： pieces = ["rook"], positions = [[1,1]]
# 输出： 15
# 解释： 
#   不动：有 1 种移动方式
#   往右：有 7 种移动方式
#   往下：有 7 种移动方式

# 输入： pieces = ["queen"], positions = [[1,1]]
# 输出： 22
# 解释： 
#   不动：有 1 种移动方式
#   往右：有 7 种移动方式
#   往下：有 7 种移动方式
#   往右下：有 7 种移动方式

# 输入： pieces = ["bishop"], positions = [[4,3]]
# 输出： 12
# 解释： 
#   不动：有 1 种移动方式
#   往右上：有 3 种移动方式
#   往右下：有 4 种移动方式
#   往左下：有 2 种移动方式
#   往左上：有 2 种移动方式

# 输入： pieces = ["rook","rook"], positions = [[1,1],[8,8]]
# 输出： 223
# 解释： 
#   每个 rook 有 15 种移动方式，共 15 * 15 = 225 种移动方式
#   有两个不合法的移动方式：
#        都移动到 (8, 1) 时，会相遇
#        都移动到 (1, 8) 时，会相遇

# 输入： pieces = ["queen","bishop"], positions = [[5,7],[3,4]]
# 输出： 281
# 解释： 
#   queen 有 12 种移动方式， bishop 有 24 种移动方式，
#   共 12 * 24 = 288 种移动方式
#   有七个不合法的移动方式：
#        queen 停在 (6, 7) ， bishop 无法移动到 (6, 7) 和 (7, 8)
#        queen 停在 (5, 6) ， bishop 无法移动到 (5, 6), (6, 7) 和 (7, 8)
#        queen 停在 (5, 2) ， bishop 无法移动到 (5, 2) 和 (5, 1)

# 思路： DFS
#
#       用 dfs 枚举每个棋子的移动方向及移动次数，然后判断这种情况是否合法，
#       如果合法则需要对统计结果 +1 。
#
#       时间复杂度： O(n * L * D * n * L * n) ， L 表示棋盘的长度， D = 3 （方向规约成 3 ） 
#       空间复杂度： O(n)


# 设定每个棋子的移动方式
DIRS = {
    "rook": [(-1, 0), (0, 1), (1, 0), (0, -1)],
    "queen": [(-1, 0), (0, 1), (1, 0), (0, -1), (-1, 1), (1, 1), (1, -1), (-1, -1)],
    "bishop": [(-1, 1), (1, 1), (1, -1), (-1, -1)],
}


class Solution:
    def countCombinations(self, pieces: List[str], positions: List[List[int]]) -> int:
        # 统计合法的结果
        ans = 0

        # 生成棋子对象
        piece_list = [Piece() for _ in range(len(pieces))]
        # dfs 遍历， i 表示该遍历的棋子下标， max_cnt 表示最大的移动次数
        def dfs(i, max_cnt):
            # 如果已经确定了所有位置的移动方式，则判断是否合法
            if i == -1:
                # 如果合法，则结果 +1
                if is_ok(piece_list, max_cnt):
                    nonlocal ans
                    ans += 1
                return
            
            # 获取棋子的初始位置
            r, c = positions[i]
            # 初始化棋子原始数据
            piece_list[i].init(r, c, 0, 0, 0)
            # 当前棋子不移动，先递归处理下一个数据
            dfs(i - 1, max_cnt)
            # 遍历当前自己的移动方向
            for dr, dc in DIRS[pieces[i]]:
                # 遍历当前棋子的移动次数
                for cnt in range(1, 8):
                    # 如果当前棋子在该放下可以移动，则递归处理下一个棋子
                    if 1 <= r + dr * cnt <= 8 and 1 <= c + dc * cnt <= 8:
                        # 初始化此时的原始数据
                        piece_list[i].init(r, c, dr, dc, cnt)
                        # 递归处理下一个数据
                        dfs(i - 1, max(max_cnt, cnt))
                    else:
                        break
        
        dfs(len(pieces) - 1, 0)
        return ans


# 判断当前棋子及其移动方式是否为合法
def is_ok(piece_list, max_cnt):
    # 如果不移动，则直接返回 True
    if not max_cnt:
        return True
    
    # 每个棋子重置为原始状态
    for p in piece_list:
        p.reset()
    
    # 如果还有棋子需要移动，则继续移动
    while max_cnt:
        max_cnt -= 1
        
        # 每个棋子移动到下一个位置
        for p in piece_list:
            p.next()

        # 如果有棋子在同一个位置，则直接返回 False
        if len(set(piece_list)) != len(piece_list):
            return False
    return True
    
    
class Piece:
    # 初始化棋子的数据
    def init(self, r, c, dr, dc, remain_move_cnt):
        # 保存原始数据，方便重置
        self._r = r
        self._c = c
        self._remain_move_cnt = remain_move_cnt
        # 棋子当前的实时数据
        self.r = r
        self.c = c
        self.dr = dr
        self.dc = dc
        self.remain_move_cnt = remain_move_cnt

    def next(self):
        # 如果还能移动，则继续移动
        if self.remain_move_cnt:
            self.r += self.dr
            self.c += self.dc
            self.remain_move_cnt -= 1
            return 1 <= self.r <= 8 and 1 <= self.c <= 8

        return True
    
    def reset(self):
        # 重置为原始数据
        self.r = self._r
        self.c = self._c
        self.remain_move_cnt = self._remain_move_cnt
    
    # 定义 set 需要的方法
    def __eq__(self, other):
        return self.r == other.r and self.c == other.c

    def __hash__(self):
        return hash((self.r, self.c))
