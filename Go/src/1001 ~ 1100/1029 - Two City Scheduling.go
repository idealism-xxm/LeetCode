// 链接：https://leetcode.com/problems/two-city-scheduling/
// 题意：给定一个二维整型数组 costs ，长度为 2 * n ，
//      其中 costs[i] = [aCost_i, bCost_i] ，
//      表示第 i 个人去 a 城市花费为 aCost_i ，去 b 城市花费为 bCost_i 。
//
//      求 a, b 城市各去 n 人的最小花费和？


// 数据限制：
//  2 * n == costs.length
//  2 <= costs.length <= 100
//  1 <= aCost_i, bCost_i <= 1000


// 输入： costs = [[10,20],[30,200],[400,50],[30,20]]
// 输出： 110
// 解释： 第 1 个人去 a 城市，花费为 10
//       第 2 个人去 b 城市，花费为 30
//       第 3 个人去 a 城市，花费为 50
//       第 4 个人去 b 城市，花费为 20
//
//       最小花费和 = 10 + 30 + 50 + 20 = 110   

// 输入： costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
// 输出： 1859

// 输入： costs = [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]
// 输出： 3086


// 思路： 贪心 + 排序
//
//      假设让全部 2 * n 个人都去 b 城市，此时所有的花费为 total ，
//      并且这个值对于给定的 costs 数组是固定的。
//
//      此时需要选择 n 个人去 b 城市，假设选择的人为 i ，
//      那么花费将变为 total - costs[i][1] + costs[i][0] 。
//
//      我们需要让这个值尽可能小，
//      那么就是选择 costs[i][0] - costs[i][1] 尽可能小的值即可。
//
//      所以我们可以按照 costs[i][0] - costs[i][1] 升序排序，
//      让前 n 个人去 a 城市，后 n 个人去 b 城市，
//      这样最终的花费和就是最小的。
//
//      
//      时间复杂度：O(nlogn)
//          1. 需要对 costs 进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历 costs 中全部 O(n) 个元素
//      空间复杂度：O(n)
//          1. 需要存储 costs 排序后的全部 O(n) 个元素


func twoCitySchedCost(costs [][]int) int {
    // 按照 costs[i][0] - costs[i][1] 升序排序
    sort.Sort(Costs(costs))

    // 维护所有人的费用和
    ans := 0
    // 计算每个城市要去的人数
    n := len(costs) >> 1
    // 前 n 个人去 a 城市，后 n 个人去 b 城市
    for i := 0; i < n; i++ {
        // 加上去 a 城市的人的费用
        ans += costs[i][0]
        // 加上去 b 城市的人的费用
        ans += costs[i + n][1]
    }

    return ans
}

// 定义数组类型
type Costs [][]int

// 返回 costs 的长度
func (costs Costs) Len() int {
	return len(costs)
}

// 判断 costs[i] 是否小于 costs[j]
func (costs Costs) Less(i, j int) bool {
	return costs[i][0] - costs[i][1] < costs[j][0] - costs[j][1]
}

// 交换两个元素
func (costs Costs) Swap(i, j int) {
	costs[i], costs[j] = costs[j], costs[i]
}