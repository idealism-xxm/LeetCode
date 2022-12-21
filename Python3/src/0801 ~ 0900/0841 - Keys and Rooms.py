# 链接：https://leetcode.com/problems/keys-and-rooms/
# 题意：有 n 个房间，标号为 0 到 n - 1 。
#
#      给定长度为 n 的数组 rooms ，其中 rooms[i] 表示 i 号房间内的钥匙集合。
#      rooms[i][j] 表示 i 号房间内的第 j 把钥匙能打开的房间号。
#
#      初始时只有 0 号房间未锁，其他房间都锁着，不能在没有获得钥匙时进入锁着的房间。
#      判断是否能进入全部房间？


# 数据限制：
#  n == rooms.length
#  2 <= n <= 1000
#  0 <= rooms[i].length <= 1000
#  1 <= sum(rooms[i].length) <= 3000
#  0 <= rooms[i][j] < n
#  所有的 rooms[i] 都各不相同


# 输入： rooms = [[1],[2],[3],[]]
# 输出： true
# 解释： 先访问 0 号房间，获得钥匙 1 ；
#       再访问 1 号房间，获得钥匙 2 ；
#       再访问 2 号房间，获得钥匙 3 ；
#       再访问 3 号房间。

# 输入： rooms = [[1,3],[3,0,1],[2],[0]]
# 输出： false
# 解释： 无法访问 2 号房间，
#       因为唯一一把 2 号房间的钥匙就在 2 号房间内。


# 思路： BFS
#
#      本题是 LeetCode 1971 的加强版，不再是简单的无向图，
#      增加了前后遍历的依赖，但依旧可以使用 BFS 求解。
#
#
#      直接用 BFS 从 0 号房间开始遍历，将未遍历过且获得钥匙的房间号放入队列中。
#
#      如果 BFS 结束时，遍历过的房间数为 n ，则说明进入了全部房间，返回 true ；
#      否则返回 false 。
#
#
#      时间复杂度：O(n + m)
#          1. 需要遍历全部 O(n) 个房间
#          2. 需要遍历全部 O(m) 把钥匙
#      空间复杂度：O(n)
#          1. 需要维护 visited 中全部 O(n) 个房间的访问状态
#          2. 需要维护 q 中全部 O(n) 个房间


class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        # visited 标记某个房间是否已访问
        visited: Set[int] = set()
        # 队列 q ，用于 BFS 遍历
        q: Deque = deque()
        # 初始时标记 0 号房间已访问，并放入队列中
        visited.add(0)
        q.append(0)
        # 当 q 不为空时，按照以下逻辑循环处理
        while q:
            cur: int = q.popleft()
            # 遍历与房间 cur 内的钥匙对应的房间号 nxt
            for nxt in rooms[cur]:
                # 如果 nxt 还没访问过，则标记已访问，并放入队列中
                if nxt not in visited:
                    visited.add(nxt)
                    q.append(nxt)

        # 如果遍历过的房间数为全部房间数，则说明进入了全部房间
        return len(visited) == len(rooms)
