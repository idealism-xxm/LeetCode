// 链接：https://leetcode.com/problems/reordered-power-of-2/
// 题意：给定一个正整数 n ，判断对 n 的十进制位重排后（不能含有前导零），
//      能否形成 2 的幂次方？


// 数据限制：
//  1 <= n <= 10 ^ 9


// 输入： n = 1
// 输出： true
// 解释： 1 = 2 ^ 0

// 输入： n = 10
// 输出： false
// 解释： 10 只能重排成 10 ，而 2 ^ 3 < 10 < 2 ^ 4


// 思路： Map
//
//      两个数是按十进制位重排的，当且仅当所有十进制数位 (0~9) 的出现次数都相同。
//
//      所以我们可以枚举 32 位整型能表示的 2 的幂次方 x (2 ^ 0, 2 ^ 1, ..., 2 ^ 30) ，
//      若 n 和 x 所有十进制数位的出现次数都相同，则 n 和 x 是按十进制重排的，
//      则 n 满足题意，直接返回 true 。
//
//
//      设十进制的不同数位个数为 C = 10 。
//
//      时间复杂度：O(logn * (logn + C))
//          1. 需要遍历全部 O(logn) 个 2 的幂次方，
//              每次遍历时都需要遍历全部 O(logn) 个十进制位
//              和全部 O(C) 个不同的十进制数位
//      空间复杂度：O(C)
//          1. 需要维护全部 O(C) 个不同十进制数位的出现次数


func reorderedPowerOf2(n int) bool {
    // 先计算 n 的十进制位出现次数，方便后续复用
    digitToCnt := getDigitToCnt(n)
    // 枚举所有 2 的幂次方
    for i := 0; i < 31; i++ {
        // 如果所有十进制数位的出现次数都相同，则 n 和 1 << i 是按十进制位重排的，
        // 则 n 满足题意，直接返回 true 。
        if isMapEqual(digitToCnt, getDigitToCnt(1 << i)) {
            return true
        }
    }

    // 此时 n 必定不满足题意
    return false
}

// 获取正整数 num 十进制数位 (0~9) 的出现次数
func getDigitToCnt(num int) map[int]int {
    digitToCnt := make(map[int]int, 10)
    for num > 0 {
        digitToCnt[num % 10] += 1
        num /= 10
    }

    return digitToCnt
}

func isMapEqual(a, b map[int]int) bool {
    // 如果长度不一样，则必定不同
    if len(a) != len(b) {
        return false
    }

    for key, val := range a {
        // 如果 b 中不存在 key ，或者对应的值不为 val ，则必定不同
        if value, exists := b[key]; !exists || val != value {
            return false
        }
    }

    return true
}