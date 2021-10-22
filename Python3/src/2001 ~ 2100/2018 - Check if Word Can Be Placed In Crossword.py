# 链接：https://leetcode.com/problems/check-if-word-can-be-placed-in-crossword/
# 题意：给定一个 m * n 的矩阵 board ，表明一个填字游戏的当前状态，
#       board 仅由英文小写字母， ' ' 和 '#' 组成，
#       ' ' 表示可以填入字母，'#' 表示不能填入字母。
#       现在判断给定的单词 word 是否能在满足以下条件下被填入矩阵中？
#       条件：
#       1. '#' 不能被填入；
#       2. ' ' 可以填入一个字母；
#       3. 如果 word 可以被横向填入，既可以从左往右填入，也可以从右往左填入，
#           但要保证 word 左右两侧都是没有 ' ' 和字母；
#       4. 如果 word 可以被纵向填入，既可以从上往下填入，也可以从下往上填入，
#           但要保证 word 上下两侧都是没有 ' ' 和字母；

# 数据限制：
#   m == board.length
#   n == board[i].length
#   1 <= m * n <= 2 * 10 ^ 5
#   board[i][j] 是 ' ', '#', 或者一个英文小写字母
#   1 <= word.length <= max(m, n)
#   word 只含有英文小写字母

# 输入： board = [["#", " ", "#"], [" ", " ", "#"], ["#", "c", " "]], word = "abc"
# 输出： true
# 解释：
#   word 可以被纵向从上往下填入第二列

# 输入： board = [[" ", "#", "a"], [" ", "#", "c"], [" ", "#", "a"]], word = "ac"
# 输出： false
# 解释：
#   无论如何填入，最后 word 的两侧都会有 ' ' 或者字母

# 输入： board = [["#", " ", "#"], [" ", " ", "#"], ["#", " ", "c"]], word = "ca"
# 输出： true
# 解释：
#   "ca" 可以被横向从右往左填入右下角


# 思路： 枚举
#
#       数据量看起来很唬人，但只要直接暴力即可，
#       我们枚举所有点，如果是 '#' ，则它的四个方向是可能被填入的，
#       枚举判断每个方向是否能按题意填入即可，
#       
#       开始寻找边界点的时间复杂度： O(m * n)
#       然后判断是否能填入时，每个 字母 和 ' ' 最多只会被访问四次，
#       所以时间复杂度仍是 O(m * n)
#
#       时间复杂度： O(m * n)
#       空间复杂度： O(1)


DIR = ((0, 1), (0, -1), (1, 0), (-1, 0))


class Solution:
    def placeWordInCrossword(self, board: List[List[str]], word: str) -> bool:
        # 获取矩阵大小
        m, n = len(board), len(board[0])
        first_ch = word[0]

        # 判断一个位置是否在矩阵中
        def is_ok(r: int, c: int) -> bool:
            return 0 <= r < m and 0 <= c < n

        # 判断从边界点 (r, c) 开始，按照 (dr, dc) 方向是否能填入 word
        def can_place(r: int, c: int, dr: int, dc: int) -> bool:
            # 如果前一个位置不是边界点，则直接返回 False
            if is_ok(r - dr, c - dc) and board[r - dr][c - dc] != '#':
                return False
            # 开始填入单词
            for ch in word:
                # 如果这个位置在矩阵外 或 这个位置既不是 ' ' 也不是 ch ，
                # 则此时不能填入，直接返回 False
                if not is_ok(r, c) or (board[r][c] != ' ' and board[r][c] != ch):
                    return False
                # 走到下一个位置
                r, c = r + dr, c + dc
            # 已经填入全部 word ，判断当前位置是否处于边界点
            return not is_ok(r, c) or board[r][c] == '#'

        # 枚举边界点
        for r, row in enumerate(board):
            for c, ch in enumerate(row):
                # 如果不是 ' ' ，则不是边界点，直接跳过
                if ch == ' ' or ch == first_ch:
                    # 此时可能能合法填入，判断是四个方向是否能够填入
                    for dr, dc in DIR:
                        # 如果能够填入 word ，则直接返回 True
                        if can_place(r, c, dr, dc):
                            return True
        # 所有情况都不能填入，返回 False
        return False
