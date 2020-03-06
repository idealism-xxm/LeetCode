// 链接：https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
// 题意：给定一个二叉树，结点的定义如下：
//      struct Node {
//        int val;
//        Node *left;
//        Node *right;
//        Node *next;
//      }
//      初始情况下每个结点的 next 都为空，
//      现在需要处理成每个结点的 next 为当前一层的下一个结点？
//      （使用 O(1) 的额外空间，本题递归栈不算作额外空间）

// 输入： root = [1,2,3,4,5,null,7]
// 输出： [1,#,2,3,#,4,5,7,#]
// 解释： 输出是按照每一层通过 next 进行的遍历，每一层通过 # 分隔

// 思路：循环
//
//		思路同 0117 一样，不过需要加一点变化适配当前题目
//
//      1. 第 1 层的 next 不需要处理即符合题意
//      2. 假设第 i 层的 next 已经全部处理完
//         设第 i 层对应的最左边的结点分别为 leftMost
//         每一层的初始状态如下： cur = leftMost ，并记录第 i + 1 层在链上第最后一个结点 last
//         则可以通过以下方式更新第 i + 1 层的 next （注意：还需要记录下一层最左边的结点）
//         (1) cur 的左子结点存在，则将其挂在链上，可以进行更新：
//              last.next = cur.left
//              last = cur.left
//         (1) cur 的右子结点存在，则将其挂在链上，可以进行更新：
//              last.next = cur.right
//              last = cur.right
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/
class Solution {
    public Node connect(Node root) {
        Node leftMost = root;
        // 若 leftMost 存在，则可以继续处理下一层的 next
        while (leftMost != null) {
            Node cur = leftMost;
            // 用一个空结点表示第 i + 1 层链上第最后一个结点，方便统一逻辑
            Node last = new Node();
            // 下一层最左边的结点
            Node nextLeftMost = null;

            // 若 cur 存在，则可更新 i + 1 层
            while (cur != null) {
                // 如果左子结点存在，则将其挂在链上
                if (cur.left != null) {
                    // 如果还未找到下一层最左边的结点，则 cur.left 就是
                    if (nextLeftMost == null) {
                        nextLeftMost = cur.left;
                    }
                    last.next = cur.left;
                    last = last.next;
                }
                // 如果左子结点存在，则将其挂在链上
                if (cur.right != null) {
                    // 如果还未找到下一层最左边的结点，则 cur.left 就是
                    if (nextLeftMost == null) {
                        nextLeftMost = cur.right;
                    }
                    last.next = cur.right;
                    last = last.next;
                }
                // 处理下一个结点
                cur = cur.next;
            }
            // 继续处理下一层
            leftMost = nextLeftMost;
        }
        return root;
    }
}
