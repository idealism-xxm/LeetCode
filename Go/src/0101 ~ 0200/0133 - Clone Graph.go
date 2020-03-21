// 链接：https://leetcode.com/problems/clone-graph/
// 题意：给定一个无向图的邻接表，将这个邻接表深拷贝一份 ？

// 输入： adjList = [[2,4],[1,3],[2,4],[1,3]]
// 输出： [[2,4],[1,3],[2,4],[1,3]]
// 解释： 图中有四个结点
//		第一个结点 (val = 1) 与第二个结点 (val = 2) 和第四个结点 (val = 4) 相连
//		第二个结点 (val = 2) 与第一个结点 (val = 1) 和第三个结点 (val = 3) 相连
//		第三个结点 (val = 3) 与第二个结点 (val = 2) 和第四个结点 (val = 4) 相连
//		第四个结点 (val = 4) 与第一个结点 (val = 1) 和第三个结点 (val = 3) 相连

// 输入： adjList = [[]]
// 输出： [[]]
// 解释： 图中有一个结点
//		第一个结点 (val = 1) 没有相连的结点

// 输入： adjList = []
// 输出： []
// 解释： 图中有无结点

// 输入： adjList = [[2],[1]]
// 输出： [[2],[1]]
// 解释： 图中有两个结点
//		第一个结点 (val = 1) 与第二个结点 (val = 2) 相连
//		第二个结点 (val = 2) 与第一个结点 (val = 1) 相连

// 思路： 递归
//
//		从给定的结点递归拷贝即可，
//			1. 如果某一个结点未拷贝过，则创建一个新的，然后进行引用
//			2. 如果某一个结点已拷贝过，则直接进行引用
//
//		这样的方法让我想起了 Spring 中初始化实例时解决循环引用
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 标记已经创建过的结点，方便重复使用
	created := make(map[int]*Node)
	return dfs(created, node)
}

func dfs(created map[int]*Node, node *Node) *Node {
	// 如果已经创建过，则直接返回
	if copiedNode, exists := created[node.Val]; exists {
		return copiedNode
	}

	// 没有创建过，则创建并保存
	copiedNode := &Node {
		Val: node.Val,
		Neighbors: nil,
	}
	created[node.Val] = copiedNode
	// 递归处理邻接的结点
	for _, neighbor := range node.Neighbors {
		copiedNode.Neighbors = append(copiedNode.Neighbors, dfs(created, neighbor))
	}
	return copiedNode
}
