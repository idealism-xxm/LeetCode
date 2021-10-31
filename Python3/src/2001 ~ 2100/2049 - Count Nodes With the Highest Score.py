# 链接：https://leetcode.com/problems/count-nodes-with-the-highest-score/
# 题意：给定一个二叉树，用 parents 表示每个节点的父节点，
#       parents[i] = j 表示节点 i 的父节点是节点 j ,
#       parents[0] = -1 表示根节点时节点 0 。
#       每个节点的得分计算方式为：取出这个节点后，剩余的部分形成的树的大小之积。
#       求最高得分的节点有多少个？

# 数据限制：
#   n == parents.length
#   2 <= n <= 10 ^ 5
#   parents[0] == -1
#   0 <= parents[i] <= n - 1 for i != 0
#   parents 表示一颗合法的二叉树

# 输入： parents = [-1,2,0,2,0]
# 输出： 3
# 解释： 
#   - 节点 0 的得分： 3 * 1 = 3
#   - 节点 1 的得分： 4 = 4
#   - 节点 2 的得分： 1 * 1 * 2 = 2
#   - 节点 3 的得分： 4 = 4
#   - 节点 4 的得分： 4 = 4
#   最高得分是 4 ，总共有 3 个节点

# 输入： parents = [-1,2,0]
# 输出： 2
# 解释： 
#   - 节点 0 的得分： 2 = 2
#   - 节点 1 的得分： 2 = 2
#   - 节点 2 的得分： 1 * 1 = 1
#   最高得分是 2 ，总共有 2 个节点


# 思路： DFS
#
#       我们维护一个 map ， key 为节点得分， value 为该得分的节点数。
#       那么我们直接 dfs 后续遍历计算每个节点的得分即可。
#
#       每个节点得分计算方式很简单：
#           (1) 统计所有子树的大小，然后将它们乘起来
#           (2) 如果父节点部分还有节点，那么还要乘以剩余节点的数量
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def countHighestScoreNodes(self, parents: List[int]) -> int:
        # 记录所有节点数量
        n = len(parents)
        # 初始化邻接表
        edges = collections.defaultdict(list)
        for child in range(1, n):
            edges[parents[child]].append(child)

        # 记录每个得分的节点数
        cnt = defaultdict(int)
        # dfs 后续遍历计算每个节点的得分
        def dfs(cur: int) -> int:
            # res 表示当前节点的得分
            # sze 表示当前节点的子树大小
            res, sze = 1, 1
            # 遍历子节点递归
            for child in edges[cur]:
                # 获取子节点的子树大小
                sub_sze = dfs(child)
                # 当前节点的子树大小 需要加上 子节点的子树大小
                sze += sub_sze
                # 当前节点得分 需要诚意 子节点的子树大小
                res *= sub_sze

            # 如果父节点部分还有节点，则还需要乘以剩余节点的数量
            if n - sze > 0:
                res *= n - sze

            # 当前节点得分的节点数 +1
            cnt[res] += 1
            # 返回当前节点的子树大小
            return sze

        # 调用计算
        dfs(0)

        # 找到最大的得分
        ans = max(cnt.keys())
        # 返回最大的分的节点数
        return cnt[ans]
