// 链接：https://leetcode.com/problems/largest-3-same-digit-number-in-string/
// 题意：给定一个数字字符串 num ，求满足以下条件的最大字符串，不存在则返回空串。
//          1. num 中长度为 3 的子串
//          2. 子串中的每个数字都是相同的


// 数据限制：
//  3 <= num.length <= 1000
//  num 仅含有数字


// 输入： num = "6777133339"
// 输出： "777"
// 解释： 有两个满足条件的子串 "777" 和 "333" ，其中 "777" 是最大的。

// 输入： num = "2300019"
// 输出： "000"
// 解释： 只有一个满足条件的子串 "000" 。

// 输入： num = "42352338"
// 输出： ""
// 解释： 没有满足条件的子串。


// 思路： 模拟
//
//      按照题目模拟即可，用 ans 表示满足题意的最大子串中的数字，
//      初始化为 0 ，表示不存在。
//
//      然后枚举以 num[i] 结尾的长度为 3 的子串，
//      如果该子串的字符都相同，则满足条件，更新 ans 的最大值。
//
//      最后如果 ans 不为 0 ，则返回对应的子串；否则返回空串。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 num 中全部 O(n) 个字符
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量


func largestGoodInteger(num string) string {
    // ans 表示满足题意的最大子串中的数字，初始化不存在
    ans := byte(0)
    for i := len(num) - 1; i >= 2; i-- {
        // 如果以 num[i] 结尾的子串满足条件，则更新 ans 的最大数字
        if num[i] == num[i - 1] && num[i] == num[i - 2] {
            ans = max(ans, num[i])
        }
    }

    if ans != 0 {
        // 如果 ans 存在，则返回其对应的长度为 3 的字符串
        return string([]byte{ans, ans, ans})
    } else {
        // 如果 ans 不存在，则返回空串
        return ""
    }
}

func max(a, b byte) byte {
    if a > b {
        return a
    } else {
        return b
    }
}
