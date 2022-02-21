#  链接：https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
#  题意：给定一棵二叉树，每个结点的值是 0 或 1 ，
#       现在将每一个根到叶的路径上的值形成的二进制数字相加，
#       返回最后相加的和的十进制值。

#  数据限制：
#   树的结点数量在 [1, 1000] 范围内
#   Node.val 是 0 或 1

#  输入： root = [1,0,1,0,1,0,1]
#  输出： 22
#  解释： (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22

#  输入： root = [0]
#  输出： 0

#  思路： DFS
# 
#       从根结点 dfs ，不断计算路径上 01 串代表的十进制数字 cur ，
#           1. 如果当前结点 node 是 none ，则返回 0 ，
#           2. 如果当前结点 node 不是 none ，则计算 val = (cur << 1) + node.val ，
#               (1) 如果 node 是叶子结点，则直接返回 val
#               (2) 如果 node 不是叶子结点，则返回 dfs(val, node.left) + dfs(val, node.right)
# 
#       数据范围决定了这是完全二叉树，所以 dfs 的递归深度为 O(logn) ，空间复杂度为 O(logn)
# 
#       时间复杂度： O(n)
#       空间复杂度： O(logn)


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumRootToLeaf(self, root: Optional[TreeNode]) -> int:
        def dfs(root: Optional[TreeNode], cur: int) -> int:
            # 当 root 是 None 时，直接返回 0
            # 这样就不用在上一层递归时判断左右子结点是否存在了
            if root is None:
                return 0

            # 更新路径上的值
            cur = (cur << 1) | root.val
            # 如果当前结点为叶子结点，则 cur 就是根到叶子结点的值
            if root.left is None and root.right is None:
                return cur

            # 如果当前结点不是叶子结点，则返回左右子结点的值的和
            return dfs(root.left, cur) + dfs(root.right, cur)

        return dfs(root, 0)
