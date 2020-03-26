// 链接：https://leetcode.com/problems/copy-list-with-random-pointer/
// 题意：给一个单链表，每一个结点除了有 Next 指向下一个结点，
//		还有一个 Random 指向链表中随机的一个结点，深拷贝这个链表？
//
// 输入： head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出： [[7,null],[13,0],[11,4],[10,2],[1,0]]

// 输入： head = [[1,1],[2,1]]
// 输出： [[1,1],[2,1]]

// 输入： head = []
// 输出： []

// 思路： 递归
//
//		思路和 0133 一致，先递归处理 Next 对应的结点，
//		然后处理 Random 对应的结点（此时对应对应的结点已经被深拷贝过）
//
//		时间复杂度： O(n)
//		空间复杂度： O(1)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	return dfs(make(map[*Node]*Node), head)
}

func dfs(nodeMap map[*Node]*Node, head *Node) *Node {
	if head == nil {
		return nil
	}
	headCopy := &Node{
		Val: head.Val,
		Next: nil,
		Random: nil,
	}
	// 先记录本结点的拷贝结点，方便后续 Random 直接拿出来用
	nodeMap[head] = headCopy
	// 递归处理下一个结点
	headCopy.Next = dfs(nodeMap, head.Next)
	// 直接从 map 中获取对应的拷贝结点
	headCopy.Random = nodeMap[head.Random]

	return headCopy
}
