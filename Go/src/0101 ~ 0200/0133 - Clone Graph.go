//  链接：https://leetcode.com/problems/clone-graph/
//  题意：给定一个无向图，将这个图深拷贝一份。


// 数据限制：
//   图中的结点数范围是 [0, 100]
//   1 <= Node.val <= 100
//   每个节点的值 Node.val 各不相同
//   图中没有重复边和自环
//   图是联通的，并且所有的结点都能从给定的结点被访问到


// 输入： adjList = [[2,4],[1,3],[2,4],[1,3]]
// 输出： [[2,4],[1,3],[2,4],[1,3]]
// 解释： 图中有四个结点：
//
//       1 —— 2
//       |    |
//       4 —— 3
//
//		第一个结点 (val = 1) 与第二个结点 (val = 2) 和第四个结点 (val = 4) 相连
//		第二个结点 (val = 2) 与第一个结点 (val = 1) 和第三个结点 (val = 3) 相连
//		第三个结点 (val = 3) 与第二个结点 (val = 2) 和第四个结点 (val = 4) 相连
//		第四个结点 (val = 4) 与第一个结点 (val = 1) 和第三个结点 (val = 3) 相连

// 输入： adjList = [[]]
// 输出： [[]]
// 解释： 图中有一个结点：
//			第一个结点 (val = 1) 没有相连的结点

// 输入： adjList = []
// 输出： []
// 解释： 图中有无结点


// 思路： 递归 + Map
//
//      相同的结点只会被拷贝一次，因此我们可以维护一个 Map 来记录，
//      valToClonedNode[val] 表示 val 对应的已经拷贝过的结点。
//
//		从给定的结点开始递归拷贝即可，
//			1. 如果当前结点 cur 未拷贝过，则创建一个拷贝的结点 cloned ，
//          	将其放入到 valToClonedNode 中
//			2. 如果当前结点 cur 已拷贝过，则直接进行引用
//
//          然后递归处理 cur 的所有相邻结点，最后返回 cloned 即可
//
//
//		关联题目： LeetCode 138 - 复制带随机指针的链表
//
//
//		时间复杂度： O(m)
//          1. 需要遍历全部 O(m)条边
//		空间复杂度： O(n)
//          1. 需要存储全部 O(n) 的结点的拷贝结点
//          2. 栈递归深度最大为 O(n)


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	// 如果是图是空的，则直接返回
	if node == nil {
		return nil
	}

	// 维护每个值 val 对应的拷贝结点
	valToClonedNode := make(map[int]*Node)

	var dfs func(*Node) *Node
	dfs = func (cur *Node) *Node {
		// 如果当前结点 cur 已拷贝过，则直接返回即可
		if cloned, exists := valToClonedNode[cur.Val]; exists {
			return cloned
		}

		// 此时需要创建当前结点 cur 的拷贝结点 cloned
		cloned := &Node{Val: cur.Val}
		// 将 cloned 放入到 valToClonedNode 中
		valToClonedNode[cur.Val] = cloned
		// 递归处理当前结点 cur 的所有相邻结点，
		// 并将拷贝后的结点放入到 cloned.neighbors 中
		for _, neighbor := range cur.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor))
		}

		// 返回当前结点的拷贝结点 cloned
		return cloned
	}

	return dfs(node)
}
