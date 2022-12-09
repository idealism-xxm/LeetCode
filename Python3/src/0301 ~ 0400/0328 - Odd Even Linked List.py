# 链接：https://leetcode.com/problems/odd-even-linked-list/
# 题意：给定一个单链表 head ，对结点按照下标奇偶性分组，组内保持原有顺序。
#      再将偶数下标的结点放在奇数下标的结点之后，返回结果链表。
#
#      第一个结点的下标是奇数，第二个结点的下标是偶数。
#
#      进阶：使用时间复杂度为 O(n) 且空间复杂度为 O(1) 算法。


# 数据限制：
#  链表的结点数范围为 [0, 10 ^ 4]
#  -(10 ^ 6) <= Node.val <= 10 ^ 6


# 输入： head = [1,2,3,4,5]
# 输出： [1,3,5,2,4]

# 输入： head = [2,1,3,5,6,4,7]
# 输出： [2,3,6,7,1,5,4]


# 思路： 一次迭代
#
#      我们可以使用奇偶链表 odd 和 even ，分别收集下标为奇数和偶数的结点，
#      最后将链表 even 挂在 odd 后面即可。
#
#      对于链表的题目，一般都可以使用一个哨兵结点。
#      本题使用哨兵结点，方便处理方便处理奇偶链表为空的情况。
#
#
#      时间复杂度：O(n)
#          1. 需要遍历全部 O(n) 个结点
#      空间复杂度：O(1)
#          1. 只需要维护常数个额外变量即可


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # 初始化奇偶链表的哨兵结点，方便处理链表为空的情况
        odd_dummy: ListNode = ListNode()
        even_dummy: ListNode = ListNode()
        # 初始化每个链表的尾部结点，方便使用尾插法插入结点
        odd_tail, even_tail = odd_dummy, even_dummy
        # is_odd 表示当前结点的下标是否为奇数
        is_odd: bool = True
        # 当还有结点时，继续分组到对应的链表
        while head:
            if is_odd:
                # 如果下标为奇数，则将 head 插入到奇链表尾部
                odd_tail.next = head
                odd_tail = odd_tail.next
            else:
                # 如果下标为偶数，则将 head 插入到偶链表尾部
                even_tail.next = head
                even_tail = even_tail.next
            
            # head 移动到下一个结点
            head = head.next
            # 下一个结点下标的奇偶性和当前相反
            is_odd = not is_odd

        # 最后将偶链表挂在奇链表后
        odd_tail.next = even_dummy.next
        # 防止最后一个结点下标为奇数时形成环
        even_tail.next = None

        return odd_dummy.next
