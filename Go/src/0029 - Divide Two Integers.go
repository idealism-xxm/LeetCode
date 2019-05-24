// 链接：https://leetcode.com/problems/divide-two-integers/
// 题意：给定两个有符号整数，不适用乘除模运算，求其整数除法结果

// 输入：10 3
// 输出：3

// 输入：7 -3
// 输出：-2

// 思路1：不断执行减法操作即可，可以用快速幂的思想优化

import "math"

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 { // 只有这种情况会溢出
        return math.MaxInt32
    }

    sign := getDivideResultSign(dividend, divisor) // 获取除法结果结果符号
    dividend, divisor = abs(dividend), abs(divisor) // 除数和被除数取绝对值
    if dividend < divisor {
        return 0
    }


    // multi[i] = divisor * 2 ^ i
    multi := []int{divisor}
    for i := 0; multi[i] <= dividend; i++ { // 计算 multi，直到第一次大于 dividend
        multi = append(multi, multi[i] << 1)
    }

    result := 0
    for i := len(multi) - 2; i >= 0; i-- { // 上面计算保证最后一个 multi > dividend，所以直接从 倒数第二个开始
        if multi[i] <= dividend { // 从大到小，保证每次减去最多的 divisor 的倍数
            dividend -= multi[i]
            result += 1 << uint(i)
        }
    }

    if sign == 1 {
        return result
    }
    return -result
}

func getDivideResultSign(a, b int) int {
    if (a > 0 && b < 0) || (a < 0 && b > 0) {
        return -1
    }
    return 1
}


func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

// 思路2：不断执行减法操作即可，可以用快速幂的思想优化
//      递归实现，简洁易懂，还可以同时计算出余数

import "math"

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 { // 只有这种情况会溢出
        return math.MaxInt32
    }

    sign := getDivideResultSign(dividend, divisor) // 获取除法结果结果符号
    dividend, divisor = abs(dividend), abs(divisor) // 除数和被除数取绝对值

    result, _ := doDivide(dividend, divisor, 1) // 递归执行

    if sign == 1 {
        return result
    }
    return -result
}

// 传入被除数，除数, 基础除数的倍数
// 返回结果和余数
func doDivide(dividend int, divisor int, cnt int) (result, remain int) {
    if dividend < divisor {
        return 0, dividend
    }

    result, remain = doDivide(dividend, divisor << 1, cnt << 1)
    if remain >= divisor { // 如果余数还比除数大，则当前还可以减去被除数
        remain -= divisor
        result += cnt
    }
    return result, remain
}

func getDivideResultSign(a, b int) int {
    if (a > 0 && b < 0) || (a < 0 && b > 0) {
        return -1
    }
    return 1
}


func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}