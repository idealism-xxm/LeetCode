# 链接：https://leetcode.com/problems/convert-bst-to-greater-tree/
# 题意：给定一棵二叉搜索树 root ，
#      将这个树的每个结点的值加上所有大于它的结点的值之和，
#      最后再返回结果树的根节点。


# 数据限制：
#  这棵树的结点数在  [0, 10 ^ 4] 内
#  -(10 ^ 4) <= Node.val <= 10 ^ 4
#  这棵树所有结点的值均不相同
#  root 是一棵合法的二叉搜索树


# 输入： root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
# 输出： [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
# 解释：       4                          30
#          /     \                    /      \
#         1       6                 36       21
#        / \     / \       →        / \      / \   
#       0   2   5   7             36  35    26  15
#          /         \                /          \
#         3           8              33           8

# 输入： root = [0,null,1]
# 输出： [1,null,1]
# 解释： 0            1
#        \     →      \
#         1            1


# 思路： 递归
#
#      二叉搜索树是左子树的值小于根节点，右子树的值大于根节点。
#
#      所以我们可以通过递归的方式来遍历这棵树，
#      并维护一个值 sum ，
#      表示大于等于当前结点值的所有结点值之和。
#
#      递归时，如果当前结点 root 为空，则直接返回。
#
#      如果当前结点 root 不为空，则先递归处理右子树，
#      因为右子树的结点值一定大于 root.val 。
#
#      然后再对 sum 加上当前结点值 root.val ，
#      这时 sum 变为所有大于等于 root.val 的结点值之和。
#
#      此时，按照题意 root.val 应该等于 sum ，
#      即 root.val = sum 。
#
#      最后再递归处理左子树即可。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历二叉搜索树中全部 O(n) 个结点
#      空间复杂度：O(n)
#          1. 栈递归深度就是树高，最差情况下，全部 O(n) 个结点在一条链上


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def convertBST(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        # sum 维护大于等于当前结点值的所有结点值之和
        sm: int = 0
        def convert(root: Optional[TreeNode]):
            # 如果 root 为空，则直接返回
            if not root:
                return

            # 此时 root 不为空，则继续处理。
            # 先递归处理右子树，因为右子树的结点值一定大于 root.val
            convert(root.right)
            # 再加上当前结点的值
            nonlocal sm
            sm += root.val
            # 此时 sum 是大于等于 root.val 的所有结点值之和，
            # 按照题意 root.val 应该等于 sum
            root.val = sm
            # 再递归处理左子树
            convert(root.left)

        convert(root)
        return root
