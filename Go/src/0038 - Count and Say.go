// 链接：https://leetcode-cn.com/problems/count-and-say/
// 题意：存在以下规律的字符串：后一个字符串是前一个字符串把一段连续的数字说出来的形式
//      1.      1
//      2.      11              前一个字符串有 1个1，写作：11
//      3.      21              前一个字符串有 2个1，写作：21
//      4.      1211            前一个字符串有 1个2 和 1个1，写作：1211
//      5.      111221          前一个字符串有 1个1 和 1个2 和 2个1，写作：111221
//      输入 n(1 <= n <= 30)，输出其表示的字符串

// 输入：1
// 输出："1"

// 输入：4
// 输出："1211"

// 思路：模拟即可，
//      从后往前找，每次记住最大的数 maxNum
//      找到第一个小于当前最大的数的下标 index
//      在从 index + 1 往后找比 nums[index] 大的最小数的下标 targetIndex
//      交换 nums[index] 和 nums[targetIndex]，然后对 nums[index + 1] 按升序排序
//      时间复杂度：O(nlogn)

import (
    "bytes"
    "strconv"
)

func countAndSay(n int) string {
    cur := "1"
    for n--; n > 0; n--  { // 当还没计算完时，继续通过前一个字符串计算当前字符串
        pre := cur
        length := len(pre)
        var curBuffer bytes.Buffer
        count := 1  // 计入第一个字符
        for i := 1; i < length; i++ {
            if pre[i - 1] != pre[i] { // 当前一个字符不等于当前字符时
                curBuffer.WriteString(strconv.Itoa(count))
                curBuffer.WriteByte(pre[i - 1]) // 把前一段字符“说”出来
                count = 0 // 重置统计次数
            }

            count++ // 增加相同字符的次数
        }
        curBuffer.WriteString(strconv.Itoa(count))
        curBuffer.WriteByte(pre[length - 1]) // 把最后一段字符“说”出来

        cur = curBuffer.String() // 获取字符串
    }

    return cur
}