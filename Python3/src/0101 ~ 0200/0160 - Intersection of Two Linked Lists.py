# 链接：https://leetcode.com/problems/intersection-of-two-linked-lists/
# 题意：返回两个链表的相交结点，如果不存在，则返回 null 。
#
#      进阶：使用时间复杂度为 O(m + n) ，空间复杂度为 O(1) 方法求解。


# 数据限制：
#  listA 的结点数为 m
#  listB 的结点数为 n
#  1 <= m, n <= 3 * 10 ^ 4
#  1 <= Node.val <= 10 ^ 5
#  0 <= skipA < m
#  0 <= skipB < n
#  如果 listA 和 listB 不相交，则 intersectVal 是 0
#  如果 listA 和 listB 相交，则 intersectVal == listA[skipA] == listB[skipB] 


# 输入： intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
# 输出： Intersected at '8'
# 解释： A:       4 -> 1 ↘
#                         8 -> 4 -> 5
#       B:  5 -> 6 -> 1 ↗

# 输入： intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
# 输出： Intersected at '2'
# 解释： A:  1 -> 9 -> 1 ↘
#                         2 -> 4
#       B:            3 ↗

# 输入： intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
# 输出： No intersection
# 解释： A:  2 -> 6 -> 4
#
#       B:       1 -> 5


# 思路： 双指针
#
#      很容易就能想到进阶的解法：先求出两个链表的长度，计算它们的长度差 diff 。
#
#      然后让长的链表先走 diff 步，最后两个链表同时走，直到走到相同结点 node 。
#
#      此时必有两种情况：
#          1. node 不为空，则说明两个链表相交，且 node 就是相交结点
#          2. node 为空，则说明两个链表不相交
#
#      综上，直接返回 node 即可。
#
#
#      直接按照这样的思路实现的话，代码清晰但不够简洁，处理起来比较麻烦，
#      可以利用内在关系简化成一个循环。
#
#      我们先回想一下为什么要计算长的链表的长度？
#      因为要处理链表长度不相等的情况，好让链表尾部对齐。
#
#      那么只要想办法对齐链表尾部，就只需要无脑同时走，直至走到相同结点。
#
#      如果我们使用指针 cur_a 先遍历链表 a ，再遍历链表 b ；
#      然后用指针 cur_b 先遍历链表 b ，再遍历链表 a 。
#
#      那么 cur_a 和 cur_b 实际上就是将链表 b 和链表 a 进行了尾部对齐，
#      可以直接同时遍历。
#
#      我们分析一下不同的情况：
#          1. len(a) == len(b): 那么链表 a 和链表 b 本身就是对齐的，
#              在 cur_a 和 cur_b 第一次遍历 a 和 b 时，
#              就会走到相同结点（相交结点或空结点）。
#          2. a 和 b 不相交: 那么 cur_a 和 cur_b 在第二次遍历完 a 和 b 时，
#              就会走到空结点。
#          3. len(a) != len(b) && a 和 b 相交: 
#              那么 cur_a 和 cur_b 在均开始第二次遍历 a 和 b 时，
#              cur_a 和 cur_b 距离链表 b 和 a 尾部的长度相同，
#              相当于长的链表已经先走完了 diff 步，
#              后续同时走就会走到相同结点（相交结点或空结点）。
#
#
#      时间复杂度：O(m + n)
#          1. 需要遍历链表 a 全部 O(m) 个结点两次
#          2. 需要遍历链表 b 全部 O(n) 个结点两次
#      空间复杂度：O(1)
#          1. 只需要使用常数个额外变量


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        # 使用新的副本遍历，方便后续交换从头遍历
        cur_a, cur_b = headA, headB
        # 如果还未遇到相同的结点，则继续移动到下一个结点
        while cur_a != cur_b:
            # 如果 a/b 链表已遍历完成，则从头遍历另一个链表
            cur_a = cur_a.next if cur_a else headB
            cur_b = cur_b.next if cur_b else headA

        return cur_a
