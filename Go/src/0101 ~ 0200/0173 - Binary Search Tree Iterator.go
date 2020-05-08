// 链接：https://leetcode.com/problems/binary-search-tree-iterator/
// 题意：实现一个二叉搜索树迭代器，能够通过一个根结点初始化迭代器，
//		实现 next() 和 hasNext() 方法， next() 返回下一个最小的数，
//		要求平均时间复杂度为 O(1) ，空间复杂度为 O(h) ，其中 h 为树高？

//     7
//   /   \
//  3     15
//      /    \
//     9     20
//
// BSTIterator iterator = new BSTIterator(root);
// iterator.next();    // return 3
// iterator.next();    // return 7
// iterator.hasNext(); // return true
// iterator.next();    // return 9
// iterator.hasNext(); // return true
// iterator.next();    // return 15
// iterator.hasNext(); // return true
// iterator.next();    // return 20
// iterator.hasNext(); // return false

// 思路： 栈
//
//		刚开始没看到要求的是平均时间复杂度，
//		想了一会儿还是只能想到用栈去模拟，并且只能是平均时间复杂度为 O(1) ，
//		因为第一次一定要找到最小的结点，这个就注定了不能是完全 O(1) 的，
//		由于遍历顺序为 左中右，
//		我们只需要保证遍历栈中的某个节点时，其左节点一定被遍历了即可，
//		那么我们入栈时，就可以不断将其左子节点入栈，直至为 nil
//
//		时间复杂度： O(1)
//		空间复杂度： O(h)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	// 存储结点栈（并且其左子结点都已经入栈过）
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{}
	bstIterator.putBst(root)
	return bstIterator
}

// 将以 root 为根的子二叉搜索树入栈
func (this *BSTIterator) putBst(root *TreeNode) {
	// 由于遍历顺序是 左中右 ，所以只要左节点部位空，就继续入栈，
	// 保证遍历到当前节点时，左节点已遍历过
	for ; root != nil; root = root.Left {
		this.stack = append(this.stack, root)
	}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	top := len(this.stack) - 1
	// 栈顶元素出栈
	topNode := this.stack[top]
	this.stack = this.stack[:top]
	// 将右子树入栈
	this.putBst(topNode.Right)
	return topNode.Val
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	// 只要栈不为空，就必定有下一个
	return len(this.stack) != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
