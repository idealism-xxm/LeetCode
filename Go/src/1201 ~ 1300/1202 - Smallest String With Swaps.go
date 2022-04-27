// 链接：https://leetcode.com/problems/smallest-string-with-swaps/
// 题意：给定一个字符串 s 和一个下标二元组列表 pairs ，
//      其中 pairs[i] = [a, b] 。
//      你能任意次交换 s 中任意一对下标二元组对应的字符，
//      求能获得的字典序最小的字符串？


// 数据限制：
//  1 <= s.length <= 10 ^ 5
//  0 <= pairs.length <= 10 ^ 5
//  0 <= pairs[i][0], pairs[i][1] < s.length
//  s 仅含有英文小写字母


// 输入： s = "dcab", pairs = [[0,3],[1,2]]
// 输出： "bacd"
// 解释： 交换 s[0] 和 s[3], s = "bcad"
//       交换 s[1] 和 s[2], s = "bacd"

// 输入： s = "dcab", pairs = [[0,3],[1,2],[0,2]]
// 输出： "abcd"
// 解释： 交换 s[0] 和 s[3], s = "bcad"
//       交换 s[0] 和 s[2], s = "acbd"
//       交换 s[1] 和 s[2], s = "abcd"

// 输入： s = "cba", pairs = [[0,1],[1,2]]
// 输出： "abc"
// 解释： 交换 s[0] 和 s[1], s = "bca"
//       交换 s[1] 和 s[2], s = "bac"
//       交换 s[0] 和 s[1], s = "abc"


// 思路： 并查集
//
//      我们可以发现如果 (x, y) 和 (y, z) 位置的字符均可以交换，
//      那么 (x, z) 位置的字符也可以交换。
//
//      所以只要把所有能交换的位置通过并查集合并，
//      那么最终在同一个集合中的所有位置都能互相交换。
//
//      可以将同一个集合中的所有字符拿出来按升序排序，
//      再按顺序将排序后的字符放入结果字符串对应的位置中。
//
//      遍历处理完所有的集合后，形成的结果字符串就是字典序最小的。
//
//
//      设 n 为字符串的长度， m 为 pairs 的长度。
//
//      时间复杂度：O((n + m) * α(n) + nlogn)
//          1. 并查集每一次操作的时间复杂度都是 O(α(n))
//          2. 最初合并时，需要遍历 pairs 的全部 O(m) 个元组对，
//              每次都要执行常数次并查集操作
//          3. 后续获取每个集合的元素时，需要遍历全部 O(n) 个元素，
//              每次都要执行常数次并查集操作
//          4. 需要对一个集合中的字符排序，最差情况下有 O(n) 个字符，
//              所以排序时间复杂度为 O(nlogn)
//      空间复杂度：O(n)
//          1. 需要存储全部 O(n) 个元素的父元素和深度（秩）
//          2. 需要维护全部集合的元素列表，存储全部的 O(n) 个元素
//          3. 需要存储每个集合的字符，最差情况下有 O(n) 个字符
//          4. 需要存储结果字符串的全部 O(n) 个字符


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
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] == x {
        // 如果 x 的父元素是自己，那么 x 是根元素
        return x
    }

    // 如果 x 的父元素不是自己，那么递归查找其所在集合的根元素。
    // 这里使用路径压缩优化，将路径上所有的元素都直接挂在根元素下
    uf.parent[x] = uf.Find(uf.parent[x])
    // 返回 x 所在集合的根元素
    return uf.parent[x]
}

// 合并元素 x 和 y 所在的集合
func (uf *UnionFind) union(x int, y int) {
    // 找到 x 和 y 所在集合的根元素
    xRoot := uf.Find(x);
    yRoot := uf.Find(y);
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


func smallestStringWithSwaps(s string, pairs [][]int) string {
    // 初始化一个大小为 n 的并查集
    unionFind := newUnionFind(len(s))
    // 遍历每一对可互换的位置
    for _, pair := range pairs {
        // 将这两个位置对应的集合合并
        unionFind.union(pair[0], pair[1])
    }

    // 定义每个集合中的元素列表，
    //  key 为集合的根元素，
    //  value 为集合中的元素列表
    rootToIndices := make(map[int][]int)
    // 遍历每个元素
    for i := 0; i < len(s); i++ {
        // 找到元素所在集合的根元素
        root := unionFind.Find(i)
        // 将元素 i 添加到元素列表中
        rootToIndices[root] = append(rootToIndices[root], i)
    }

    // ans[i] 表示结果字符串中第 i 个位置的字符
    ans := make([]byte, len(s))
    // 遍历每个集合
    for _, indices := range rootToIndices {
        // 从原始字符串中收集当前元素列表对应的字符
        chars := make([]int, len(indices))
        for i, index := range indices {
            chars[i] = int(s[index])
        }
        // 将这些字符按升序排序
        sort.Ints(chars)
        // 按顺序将排序后的字符放入结果字符串中
        for i, index := range indices {
            ans[index] = byte(chars[i])
        }
    }

    // 转换成字符串并返回
    return string(ans)
}
