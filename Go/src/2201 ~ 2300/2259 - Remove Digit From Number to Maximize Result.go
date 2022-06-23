// 链接：https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/
// 题意：给定一个数字字符串 number 和一个数字字符 digit ，
//      求删除 number 中的一个 digit 数字字符后，能形成的最大的数字字符串？


// 数据限制：
//  2 <= number.length <= 100
//  number 仅由 '1' - '9' 组成
//  digit 是 '1' - '9' 中的数字字符
//  digit 在 number 至少出现一次


// 输入： number = "123", digit = "3"
// 输出： "12"
// 解释： "123" 中只有一个 '3' ，删除后得到 "12"

// 输入： number = "1231", digit = "1"
// 输出： "231"
// 解释： 删除第一个 '1' 可得 "231" ，
//       删除第二个 '1' 可得 "123" ，
//       因为 231 > 123 ，所以返回 "231"

// 输入： number = "551", digit = "5"
// 输出： "51"
// 解释： 删除第一个或第二个 '5' 都会得到 "51"


// 思路： 贪心
//
//      从前往后遍历要删除的数的位置 i ，
//      如果 number[i] 就是 digit ，且 number[i] < number[i + 1] ，
//      那么删除该位形成的数必定最大。
//
//      假设 number 为 xxcaxx ，其中 c 是可删除的数，共有 3 种情况：
//          1. c == a: 则删除 c 和删除 a 形成的结果一样，可以暂时不删除
//          2. c < a: 删除后面的 c 形成的结果为 xxca_ ，删除当前的 c 形成的结果为 xxa__ 。
//                   由于 c < a ，所以必定有 xxca_ < xxa__ ，删除当前的 c 的答案必定更优
//                   必定选择删除
//          3. c > a: 删除后面的 1 形成的结果为 xx1a_ ，删除当前的 c 形成的结果为 xxa__ 。
//                   由于 c > a ，所以必定有 xxca_ > xxa__ ，删除当前的 c 的答案必定更差，
//                   必定选择不删除
//
//      如果没有满足这些条件的位置，则删除最后一个 digit 即可。
//
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 number 中全部 O(n) 个字符
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量


func removeDigit(number string, digit byte) string {
    lastIndex := 0
    // 枚举要删除的数的位置 i
    for i := range number {
        // 如果第 i 位的数是 digit 且 小于第 i + 1 位的数，
        // 则删除该位形成的数必定最大，直接返回即可
        if number[i] == digit {
            if i < len(number) - 1 && number[i] < number[i + 1] {
                return number[:i] + number[i + 1:]
            }
            lastIndex = i
        }
    }

    // 此时最后一个数必定是 digit ，直接删除即可
    return number[:lastIndex] + number[lastIndex + 1:]
}
