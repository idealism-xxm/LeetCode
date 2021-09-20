# 链接：https://leetcode.com/problems/smallest-missing-genetic-value-in-each-subtree/
# 题意：给定一颗大小为 n 的树，根节点为节点 0 ，每个节点都有一个唯一的基因值，
#       返回一个长度为 n 的数组 ans ，ans[i] 表示节点 i 及其所有子节点的基因值中，
#       最小的未出现的基因值。

# 数据限制：
#   n == parents.length == nums.length
#   2 <= n <= 10 ^ 5
#   0 <= parents[i] <= n - 1 (i != 0)
#   parents[0] == -1
#   parents 表示一颗合法的树
#   1 <= nums[i] <= 10 ^ 5
#   每个 nums[i] 都是唯一的

# 输入： parents = [-1,0,0,2], nums = [1,2,3,4]
# 输出： [5,1,1,1]
# 解释：
#   - 0: 子树含有节点 [0,1,2,3] ，基因值为 [1,2,3,4] 。 5 最小的未出现的基因值
#   - 1: 子树含有节点 [1] ，基因值为 [2] 。 1 最小的未出现的基因值
#   - 2: 子树含有节点 [2,3] ，基因值为 [3,4] 。 1 最小的未出现的基因值
#   - 3: 子树含有节点 [3] ，基因值为 [4] 。 1 最小的未出现的基因值

# 输入： parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
# 输出： [7,1,1,4,2,1]
# 解释：
#   - 0: 子树含有节点 [0,1,2,3,4,5] ，基因值为 [5,4,6,2,1,3] 。 7 最小的未出现的基因值
#   - 1: 子树含有节点 [1,2] ，基因值为 [4,6] 。 1 最小的未出现的基因值
#   - 2: 子树含有节点 [2] ，基因值为 [6] 。 1 最小的未出现的基因值
#   - 3: 子树含有节点 [3,4,5] ，基因值为 [2,1,3] 。 4 最小的未出现的基因值
#   - 4: 子树含有节点 [4] ，基因值为 [1] 。 2 最小的未出现的基因值
#   - 5: 子树含有节点 [5] ，基因值为 [3] 。 1 最小的未出现的基因值

# 输入： parents = [-1,2,3,0,2,4,1], nums = [2,3,4,5,6,7,8]
# 输出： [1,1,1,1,1,1,1]
# 解释： 1 在所有的子树中都未出现


# 思路： DFS
#
#       因为所有的的基因值都不一样，所以从根节点 0 到基因值为 1 的那个节点需要进行计算处理，
#       其他节点对应的答案都为 1 。
#
#       那我们可以找到基因值为 1 的那个节点，设为 cur ，则收集节点 cur 的所有子节点的基因值，
#       这样我们就知道节点 cur 对应的答案是多少。
#       接着我们再用相同方法处理 cur 的父节点计算答案，直到处理完根节点。
#       【注意】已经处理过的子树不应该再处理，而只有我们上一次 cur 的节点是被处理过的，
#           所以遇到上一个节点 pre 则可直接跳过。
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


class Solution:
    def smallestMissingValueSubtree(self, parents: List[int], nums: List[int]) -> List[int]:
        # 先初始化邻接表
        children = defaultdict(list)
        for i in range(1, len(parents)):
            children[parents[i]].append(i)
        
        # 初始化答案
        ans = [1] * len(parents)
        # 找到基因值为 1 的节点
        cur = -1
        for i, num in enumerate(nums):
            if num == 1:
                cur = i
                break
        # 如果没有基因值为 1 的节点，那么所有字数的答案都是 1
        if cur == -1:
            return ans

        # 收集从 cur 到根节点上所有子节点的基因值
        subtree_nums = set()
        def dfs(u):
            # 将当前节点的基因值加入集合
            subtree_nums.add(nums[u])
            # 递归处理子节点
            for v in children[u]:
                dfs(v)

        # 上一次处理过的节点，防止重复收集
        pre = -1
        # 上一次子树中确实的基因值（这个随着往上遍历是非减小的）
        miss = 1
        while cur != -1:
            # 收集所有子树的基因值
            subtree_nums.add(nums[cur])
            for child in children[cur]:
                if child != pre:
                    dfs(child)

            # 计算答案，如果当前基因值存在，则寻找下一个基因值
            while miss in subtree_nums:
                miss += 1
            ans[cur] = miss

            # 当前节点已经处理过，下一次不需要在处理 cur 子树
            pre = cur
            # 继续处理父节点
            cur = parents[cur]
        return ans
