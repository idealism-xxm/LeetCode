// 链接：https://leetcode.com/problems/multiply-strings/
// 题意：给定两个位数小于 110 的没有前导零的非负整数，
//      求这两个数的乘积？

// 输入：num1 = "2", num2 = "3"
// 输出："6"

// 输入：num1 = "123", num2 = "456"
// 输出："56088"

// 思路：模拟
//      按照乘法的竖式计算模拟即可
//      从个位开始：第一个数的第 i 位 * 第二个数的第 j 位第结果放在第结果的第 i + j 位上（ i, j 从 0 开始，且不考虑进位）
//      然后再从结果的个位开始计算每一位的最后结果，考虑进位情况
//      且结果位数不超过 len(num1) + len(num2) 位
//      时间复杂度： O(len(num1) * len(num2)) ，空间复杂度： O(len(num1) + len(num2))
//
//      还能想到就是使用快速乘法进行操作，将乘法转变为加法
//     （不过由于两个数字都很大，不能运用位运算加速，所以时间复杂度差不多）

func multiply(num1 string, num2 string) string {
    resultLength := len(num1) + len(num2)
    result := make([]int32, resultLength, resultLength)
    for i := len(num1) - 1; i >= 0; i-- {
        a := int32(num1[i] - '0')  // num1 第 i 位转为数字
        for j := len(num2) - 1; j >= 0; j-- {
            b := int32(num2[j] - '0')  // num2 第 j 位转为数字
            result[i + j + 1] += a * b  // 由于 i, j 分别从最高位开始，所以应从结果数组最后一位开始放入
        }
    }

    for i := resultLength - 1; i > 0; i-- {  // 从个位开始计算进位后的结果，并转成对应字符的 ASCII 码
        result[i - 1] += result[i] / 10  // 计算进位
        result[i] = result[i] % 10 + '0'  // 计算当前位最终结果
    }
    result[0] += '0'

    for i := 0; i < resultLength; i++ {
        if result[i] != '0' {  // 返回不含前导 0 的结果字符串
            return string(result[i:])
        }
    }
    return string(result[resultLength - 1:])  // 若全为 0，则只返回 "0"
}

func reverse(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
