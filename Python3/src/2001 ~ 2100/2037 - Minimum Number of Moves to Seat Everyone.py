# 链接：https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
# 题意：给定两个长度为 n 的数组 seats 和 students ，其中 seats[i] 表示第 i 个座位的位置，
#       students[j] 表示第 j 个学生的位置，
#       现在每次可以将第 i 个学生移动至 students[j] - 1 或 students[j] + 1 ，
#       求最少移动多少次后，所有的学生都有座位？
#       （注意初始可能有多个学生或者座位在同一个位置。）

# 数据限制：
#   n == seats.length == students.length
#   1 <= n <= 100
#   1 <= seats[i], students[j] <= 100

# 输入： seats = [3,1,5], students = [2,7,4]
# 输出： 4
# 解释：
#   第 0 个学生从 2 移动至 1 ，需要移动 1 次
#   第 1 个学生从 7 移动至 5 ，需要移动 2 次
#   第 2 个学生从 4 移动至 3 ，需要移动 1 次
#   总共需要移动： 1 + 2 + 1 = 4 次 

# 输入： seats = [4,1,5,9], students = [1,3,2,6]
# 输出： 7
# 解释：
#   第 0 个学生不移动
#   第 1 个学生从 3 移动至 4 ，需要移动 1 次
#   第 2 个学生从 2 移动至 5 ，需要移动 3 次
#   第 3 个学生从 6 移动至 9 ，需要移动 3 次
#   总共需要移动： 0 + 1 + 3 + 3 = 7 次 

# 输入： seats = [2,2,6,6], students = [1,3,2,6]
# 输出： 4
# 解释：
#   第 0 个学生从 1 移动至 2 ，需要移动 1 次
#   第 1 个学生从 3 移动至 6 ，需要移动 3 次
#   第 2 个学生不移动
#   第 3 个学生不移动
#   总共需要移动： 1 + 3 + 0 + 0 = 4 次 


# 思路： 排序
#
#       我们对座位和学生都按照位置进行排序，那么第 i 个学生需要从 students[i] 移动到 seats[i] ，
#       所有需要移动的次数为： sum(abs(students[i] - seats[i]))
#
#       时间复杂度： O(nlogn)
#       空间复杂度： O(1)


class Solution:
    def minMovesToSeat(self, seats: List[int], students: List[int]) -> int:
        # 座位和学生均按照位置排序
        seats.sort()
        students.sort()
        # 对座位和学生都按照位置进行排序后，
        # 需要移动的次数为： sum(abs(students[i] - seats[i]))
        return sum(abs(student - seat) for seat, student in zip(seats, students))
