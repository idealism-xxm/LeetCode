# 链接：https://leetcode.com/problems/maximum-compatibility-score-sum/
# 题意：给定两个 m * n 的 01 二维数组 students 和 mentors ，
#       分别表示学生和导师对某 n 题的解答，
#       如果一个学生的解答和一个导师的某题解答相同，则这个学生的分数 +1 ，
#       每个学生和导师只能配对一次，
#       求如何分配学生和导师，才能使得所有学生的分数和最大？


# 数据限制：
#   m == students.length == mentors.length
#   n == students[i].length == mentors[j].length
#   1 <= m, n <= 8
#   students[i][k] 是 0 或 1
#   mentors[j][k] 是 0 或 1

# 输入： students = [[1,1,0],[1,0,1],[0,0,1]], mentors = [[1,0,0],[0,0,1],[1,1,0]]
# 输出： 8
# 解释： 
#       学生 0 分配给导师 2 ，得分为 3
#       学生 1 分配给导师 0 ，得分为 2
#       学生 2 分配给导师 1 ，得分为 3

# 输入： students = [[0,0],[0,0],[0,0]], mentors = [[1,1],[1,1],[1,1]]
# 输出： 0
# 解释： 
#       无论如何分配，学生的分数都是 0


# 思路： 状压 DP
#
#       学生和导师数很小，所以可以使用状态压缩。
#       比赛时使用了最简单的方法，同时考虑了学生和导师的顺序，
#       所以时间复杂度为 O(m * (2 ^ (2m))) ，空间复杂度为 O(2 ^ (2m)) 。
#       设 dp[i][j] 表示已选学生状态 i 和已选导师状态 j 时的最大分数。
#
#       这样存在状态重复，因为两者都是全部顺序，我们可以固定学生的顺序再次简化，
#       设 dp[i] 表示已选导师状态为 i 时的最大分数，
#       （此时已选的学生是前 get_one_count(i) 个）
#
#       初始化： dp[i] = 0
#       状态转移：
#           遍历所有导师 j ，如果该导师在 i 中未被选用，则可以和学生 get_one_count(i) 配对，
#           即 dp[i | (1 << j)] = max(dp[i | (1 << j)], dp[i] + score[get_one_count(i)][j])
#
#       其中 score[get_one_count(i)][j] 表示学生 get_one_count(i) 和导师 j 配对时能获得的分数
#
#       时间复杂度： O(m * (2 ^ m))
#       空间复杂度： O(2 ^ m)

class Solution:
    def maxCompatibilitySum(self, students: List[List[int]], mentors: List[List[int]]) -> int:
        # 计算学生和老师的配对分数
        score = [
            [
                sum((student[k] == mentor[k] for k in range(len(student)))) 
                for mentor in mentors
            ]
            for student in students
        ]
        # 初始化状态
        mx = 1 << len(students)
        dp = [0] * (mx + 1)
        for i in range(mx):
            # 遍历当前的导师
            for j in range(len(mentors)):
                # 如果该导师在 i 中未被选用，则可以和学生 get_one_count(i) 配对
                if not (i & (1 << j)):
                    dp[i | (1 << j)] = max(dp[i | (1 << j)], dp[i] + score[self.get_one_count(i)][j])
            
        return dp[mx - 1]

    def get_one_count(self, num: int) -> int:
        """计算一个数中二进制下 1 的个数"""
        cnt = 0
        while num > 0:
            cnt += num & 1
            num >>= 1
        return cnt
