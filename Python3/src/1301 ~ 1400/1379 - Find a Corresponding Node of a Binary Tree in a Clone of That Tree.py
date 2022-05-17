# 链接：https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/
# 题意：给定两棵二叉树 original 和 cloned ，以及 original 中的一个结点的引用。
#       其中 cloned 是 original 的拷贝，反馈 cloned 中对应的结点的引用。
#
#       进阶：树中的结点值可重复时该如何处理？


# 数据限制：
#  树中的结点数在 [1, 10 ^ 4] 内
#  树中的结点值都是唯一的
#  target 是 original 中的一个结点，并且不为 null


# 输入： tree = [7,4,3,null,null,6,19], target = 3
# 输出： 3
# 解释： 7
#      / \
#     4  (3)
#        / \
#       6   19

# 输入： tree = [7], target =  7
# 输出： 7
# 解释： (7)

# 输入： tree = [8,null,6,null,5,null,4,null,3,null,2,null,1], target = 4
# 输出： 4


# 思路： DFS
#
#		由于 cloned 是 original 的拷贝，
#       所以按照相同的模式遍历两棵树即可。
#
#       当遍历到 original 的一个结点就是 target 时，
#       此时遍历到的 cloned 的对应结点就是所求结点。
#   
#
#       时间复杂度：O(n)
#           1. 需要遍历树的结点，直至找到对应的结点，
#               最差情况下，需要遍历全部 O(n) 个结点
#       空间复杂度：O(1)
#           1. 栈的深度是树的高度，最差情况下，所有结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        def dfs(original: TreeNode, cloned: TreeNode):
            # 如果当前子树为空，则直接返回空结点
            if not original:
                return None

            # 如果已在 original 中找到 target 结点，
            # 则此时 cloned 就是所求的结点
            if original == target:
                return cloned
        
            # 先在左子树中找 target ，如果找到就直接返回；
            # 否则 target 必定在右子树中，返回右子树寻找的结果
            return dfs(original.left, cloned.left) or dfs(original.right, cloned.right)
        
        # 由于 cloned 是 original 的拷贝，
        # 所以按照相同的模式遍历两棵树即可
        return dfs(original, cloned)
