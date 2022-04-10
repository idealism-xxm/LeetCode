// 链接：https://leetcode.com/problems/baseball-game/
// 题意：给定一个操作列表 ops ，
//      在一个空的分数列表 record 中按顺序执行这些操作，
//      返回最终的分数列表 record 中所有分数之和。
//
//      操作 op 有以下 4 种：
//          1. 整数：将该整数放入 record
//          2. "+" ：将 record 中的最后两个整数相加，并将结果放入 record
//          3. "D" ：将 record 中的最后一个整数乘 2 ，并将结果放入 record
//          4. "C" ：将 record 中的最后一个整数移除  


// 数据限制：
//  1 <= ops.length <= 1000
//  ops[i] 是 "C", "D", "+", 范围在 [-(3 * 10 ^ 4), 3 * 10 ^ 4] 内的整数字符串
//  对于操作 "+" ， 分数列表中至少存在两个整数
//  对于操作 "C" 和 "D" ，分数列表中至少存在一个整数


// 输入： ops = ["5","2","C","D","+"]
// 输出： 30
// 解释：
//      "5" -> 将 5 放入 record 中， record 现在是 [5]
//      "2" -> 将 2 放入 record 中， record 现在是 [5, 2]
//      "C" -> 将 record 中的最后一个数移除，现在 record 现在是 [5]
//      "D" -> 将 record 中的最后一个数乘以 2，
//              并将结果放入 record，现在 record 现在是 [5, 10]
//      "+" -> 将 record 中的最后两个数相加，并将结果放入 record，
//              现在 record 现在是 [5, 10, 15]
//      最后， record 中所有数的和为 5 + 10 + 15 = 30

// 输入： ops = ["5","-2","4","C","D","9","+","+"]
// 输出： 27
// 解释：
//      "5"  -> 将 5  放入 record 中， record 现在是 [5]
//      "-2" -> 将 -2 放入 record 中， record 现在是 [5, -2]
//      "4"  -> 将 4  放入 record 中， record 现在是 [5, -2, 4]
//      "C"  -> 将 record 中的最后一个数移除，现在 record 现在是 [5, -2]
//      "D"  -> 将 record 中的最后一个数乘以 2，
//              并将结果放入 record，现在 record 现在是 [5, -2, -4]
//      "9"  -> 将 9  放入 record 中， record 现在是 [5, -2, -4, 9]
//      "+"  -> 将 record 中的最后两个数相加，并将结果放入 record，
//              现在 record 现在是 [5, -2, -4, 9, 5]
//      "+"  -> 将 record 中的最后两个数相加，并将结果放入 record，
//              现在 record 现在是 [5, -2, -4, 9, 5, 14]
//      最后， record 中所有数的和为 5 + (-2) + (-4) + 9 + 5 + 14 = 27

// 输入： ops = ["1"]
// 输出： 1


// 思路： 模拟
//
//      初始化一个空的分数列表 record ，
//      然后按顺序执行 ops 中的操作处理即可，
//      最后直接返回 record 中所有数的和。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 ops 中全部 O(n) 个操作
//          2. 需要对 record 中全部数字求和，最差情况下有 O(n) 个数字
//      空间复杂度：O(n)
//          1. 需要维护包含所有数字的分数列表 record ，
//              最差情况下有 O(n) 个数字


func calPoints(ops []string) int {
    // 初始化空的分数列表 record
    record := make([]int, 0, len(ops))
    // 按顺序枚举每一个操作
    for _, op := range ops {
        // 根据操作的不同，更新 record
        switch op {
            case "+":
                // 如果是 "+" ，则将 record 中的最后两个数相加，并将结果放入 record
                record = append(record, record[len(record) - 1] + record[len(record) - 2])
            case "D":
                // 如果是 "D" ，则将 record 中的最后一个数乘以 2，并将结果放入 record
                record = append(record, record[len(record) - 1] * 2)
            case "C":
                // 如果是 "C" ，则移除 record 中的最后一个数
                record = record[:len(record) - 1]
            default:
                // 其他情况，则认为是整数，转换成整数后，放入 record
                num, _ := strconv.Atoi(op)
                record = append(record, num)
        }
    }

    // 对 reocrd 中的所有数求和并返回
    sum := 0
    for _, num := range record {
        sum += num
    }
    return sum
}
