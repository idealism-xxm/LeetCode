// 链接：https://leetcode.com/problems/copy-list-with-random-pointer/
// 题意：给一个单链表，每一个结点除了有 next 指向下一个结点，
//		还有一个 random 指向链表中随机的一个结点，
//		现在深拷贝这个链表，并返回结果链表的头结点。

// 数据限制：
//	0 <= n <= 1000
//	-(10 ^ 4) <= Node.val <= 10 ^ 4
//	Node.random 为空或者链表中存在的某个结点


// 输入： head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出： [[7,null],[13,0],[11,4],[10,2],[1,0]]

// 输入： head = [[1,1],[2,1]]
// 输出： [[1,1],[2,1]]

// 输入： head = [[3,null],[3,0],[3,null]]
// 输出： [[3,null],[3,0],[3,null]]


// 思路1： 递归 + Map
//
//      相同的结点只会被拷贝一次，因此我们可以维护一个 Map 来记录，
//      originToCloned[origin] 表示结点 origin 对应的已经拷贝过的结点。
//
//		从给定的结点开始递归拷贝即可，
//			1. 如果当前结点 cur 已拷贝过，则直接进行引用并返回
//			2. 如果当前结点 cur 未拷贝过，则创建一个拷贝的结点 cloned ，
//          	将其放入到 originToCloned 中
//
//          	然后递归处理 cur.next 结点。
//
//				再直接从 originToCloned 中，
//				取出 cur.random 对应的拷贝结点，
//				放入到 cloned.random 中，
//				因为此时链表中的所有结点已经被深拷贝过。
//
//
//		关联题目： LeetCode 133 - 克隆图
//
//
//		时间复杂度： O(n)
//			1. 需要遍历链表中的全部 O(n) 个结点
//		空间复杂度： O(n)
//			1. 需要对拷贝链表中的全部 O(n) 个结点
//			2. 需要维护一个 O(n) 的 Map 来记录已经拷贝过的结点

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// 维护每个结点下标对应的拷贝结点
	originToCloned := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func (cur *Node) *Node {
		// 如果当前结点 cur 为空，则直接返回即可
		if cur == nil {
			return nil
		}

		// 如果当前结点 cur 已拷贝过，则直接取出返回
		if cloned, exists := originToCloned[cur]; exists {
			return cloned
		}

		// 此时需要创建当前结点 cur 的拷贝结点 cloned
		cloned := &Node{Val: cur.Val}
		// 将 cloned 放入到 originToCloned 中
		originToCloned[cur] = cloned
		// 递归处理 cur.next 结点
		cloned.Next = dfs(cur.Next)
		// 此时链表中所有的结点都已拷贝过，
		// 直接从 originToCloned 中取出即可
		if cur.Random != nil {
			// 取出 cur.random 对应的拷贝结点
			cloned.Random = originToCloned[cur.Random]
		}

		// 返回当前结点的拷贝结点 cloned
		return cloned
	}

	return dfs(head)
}
