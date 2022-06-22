// 链接：https://leetcode.com/problems/calculate-amount-paid-in-taxes/
// 题意：给定一个二维数组 brackets ，其中 brackets[i] = [upper_i, percent_i] ，
//      表示第 i 个税级的上限是 upper_i ，税率是 percent_i 。
//      brackets 是按照 upper 升序排序的，税的计算方式如下：
//          1. 不超过 upper_0 的部分按照税率 percent_0 缴纳
//          2. 接着 upper_1 - upper_0 的部分按照 percent_1 缴纳
//          3. 然后 upper_2 - upper_1 的部分按照 percent_2 缴纳
//          4. 以此类推
//
//      现在给定总收入 income ，计算应缴税额。


// 数据限制：
//  1 <= brackets.length <= 100
//  1 <= upper_i <= 1000
//  0 <= percent_i <= 100
//  0 <= income <= 1000
//  upper_i 是按照升序排序的
//  所有的 upper_i 都可不相同
//  最后一个税级的上限大于等于 income


// 输入： brackets = [[3,50],[7,10],[12,25]], income = 10
// 输出： 2.65000
// 解释： 3 * 50% + (7 - 4) * 10% + (10 - 7) * 25% = 2.65

// 输入： brackets = [[1,0],[4,25],[5,50]], income = 2
// 输出： 0.25000
// 解释： 1 * 0% + (2 - 1) * 25% = 0.25

// 输入： brackets = [[2,50]], income = 0
// 输出： 0.00000
// 解释： 0 * 50% = 0


// 思路： 模拟
//
//      按照题目模拟计算即可。
//
//      为了方便处理，使用 pre 记录前一个税级的上限，
//      初始化为 0 ，就能兼容处理第一个税级的情况。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历触及的全部税级，最差情况下要遍历 brackets 全部 O(n) 个税级
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量


func calculateTax(brackets [][]int, income int) float64 {
    // 应缴税额，初始为 0
    ans := 0.0
    // 前一个税级的上限，初始为 0
    pre := 0
    for _, bracket := range brackets {
        if income <= bracket[0] {
            // 如果没超过当前税级上限，则计入最后触及的税级中的税，然后返回
            ans += float64((income - pre) * bracket[1]) / 100
            return ans
        }
        // 计入当前税级中的税，并更新前一个税级的上限
        ans += float64((bracket[0] - pre) * bracket[1]) / 100
        pre = bracket[0]
    }

    // 由于在循环里提前返回了，所以不会到达这里
    return 0
}
