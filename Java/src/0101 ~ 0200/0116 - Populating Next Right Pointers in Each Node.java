// 链接：https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// 题意：给定一个完全二叉树，结点的定义如下：
//      struct Node {
//        int val;
//        Node *left;
//        Node *right;
//        Node *next;
//      }
//      初始情况下每个结点的 next 都为空，
//      现在需要处理成每个结点的 next 为当前一层的下一个结点？
//      （使用 O(1) 的额外空间，本题递归栈不算作额外空间）

// 输入： root = [1,2,3,4,5,6,7]
// 输出： [1,#,2,3,#,4,5,6,7,#]
// 解释： 输出是按照每一层通过 next 进行的遍历，每一层通过 # 分隔

// 思路1：递归
//
//		做这道题的时候体会到思路逐渐清晰的流程，每一次的想法都为后续的想法提供了灵感
//      想法1：很容易就能想到，递推处理的时候，
//            每次都可以对每个非叶子结点进行如下处理：root.left.next = root.right
//
//      想法2：接下来就开始想不是同一个结点的直接子结点的情况，然后就想先处理叶子结点，
//            既然递归还有返回值没有利用，那么可以返回一个子树中最左的叶子结点 和 最右的叶子结点，
//            这样可以自底向上更新所有的叶子结点，递归完后，所有的叶子结点就都连起来了。
//            同样，我们如果加一个入参，每次限定返回的深度，就可以完成所有层的更新了。
//            这样：每次遍历的结点个数为 n, n/2, n/4, ..., 1 ，总共遍历 logn 次
//            等比数列求和可得总遍历的结点次数为： n * (1 - (1/2)^n)/(1 - 1/2)
//            当 n -> ∞ 时可得： 2n
//
//            时间复杂度： O(n)
//
//      想法3：我们还可以充分利用完全二叉树的特性，只用遍历一遍即可自顶向下地更新所有结点
//            update(root, right) 表示更新 root 结点，且 root 所在层的右边的结点为 right
//            若 root 已是最右边的结点，则 right 必定等于 root ，此时不需要更新
//            1. 处理左子结点 root.left ，则对应的当层右边结点为 root.right
//            2. 处理右子结点 root.right ，则对应的当层右边结点为：
//              (1) root 是最右边的结点： root.right
//              (2) root 不是最右边的结点： right.left
//
//            时间复杂度： O(n)
//

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
        // 根结点是最右结点，next 等于自身
        this.update(root, root);
        return root;
    }

    private void update(Node root, Node right) {
        // 空结点直接返回
        if (root == null) {
            return ;
        }
        // root 非最右结点，则更新其右边的结点
        if (root != right) {
            root.next = right;
        }
        // 递归处理左子结点
        this.update(root.left, root.right);
        // 递归处理右子结点
        if (root == right) {
            // 最右边的结点的右子结点仍是最右边的结点
            this.update(root.right, root.right);
        } else {
            this.update(root.right, right.left);
        }
    }
}

// 思路2：循环
//
//		看了题解后才发现还存在真正 O(1) 额外空间的循环解法，并支持任意二叉树
//      主要是没考虑到还可以利用已处理的 next 进行辅助
//
//      1. 第 1 层的 next 不需要处理即符合题意
//      2. 假设第 i 层的 next 已经全部处理完
//         设第 i 层对应的最左边的结点分别为 leftMost
//         每一层的初始状态如下： cur = leftMost
//         则可以通过以下方式更新第 i + 1 层的 next
//         (1) cur 的左右子结点互为先后关系，可以直接更新：
//              cur.left.next = cur.right
//         (2) 若 cur.next 存在，则其右子结点的 next 指向其 next 的左子结点，则可以更新：
//              cur.right.next = cur.next.left
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

class Solution {
    public Node connect(Node root) {
        if (root == null) {
            return null;
        }

        Node leftMost = root;
        // 若 leftMost 非叶子结点，则可以继续处理下一层的 next
        while (leftMost.left != null) {
            Node cur = leftMost;
            // cur 的左右子结点互为先后关系，可以直接更新
            cur.left.next = cur.right;

            // 若 cur.next 存在，则其右子结点的 next 指向其 next 的左子结点
            while (cur.next != null) {
                cur.right.next = cur.next.left;

                // 继续处理下一个结点
                cur = cur.next;
                // cur 的左右子结点互为先后关系，可以直接更新
                cur.left.next = cur.right;
            }
            // 继续处理下一层
            leftMost = leftMost.left;
        }
        return root;
    }
}
