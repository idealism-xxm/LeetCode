// 链接：https://leetcode.com/problems/satisfiability-of-equality-equations/
// 题意：给定一个字符串数组 equations ， equations[i] 长度为 4 ，
//      形如 x_i==y_i 或 x_i!=y_i ，表示两个变量 x_i 和 y_i 的值是否相等。
//
//      判断所有的等式是否同时成立？


// 数据限制：
//  1 <= equations.length <= 500
//  equations[i].length == 4
//  equations[i][0] 是一个英文小写字母
//  equations[i][1] 是 '=' 或 '!'
//  equations[i][2] 是 '='
//  equations[i][3] 是一个英文小写字母


// 输入： equations = ["a==b","b!=a"]
// 输出： false
// 解释： 第 0 个等式要求 a 与 b 相等，
//       但第 1 个等式要求 a 与 b 不等，
//       相互矛盾，所以所有的等式无法同时成立。

// 输入： equations = ["b==a","a==b"]
// 输出： true
// 解释： 第 0 个等式和第 1 个等式都要求 a 与 b 相等，
//       所有的等式可以同时成立。


// 思路： 并查集
//
//      等式具有传递性，即 a == b, b == c ，那么必定有 a == c 。
//
//      所以对于等式，我们可以用一个并查集维护所有等式的相等关系，
//      即在同一个集合中的变量值是相等的。
//
//      此时，我们可以再遍历所有不等关系，进行如下判断处理。
//
//      如果 a != b ，那么要让所有等式都成立的话，
//      a == b 必定不能成立，
//      即 a 与 b 在并查集中的不同集合中。
//
//      否则，表面同时存在 a == b 和 a != b 两个等式，相互矛盾，直接返回 false
//
//
//      设字符集大小为 C 。
//      
//      时间复杂度：O(n * α(C))
//          1. 需要对全部 O(n) 个等式执行并查集操作，
//              并查集每一次操作的时间复杂度都是 O(α(C))
//      空间复杂度：O(C)
//          1. 只需要在 uf 中维护全部 O(C) 个不同的字母的关系


// 获取小写字母 ch 的下标
func getIndex(ch byte) int {
    return int(ch - 'a')
}


func equationsPossible(equations []string) bool {
    // 初始化并查集，在同一个集合中的变量值是相等的。
    // 只用维护全部 26 个不同字母的关系
    uf := newUnionFind(26)
    // 将所有等式的关系连起来放入同一个集合
    for _, equation := range equations {
        if equation[1] == '=' {
            uf.union(getIndex(equation[0]), getIndex(equation[3]))
        }
    }

    // 判断所有不等关系对应的两个变量是否在同一个集合中
    for _, equation := range equations {
        // 如果当前不等关系的两个变量在同一个集合中，则所有等式无法同时满足，直接返回 false
        if equation[1] == '!' && uf.find(getIndex(equation[0])) == uf.find(getIndex(equation[3])) {
            return false
        }
    }

    // 此时所有等式可以同时满足
    return true
}


// 并查集
type UnionFind struct {
    // parent[i] 表示第 i 个元素所指向的父元素
    parent []int
    // rank[i] 表示以第 i 个元素的深度（秩），
    // 当 i 是根元素（即 parent[i] == i ）时有效
    rank []int
}


    // 初始化一个大小为 n 的并查集
func newUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        // 初始每个元素的父元素都是自己
        parent[i] = i
        // 初始化深度（秩）都是 1
        rank[i] = 1
    }
    return &UnionFind{parent, rank}
}

// 查找元素 x 所在集合的根元素
func (uf *UnionFind) find(x int) int {
    if uf.parent[x] == x {
        // 如果 x 的父元素是自己，那么 x 是根元素
        return x
    }

    // 如果 x 的父元素不是自己，那么递归查找其所在集合的根元素。
    // 这里使用路径压缩优化，将路径上所有的元素都直接挂在根元素下
    uf.parent[x] = uf.find(uf.parent[x])
    // 返回 x 所在集合的根元素
    return uf.parent[x]
}

// 合并元素 x 和 y 所在的集合
func (uf *UnionFind) union(x int, y int) {
    // 找到 x 和 y 所在集合的根元素
    xRoot := uf.find(x);
    yRoot := uf.find(y);
    // 如果 x 和 y 在同一个集合，则不需要合并
    if xRoot == yRoot {
        return
    }

    if uf.rank[xRoot] < uf.rank[yRoot] {
        // 如果 xRoot 深度（秩）更小，
        // 则将 yRoot 合并入 xRoot 中
        uf.parent[xRoot] = yRoot
    } else if uf.rank[xRoot] > uf.rank[yRoot] {
        // 如果 xRoot 深度（秩）更大，
        // 则将 xRoot 合并入 yRoot 中
        uf.parent[yRoot] = xRoot
    } else {
        // 如果 xRoot 深度（秩）相等，
        // 则将 yRoot 合并入 xRoot 中
        uf.parent[yRoot] = xRoot
        // 同时将 xRoot 的深度（秩）加 1
        uf.rank[xRoot] += 1
    }
}
