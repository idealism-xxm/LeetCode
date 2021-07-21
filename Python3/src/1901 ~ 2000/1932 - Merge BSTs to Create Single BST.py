# 链接：https://leetcode.com/problems/merge-bsts-to-create-single-bst/
# 题意：给定 n 棵 BST ，判断其是否可以合并为一棵 BST ？
#       合并方式：
#           选择两棵还存在的 BST ，例如 a 和 b ，
#           如果 a 的根节点的值是 b 的叶子节点 x 的值，
#           则可以将 x 替换为 a ，并移除 a

# 数据限制：
#   n == trees.length
#   1 <= n <= 5 * 10 ^ 4
#   每棵树的节点数在 [1, 3] 之间
#   每个节点可能会有子节点，但不会有孙节点
#   没有两棵树的根有相同值
#   所有的树都是合法的 BST
#   1 <= TreeNode.val <= 5 * 10 ^ 4

# 输入： trees = [[2,1],[3,2,5],[5,4]]
# 输出： [3,2,5,1,null,4]
# 解释：
#       [[2,1],[3,2,5],[5,4]] -> [[3,2,5,1],[5,4]] -> [3,2,5,1,null,4]
#       [3,2,5,1,null,4] 是一个合法的 BST

# 输入： trees = [[5,3,8],[3,2,6]]
# 输出： []
# 解释：
#       [[5,3,8],[3,2,6]] -> [[5,3,8,2,6]]
#       [[5,3,8,2,6]] 不是一个合法的 BST

# 输入： trees = [[5,4],[3]]
# 输出： []

# 输入： trees = [[2,1,3]]
# 输出： [[2,1,3]]
# 解释： [[2,1,3]] 是一个合法的 BST


# 思路： 模拟
#
#       按照题意模拟合并即可，我们可以在合并的过程中就进行判断，
#       保证每次合并后的树都是合法的 BST
#
#       最后找到合并后的树的根节点，然后将其返回即可
#       如果存在多个树，则不满足题意，返回 None
#
#       时间复杂度： O(n)
#       空间复杂度： O(n)


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def canMerge(self, trees: List[TreeNode]) -> TreeNode:
        # 计算所有节点的入度，并保存根节点的对应关系
        val_to_indeg = defaultdict(int)
        self.mp = {}
        for tree in trees:
            self.mp[tree.val] = tree
            if tree.val not in val_to_indeg:
                val_to_indeg[tree.val] = 0
            if tree.left:
                val_to_indeg[tree.left.val] += 1
            if tree.right:
                val_to_indeg[tree.right.val] += 1

        # 找到入度为 0 的节点列表
        root_vals = [val for val, indeg in val_to_indeg.items() if indeg == 0]
        # 如果入度为 0 的节点数不是 1 ，则不可能合并
        if len(root_vals) != 1:
            return None

        # 如果根节点无法合并，则返回 None
        if not self.merge(self.mp[root_vals[0]], -1, 100000):
            return None

        # 如果最终的树不只一个，则不满足题意
        roots = set(self.mp.values()) - {None}
        if len(roots) > 1:
            return None

        # 返回唯一的这棵树
        return roots.pop()

    def merge(self, root: TreeNode, mn: int, mx: int):
        # 如果是空节点，则必定满足题意
        if root is None:
            return True

        # 如果当前节点不在范围内，则不是合法的 BST
        if not mn < root.val < mx:
            return False

        # 如果当前节点是叶子节点，则可以合并
        if not root.left and not root.right:
            # 如果当前值对应的树存在，且不是当前节点，则能合并
            tree = self.mp.get(root.val)
            if tree and root != tree:
                # 标记该树已被合并
                self.mp[tree.val] = None
                # 将当前节点的子节点替换为该树的子节点
                root.left = tree.left
                root.right = tree.right

            # 如果此时还是叶子节点，则直接返回 True
            if not root.left and not root.right:
                return True

        # 左右子树都能成功合并，则当前能成功合并
        return self.merge(root.left, mn, root.val) and self.merge(root.right, root.val, mx)
