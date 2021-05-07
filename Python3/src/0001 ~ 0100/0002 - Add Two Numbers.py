# 链接：https://leetcode.com/problems/add-two-numbers/
# 题意：给定两个不含前导零的数，从个位用链表表示它们，返回它们的和对应的链表？
# 思路：模拟加法操作即可，可以复用节点，最高位进位时新建一个节点。

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 复用 l1 的节点，所以答案最后就是 l1
        ans = l1
        # add dummy node
        l1 = ListNode(next=l1)
        l2 = ListNode(next=l2)
        # 进位数
        carry = 0
        # 如果都还有下一个节点，则执行加法
        while l1.next and l2.next:
            l1.next.val, carry = self.add(l1.next.val, l2.next.val, carry)

            # 处理下一个节点
            l1 = l1.next
            l2 = l2.next

        # 如果 l1 没有下一个节点，则将 l2 剩余的节点挂在后面
        if not l1.next:
            l1.next = l2.next
        # 处理 l1 剩余的节点
        while l1.next:
            l1.next.val, carry = self.add(l1.next.val, 0, carry)

            # 处理下一个节点
            l1 = l1.next

        # 此时还有进位，则需要新建节点
        if carry:
            l1.next = ListNode(val=carry)

        return ans

    def add(self, a: int, b: int, carry: int) -> Tuple[int, int]:
        a += b + carry
        # 如果变成两位数，则需要进位
        if a >= 10:
            a -= 10
            carry = 1
        else:
            carry = 0

        return a, carry
