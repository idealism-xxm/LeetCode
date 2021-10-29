# 链接：https://leetcode.com/problems/parallel-courses-iii/
# 题意：有 n 节课需要学习，给定二维数组 relations 表示课程之间的依赖关系，
#       其中 relations[j] = [preCourse_j, nextCourse_j] 表示课程 preCourse_j 必须在课程 nextCourse_j 前完成，
#       给定长度为 n 的数组 time ，其中 time[i] 表示完成第 i + 1 门课程所需的月数，
#       求完成所有课程的最小月数？

# 数据限制：
#   1 <= n <= 5 * 10 ^ 4
#   0 <= relations.length <= min(n * (n - 1) / 2, 5 * 10 ^ 4)
#   relations[j].length == 2
#   1 <= prevCourse_j, nextCourse_j <= n
#   prevCourse_j != nextCourse_j
#   所有的元祖 [prevCourse_j, nextCourse_j] 都是唯一的
#   time.length == n
#   1 <= time[i] <= 10 ^ 4
#   给定的图不含自环

# 输入： n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
# 输出： 8
# 解释： 
#   第 0 月同时开始课程 1 和课程 2
#   课程 1 花费 3 个月学习完成，课程 2 花费 2 个月学习完成
#   因此，最早可以在第 3 月开始课程 3 ，总共用时 3 + 5 = 8 个月

# 输入： n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
# 输出： 12
# 解释： 
#   第 0 月同时开始课程 1 ，课程 2 和课程 3
#   课程 1 花费 1 个月学习完成，课程 2 花费 2 个月学习完成，课程 3 花费 3 个月学习完成
#   第 3 月开始课程 4 ，花费 4 个月学习完成，
#   第 7 月开始课程 5 ，花费 5 个月学习完成，
#   总共用时 3 + 4 + 5 = 12 个月


# 思路： 拓扑排序 + BFS
#
#       比赛中想到拓扑排序，然后不断将入度为 0 的点丢到优先队列中，
#       这样最后一个完成课程的月数就是完成所有课程的最小月数。
#
#       这样时间复杂度是 O(nlogn + m)
#
#       比赛后发现有题解直接用拓扑排序 + BFS 就能将时间复杂度降低为 O(n + m) ，
#       因为考虑到一点：完成所有课程的最小月数 = 入度为 0 的点到出度为 0 的点的所有路径中的最大值
#
#       所以我们可以用 dist 数组维护这个值，即 dist[i] 表示完成课程 i 的最早时间，
#       遍历方式还是使用拓扑排序，只讲入度为 0 的值放入到队列中，
#       每次更新 dist 的最大值，保证 dist[i] 是到课程 i 所有路径中的最大值
#
#       时间复杂度： O(n + m)
#       空间复杂度： O(n + m)


class Solution:
    def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
        # 入度表，初始都为 0
        ind = [0] * (n + 1)
        # 初始化邻接表，转换成从 0 开始的下标
        edges = defaultdict(list)
        for pre, nxt in relations:
            # nxt 入度 +1
            ind[nxt - 1] += 1
            # 加入有向边
            edges[pre - 1].append(nxt - 1)

        # 初始化距离数组，即完成每个可能所需的最早时间
        # 注意：这个最早时间是所有到该课程的路径中最长的一个
        dist = [0] * n
        q = deque([])
        # 先讲入度为 0 的课程放入
        for i in range(n):
            if ind[i] == 0:
                q.append(i)
                # 第 0 月开始的课程，完成时间就是 time[i]
                dist[i] = time[i]

        # 当队列不为空时，进行遍历
        while len(q):
            # 拿出队首的课程
            cur = q.popleft()
            # 遍历以来 cur 的所有课程 nxt
            for nxt in edges[cur]:
                # 更新完成课程 nxt 的最早时间，
                # 这个时间是所有路径中的最大值
                dist[nxt] = max(dist[nxt], dist[cur] + time[nxt])
                # nxt 入度 -1
                ind[nxt] -= 1
                # 如果此时 nxt 入度为 0 ，
                # 则此时课程 nxt 依赖的所有课程都已完成， dist 值也都确定了，
                # 可以开始上课程 nxt
                if ind[nxt] == 0:
                    q.append(nxt)

        # 所有课程完成的最早时间的最大值，就是完成所有课程的最大时间
        return max(dist)
