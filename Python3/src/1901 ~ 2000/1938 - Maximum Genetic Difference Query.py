# 链接：https://leetcode.com/problems/maximum-genetic-difference-query/
# 题意：给定一棵树，每个节点都有一个全数唯一的值，现在有一些查询，
#       每个查询 query[i] = [node_i, val_i] 要求从根节点到 node_i 节点的路径所有节点中，
#       节点值与 val_i 异或后结果最大的值。


# 数据限制：
#   2 <= parents.length <= 10 ^ 5
#   每个非根节点都满足 0 <= parents[i] <= parents.length - 1
#   parents[root] == -1
#   1 <= queries.length <= 3 * 10 ^ 4
#   0 <= node_i <= parents.length - 1
#   0 <= val_i <= 2 * 10 ^ 5

# 输入： parents = [-1,0,1,1], queries = [[0,2],[3,2],[2,5]]
# 输出： [2,3,7]
# 解释： 
#       [0,2]: 从根节点到节点 0 的路径中，节点 0 与 2 异或后的结果最大的值为 2
#       [3,2]: 从根节点到节点 3 的路径中，节点 1 与 2 异或后的结果最大的值为 3
#       [2,5]: 从根节点到节点 2 的路径中，节点 2 与 5 异或后的结果最大的值为 7

# 输入： parents = [3,7,-1,2,0,7,0,2], queries = [[4,6],[1,15],[0,5]]
# 输出： [6,14,7]
# 解释： 
#       [4,6]: 从根节点到节点 4 的路径中，节点 0 与 6 异或后的结果最大的值为 6
#       [1,15]: 从根节点到节点 1 的路径中，节点 1 与 15 异或后的结果最大的值为 14
#       [0,5]: 从根节点到节点 0 的路径中，节点 2 与 5 异或后的结果最大的值为 7


# 思路： Trie
#
#       好久没写字典树，看了题解才知道字典树居然可以用来求异或和的最大值，
#       以前似乎也出过类似的题，没接触过真得比较难想到
#
#       我们对路径上所有的树按位从高到低建立 01 字典树，
#       字典树构建好后，我们按位从高到低遍历 val 即可，
#       如果 val 当前的位是 0 ，那我们倾向于选 1 的子节点，没有则只能选 0 的子节点，
#       如果 val 当前的位是 1 ，那我们倾向于选 0 的子节点，没有则只能选 1 的子节点，
#
#       我们对查询按照节点分组，
#       然后在原来的树上进行 dfs ，然后将路径上节点的值按位建立字典树，
#       当遍历到某个节点后，发现该节点需要查询时，就通过字典树计算出最大值即可。
#
#
#       如果需要在线查询的化，那么就得使用可持久化字典树。
#       核心思想：每次增加/减少一个数时，我们不重新建立整棵字典树，
#       而是至多修改路径上的 k 个节点，复用其他原有的节点即可。
#
#
#       时间复杂度： O(k * (n + q)) ，其中 k 考虑的二进制位
#       空间复杂度： O(n + q + 2 ^ k) ，其中 k 考虑的二进制位



class TrieNode:
    def __init__(self):
        self.count = 0
        self.next = [None, None]


class Trie:
    def __init__(self):
        self.root = TrieNode()
    
    def add(self, num: int, count: int):
        cur = self.root
        # 按位从高到低插入到字典树中
        for i in range(18, -1, -1):
            # 获取当前位
            digit = (num >> i) & 1
            # 如果对应的节点还没有，则创建一个
            if cur.next[digit] is None:
                cur.next[digit] = TrieNode()
            
            # 移动到该节点
            cur = cur.next[digit]
            # 该节点出现次数加 count
            cur.count += count

    def get_max_xor(self, num: int):
        cur = self.root
        ans = 0
        # 按位从高到低遍历
        for i in range(18, -1, -1):
            # 获取当前位
            digit = (num >> i) & 1
            # 获取另一位
            other_digit = 1 ^ digit
            # 如果另一位存在，则可以直接选择 other_digit ， ans 的当前位置为 1
            if cur.next[other_digit] and cur.next[other_digit].count:
                ans |= (1 << i)
                cur = cur.next[other_digit]
            else:
                # 如果另一位不存在，则当前位只能选择 digit ， ans 的当前位置为 0
                cur = cur.next[digit]

        return ans

class Solution:
    def maxGeneticDifference(self, parents: List[int], queries: List[List[int]]) -> List[int]:
        # 构建邻接表
        root = None
        self.adj = {}
        for i, parent in enumerate(parents):
            if parent == -1:
                root = i
            else:
                self.adj.setdefault(parent, []).append(i)
        
        # queries 按照节点分组
        self.grouped_queries = {}
        for i, (node, val) in enumerate(queries):
            self.grouped_queries.setdefault(node, []).append((i, val))
        
        # 构建 Trie ，并查询每个节点上的最大值
        self.trie = Trie()
        self.ans = [0] * len(queries)
        self.dfs(root)
        return self.ans
    
    def dfs(self, cur: int):
        # 将当前节点的值加上
        self.trie.add(cur, 1)

        # 计算到当前节点的查询结果
        for i, val in self.grouped_queries.get(cur, []):
            self.ans[i] = self.trie.get_max_xor(val)

        # 处理子节点
        for child in self.adj.get(cur, []):
            self.dfs(child)
            
        # 将当前节点的值减去
        self.trie.add(cur, -1)
