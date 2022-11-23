# 链接：https://leetcode.com/problems/valid-sudoku/
# 题意：判断一个数独的初始状态是否合法，合法状态如下：
#          1. 数字 1-9 在每一行只能出现一次
#          2. 数字 1-9 在每一列只能出现一次
#          3. 数字 1-9 在每一个九宫格只能出现一次


# 数据限制：
#  board.length == 9
#  board[i].length == 9
#  board[i][j] 是数字 1-9 或 '.'


# 输入： board = 
#       [["5","3",".",".","7",".",".",".","."]
#       ,["6",".",".","1","9","5",".",".","."]
#       ,[".","9","8",".",".",".",".","6","."]
#       ,["8",".",".",".","6",".",".",".","3"]
#       ,["4",".",".","8",".","3",".",".","1"]
#       ,["7",".",".",".","2",".",".",".","6"]
#       ,[".","6",".",".",".",".","2","8","."]
#       ,[".",".",".","4","1","9",".",".","5"]
#       ,[".",".",".",".","8",".",".","7","9"]]
# 输出： true

# 输入： board = 
#       [["8","3",".",".","7",".",".",".","."]
#       ,["6",".",".","1","9","5",".",".","."]
#       ,[".","9","8",".",".",".",".","6","."]
#       ,["8",".",".",".","6",".",".",".","3"]
#       ,["4",".",".","8",".","3",".",".","1"]
#       ,["7",".",".",".","2",".",".",".","6"]
#       ,[".","6",".",".",".",".","2","8","."]
#       ,[".",".",".","4","1","9",".",".","5"]
#       ,[".",".",".",".","8",".",".","7","9"]]
# 输出： false
# 解释： 左上角第一个九宫格含有 2 个 8 。


# 思路： 模拟
#
#      针对每一行/列/九宫格，我们都可以使用一个集合 nums 来维护已经出现的数字。
#
#      然后遍历其中的所有数字 num ，如果 num 已在 nums 中，则不满足题意，返回 false ；
#      否则将 num 加入到 nums 中。
#
#      最后如果还未返回，则说明所有的行/列/九宫格都满足题意，直接返回 true 。
#
#
#      设数独矩阵边长为 n 。
#
#      时间复杂度：O(n ^ 2)
#          1. 需要遍历全部 O(n) 行/列/九宫格，每一行/列/九宫格都需要遍历全部 O(n) 个元素
#      空间复杂度：O(n)
#          1. 判断每一行/列/九宫格时，都需要维护全部 O(n) 个元素


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            # 如果当前行/列不合法，则直接返回 false
            if not Solution.is_valid_row(board, i) or not Solution.is_valid_col(board, i):
                return False

            # 计算第 i 个九宫格的左上角坐标
            r, c = (i // 3) * 3, (i % 3) * 3
            # 如果该九宫格不合法，则直接返回 false
            if not Solution.is_valid_square(board, r, c):
                return False

        return True

    @staticmethod
    def is_valid_row(board: List[List[str]], r: int) -> bool:
        # nums 维护当前行已经出现的数字
        nums: Set[str] = set()
        for c in range(9):
            # 如果是空，则继续
            if board[r][c] == '.':
                continue

            # 如果当前数字已出现，则当前行不合法
            if board[r][c] in nums:
                return False

            # 标记当前数字已出现
            nums.add(board[r][c])

        # 所有数字否不重复，则当前行合法
        return True

    @staticmethod
    def is_valid_col(board: List[List[str]], c: int) -> bool:
        # nums 维护当前列已经出现的数字
        nums: Set[str] = set()
        for r in range(9):
            # 如果是空，则继续
            if board[r][c] == '.':
                continue

            # 如果当前数字已出现，则当前列不合法
            if board[r][c] in nums:
                return False

            # 标记当前数字已出现
            nums.add(board[r][c])

        # 所有数字否不重复，则当前列合法
        return True

    @staticmethod
    def is_valid_square(board: List[List[str]], start_r: int, start_c: int) -> bool:
        # nums 维护当前列已经出现的数字
        nums: Set[str] = set()
        for r in range(start_r, start_r + 3):
            for c in range(start_c, start_c + 3):
                # 如果是空，则继续
                if board[r][c] == '.':
                    continue

                # 如果当前数字已出现，则当前九宫格不合法
                if board[r][c] in nums:
                    return False

                # 标记当前数字已出现
                nums.add(board[r][c])

        # 所有数字否不重复，则当前九宫格合法
        return True
