// 链接：https://leetcode.com/problems/binary-search-tree-iterator/
// 题意：给定一棵二叉搜索树 root ，实现一个迭代器，
//      支持中序遍历这棵二叉搜索树。
//
//      要求： next 和 hasNext 操作的平均时间复杂度是 O(1) ，
//          空间复杂度是 O(h) ，其中 h 是树的高度。


// 数据限制：
//  这棵树的结点数在 [0, 10 ^ 5] 内
//  0 <= Node.val <= 10 ^ 6
//  hasNext 和 next 最多会调用 10 ^ 5 次


// 输入： ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
//       [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// 输出： [null, 3, 7, true, 9, true, 15, true, 20, false]
// 解释： 7
//      / \
//     3  15
//        / \
//       9  20
//
//       BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
//       bSTIterator.next();    // 返回 3
//       bSTIterator.next();    // 返回 7
//       bSTIterator.hasNext(); // 返回 True
//       bSTIterator.next();    // 返回 9
//       bSTIterator.hasNext(); // 返回 True
//       bSTIterator.next();    // 返回 15
//       bSTIterator.hasNext(); // 返回 True
//       bSTIterator.next();    // 返回 20
//       bSTIterator.hasNext(); // 返回 False


// 思路： 栈
//
//      用栈模拟递归的中序遍历即可，维护一个栈，
//      保证遍历栈的全部结点的左子结点都一定会被先遍历。
//
//      那么在入栈一个子树 root 的根结点后，可以不断将其左子节点入栈，
//      直至左子结点为空，此时栈顶的结点就是中序遍历时的下一个结点。
//
//      当调用 next 时，返回栈顶结点 top 的值，
//      并将 top.left 这棵子树按照刚刚的方法入栈即可。
//
//      当调用 hasNext 时，只要栈不为空，就一定有下一个结点。
//
//
//		时间复杂度：平均 O(1)
//          1. next 最多会被调用 O(n) 次，且最多只有 O(n) 个结点会被遍历，
//              所以平均时间复杂度为 O(1)
//          2. hastNext 每次直接判断数组长度即可，时间复杂度为 O(1)
//		空间复杂度： O(h)
//          1. 栈递归深度就是树高 h ，最差情况下，全部 O(n) 个结点在一条链上


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	// 结点栈（所有结点的左子结点都已经入栈过）
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	// 初始化一个空栈的迭代器
	bstIterator := BSTIterator{}
	// 将 root 放入迭代器中
	bstIterator.putBst(root)
	return bstIterator
}


func (this *BSTIterator) Next() int {
	// 取出栈顶结点 top ， top 就是下一个需要遍历的结点
	top:= this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]
	// 将右子树入栈
	this.putBst(top.Right)
	// 返回 top 的值
	return top.Val
}


func (this *BSTIterator) HasNext() bool {
	// 如果栈不为空，则还有下一个结点
	return len(this.stack) != 0
}


func (this *BSTIterator) putBst(root *TreeNode) {
	// 当 root 不为空时，继续处理
	for ; root != nil; {
		// 将根结点 root 入栈
		this.stack = append(this.stack, root)
		// 更新 root 为其左子树
		root = root.Left
	}
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
