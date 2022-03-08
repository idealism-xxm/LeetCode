# 链接：https://leetcode.com/problems/linked-list-cycle/
# 题意：给定一个链表，判断是否存在环？


# 数据限制：
#  两个链表中的结点数均在 [0, 50] 内
#  -100 <= Node.val <= 100
#  list1 和 list2 均为非降序链表


# 输入： head = [3,2,0,-4], pos = 1
# 输出： true
# 解释： 链表有环，第 3 个结点会连接到第 1 个结点 
#       3 → 2 → 0 → (-4)
#           ↑         ↓
#           └────←────┘       

# 输入： head = [1,2], pos = 0
# 输出： true
# 解释： 链表有环，第 1 个结点会连接到第 0 个结点 
#       1 → 2
#       ↑   ↓
#       └─←─┘    

# 输入： head = [1], pos = -1
# 输出： false
# 解释： 链表只有一个结点，无环


# 思路： 双指针
#
#		维护两个指针 fast 和 slow ，
#       每次循环时 fast 往前走两步， slow 往前走一步：
#		    1. 若 fast 遇到 nil ，则链表无环
#		    2. 若 fast 遇到 slow ，则链表有环
#
#
#		时间复杂度： O(n)
#           1. 需要遍历链表中的全部 O(n) 个结点
#		空间复杂度： O(1)
#           1. 只需要维护常数个额外变量


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        # 如果头结点为空 或者 头结点的下一个结点为空，
        # 则不存在环，直接返回 False
        if head is None or head.next is None:
            return False
        # 定义快指针 fast ，让其先走一步，方便后续判断
        fast = head.next
        # 定义慢指针 slow ，初始化为头结点
        slow = head
        # 如果快指针 fast 不等于 慢指针 slow ，
        # 则需要继续循环处理
        while fast != slow:
            # 如果快指针为空 或者 快指针的下一个结点为空，
            # 则不存在环，直接返回 False
            if fast is None or fast.next is None:
                return False

            # 快指针 fast 往前走两步
            fast = fast.next.next
            # 慢指针 slow 往前走一步
            slow = slow.next

        # 此时必有 fast == slow ，即存在环
        return True
