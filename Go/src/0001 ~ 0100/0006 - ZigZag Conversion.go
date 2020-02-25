// 链接：https://leetcode.com/problems/zigzag-conversion/
// 题意：给定一个字符串和行数，按照 zigzag 模式写下来，然后按行拼成字符串（不包含空白字符）返回
// 输入：PAYPALISHIRING 3
// zigzag 模式：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 输出：PAHNAPLSIIGYIR

// 额外收获：Go 中结构体自定义排序需要实现 Interface 接口

// 思路：模拟，按照题意记录每个字符的位置
// 		由于方向的模式固定，所以可以定义成增量数组，统一循环逻辑，只用处理边界情况换方向
// 		最后按行号升序，再列号升序排序即可，按顺序拼接字符串返回
// 时间复杂度：O(nlogn) 空间复杂度：O(n)

import (
	"bytes"
	"sort"
)

func convert(s string, numRows int) string {
	// 初始化行列增量，分别表示 竖直向下 和 斜向右上
	dr := []int {1, -1} // 行号增量
	dc := []int {0, 1} // 列号增量

	// nodes 列表，记录每个字符及其位置
	nodes := make([]Node, len(s))
	// 初始化方向是 竖直向下，然后到达边界后斜向右上，交替进行
	direction := 0
	// 初始化为值在 1, 1 （以 1 为基准好判断，不用再多进行加减法）
	r, c := 1, 1
	for i, ch := range s  {
		nodes[i] = Node{byte(ch), r, c}
		r += dr[direction]
		c += dc[direction]

		// 到达 上边界 和 下边界 时要切换方向
		if r == 1 || r == numRows {
			direction = 1 - direction // 由于只有两个方向，所以用 1 减去可以不用取模或者多余判断
		}
	}

	// 排序
	sort.Sort(Nodes(nodes))
	// 按顺序组装结果即可
	var result bytes.Buffer
	for _, node := range nodes {
		result.WriteByte(node.Val)
	}
	return result.String()
}

// 定义元素类型
type Node struct {
	Val byte
	Row int
	Col int
}

// 定义数组类型
type Nodes []Node

// 返回 nodes 的长度
func (nodes Nodes) Len() int {
	return len(nodes)
}

// 判断 nodes[i] 是否小于 nodes[j]
func (nodes Nodes) Less(i, j int) bool {
	// 如果行号相当，则比较列号
	if nodes[i].Row == nodes[j].Row {
		return nodes[i].Col < nodes[j].Col
	}
	// 如果行号相当，则比较行号即可
	return nodes[i].Row < nodes[j].Row
}

// 交换两个元素
func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}