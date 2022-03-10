# 链接：https://leetcode.com/problems/add-two-numbers/
# 题意：给定两个非空链表表示的不含前导零的非负整数（逆序存储），
#      求这两个整数的和，并以相同形式的链表返回。


# 数据限制：
#   两个链表中的结点数在 [1, 100] 内
#   0 <= Node.val <= 9
#   两个链表所表示的数不含前导零


# 输入： l1 = [2,4,3], l2 = [5,6,4]
# 输出： [7,0,8]
# 解释： 342 + 465 = 807
#
#       2 -> 4 -> 3 
#            +
#       5 -> 6 -> 4
#            ↓
#       7 -> 0 -> 8

# 输入： l1 = [0], l2 = [0]
# 输出： [0]
# 解释： 0 + 0 = 0
#
#       0 
#       +
#       0
#       ↓
#       0

# 输入： l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
# 输出： [8,9,9,9,0,0,0,1]
# 解释： 9999999 + 9999 = 89990001
#      9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9
#                +
#      9 -> 9 -> 9 -> 9
#                ↓
#      8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1


# 思路： 模拟
#
#		按照通常的加法器模拟即可，从个位开始按位加，
#       注意进位即可，特别是最高位进位时要添加新的结点。
#
#       类似题目： LeetCode 67 - 二进制求和
#
#
#		时间复杂度： O(|l1| + |l2|)
#           1. 需要遍历 l1 中的全部 O(|l1|) 个结点
#           2. 需要遍历 l2 中的全部 O(|l2|) 个结点
#		空间复杂度： O(max(|l1|, |l2|))
#           1. 需要为结果链表中的全部 O(max(|l1|, |l2|)) 个结点分配空间
#           2. （理论上可以复用已有的结点，这样就只需要定义常数个额外结点，
#               能将空间复杂度优化为 O(1) ）


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        # 哨兵结点，方便后续处理
        head_pre = ListNode(0)
        # 结果链表的尾结点，方便用尾插法插入
        tail = head_pre
        # 进位值，初始化为 0
        carry = 0
        # 如果两个链表至少有一个还有结点 或者 还有进位，则继续循环处理
        while l1 or l2 or carry > 0:
            # 当前位的和，初始化为前一位的近位值
            sm = carry
            # 如果 l1 还有结点
            if l1:
                # 当前位的和加上 l1 中该位的值
                sm += l1.val
                # l1 向后移动一个结点
                l1 = l1.next
            # 如果 l2 还有结点
            if l2:
                # 当前位的和加上 l2 中该位的值
                sm += l2.val
                # l2 向后移动一个结点
                l2 = l2.next

            # 计算当前位的进位值
            carry = sm // 10
            # 将当前位的值加入到结果链表中
            tail.next = ListNode(sm % 10)
            # 尾结点向后移动一个结点
            tail = tail.next

        # 返回结果链表的头结点
        return head_pre.next
