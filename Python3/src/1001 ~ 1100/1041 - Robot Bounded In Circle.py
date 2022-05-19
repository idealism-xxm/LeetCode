#  链接：https://leetcode.com/problems/robot-bounded-in-circle/
#  题意：在一个无限的平面上，一个机器人初始在 (0, 0) 处，面朝北方，
#       现在给定一个字符串 instructions ，表示一连串三种行动指令：
#           "G": 向现在面朝的方向走 1 步
#           "L": 向左转 90 度
#           "R": 向右转 90 度
#       这个机器人按照 instructions 的指令行动，并且会无限重复，
#       现在判断这个机器人的行动路径是否是一个圈？

#  数据限制：
#   1 <= instructions.length <= 100 
#   instructions[i] 是 'G', 'L' 或 'R'

#  输入： instructions = "GGLLGG"
#  输出： true
#  解释： 机器人先往北走 2 步到 (-2, 0) ，然后转 180 度，再往南走 2 步到 (0, 0) 。
#        然后这个机器人一直重复这个路径，
#        在 (0, 0) -> (-2, 0) -> (0, 0) -> (2, 0) -> ... 之间循环。

#  输入： instructions = "GG"
#  输出： false
#  解释： 机器人不会转向，一直往北方走，不会回到原点。

#  输入： instructions = "GL"
#  输出： true
#  解释： 机器人一直重复以下路径： (0, 0) -> (0, 1) -> (-1, 1) -> (-1, 0) -> (0, 0) -> ...

#  思路： 模拟
# 
#       1. 如果经过一次 instructions 后，机器人回到 (0, 0) ，则必定能形成一个圈，如此反复，返回 true
#       2. 如果经过一次 instructions 后，机器人没有回到 (0, 0) ，
#           此时根据机器人的面朝方向有以下情况：
#           (1) 北方：后续的 instructions 必定还会不断远离 (0, 0) ，无法形成一个圈，返回 false
#           (2) 南方：则在进行 1 次 instructions 后，机器人必定会回到 (0, 0) ，可以形成一个圈，返回 true
#           (3) 东方：则在进行 3 次 instructions 后，机器人必定会回到 (0, 0) ，可以形成一个圈，返回 true
#                   因为执行过北方和南方的两次 instructions 后，机器人在这两个方向上的偏离会抵消
#                       执行过东方和西方的两次 instructions 后，机器人在这两个方向上的偏离会抵消
#           (4) 西方：则在进行 3 次 instructions 后，机器人必定会回到 (0, 0) ，可以形成一个圈，返回 true
#                   因为执行过北方和南方的两次 instructions 后，机器人在这两个方向上的偏离会抵消
#                       执行过东方和西方的两次 instructions 后，机器人在这两个方向上的偏离会抵消
# 
#       时间复杂度： O(n)
#       空间复杂度： O(1)

#  每个方向的位置改变量
#   0: 向北走 1 步
#   1: 向东走 1 步
#   2: 向南走 1 步
#   3: 向西走 1 步
DIRS = ((-1, 0), (0, 1), (1, 0), (0, -1))

class Solution:
    def isRobotBounded(self, instructions: str) -> bool:
        # 初始机器人在 (0, 0) ，方向 0 表示面朝北
        x, y, direction = (0, 0, 0)
        # 遍历操作指令
        for instruction in instructions:
            if instruction == 'G':
                # 获取当前方向的位置改变量
                dx, dy = DIRS[direction]
                # 机器人朝该方向走 1 步，方向不变
                x += dx
                y += dy
            elif instruction == 'L':
                # 如果指令是左转，则更新方向为 (dir + 3) % 4
                direction = (direction + 3) % 4
            else:
                # 如果指令是右转，则更新方向为 (dir + 1) % 4
                direction = (direction + 1) % 4

        #  最后机器人能形成一个圈的情况：
        #   1. 机器人还在 (0, 0)
        #   2. 机器人不是面朝北方
        return (x == 0 and y == 0) or direction != 0
