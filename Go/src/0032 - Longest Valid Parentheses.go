// 链接：https://leetcode.com/problems/longest-valid-parentheses/
// 题意：给定一个只含有 '(' 和 ')' 字符的字符串，求最长的合法括号子串的长度

// 输入：(()
// 输出：2

// 输入：)()())
// 输出：4

// 思路1：栈 + DP
//  很容易就能想到：可以按照栈的方式进行匹配前后括号
//  dp[i] 表示以 后括号s[i] 结尾的最长合法括号子串的长度（前括号记录无意义，默认为 0）
//  1. 当前字符是前括号，直接入栈，并记录下标，dp[i] = 0
//  2. 当前字符是后括号
//      (1) 栈顶是前括号（下标为 j）：匹配成功，栈顶元素出栈
//          则其长度等于匹配成功的这一段长度 (i - j + 1) 加上 以 s[j] 结尾的最长合法括号子串的长度 dp[j - 1]
//          即：dp[i] =  (i - j + 1) + dp[j - 1]
//      (2) 栈顶是后括号：匹配失败，后括号入栈，并记录下标，dp[i] = 0
//  时间复杂度：O(n)，空间复杂度：O(n)

func longestValidParentheses(s string) int {
    length := len(s)
    top := -1
    index := make([]int, length) // 下标，放入栈中元素的下标
    dp := make([]int, length) // dp[i] 表示以 后括号s[i] 结尾的最长合法括号子串的长度（前括号记录无意义，默认为 0）

    result := 0
    for i := 0; i < length; i++ {
        if top == -1 || s[i] == '(' { // 栈空 或者 前括号，则直接入栈
            top++
            index[top] = i
            continue
        }

        // 到这必定是后括号，且栈不为空
        if s[index[top]] == '(' { // 栈顶是前括号
            j := index[top]
            top-- // 栈顶元素出栈
            preLen := 0 // dp[j - 1]
            if j > 0 {
                preLen = dp[j - 1]
            }
            dp[i] = i - j + 1 + preLen
            result = max(result, dp[i]) // 更新最长值
        } else { // 栈顶是后括号，直接入栈
            top++
            index[top] = i
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 思路2：栈
//  看到题解中还可以只用栈，就又继续挖掘了一下以利用的信息
//  继续思考栈的运行流程，可以发现如果当前后口号匹配到栈顶的前括号
//  则出栈后的栈顶，必定是最长合法括号子串的前一个字符，不需要再额外用dp数组记录
//  同时，先放入 -1 表名第一个字符的前一个字符下标，方便统一处理
//  每次出栈后，拿栈顶下标 j 和 当前后括号下标 i 计算最长长度 i - j 即可
//  时间复杂度：O(n)，空间复杂度：O(n)

func longestValidParentheses(s string) int {
    length := len(s)
    index := make([]int, length + 1) // 下标，放入栈中元素的下标
    index[0] = -1
    top := 0

    result := 0
    for i := 0; i < length; i++ {
        if top == 0 || s[i] == '(' { // 栈空 或者 前括号，则直接入栈
            top++
            index[top] = i
            continue
        }

        // 到这必定是后括号，且栈不为空
        if s[index[top]] == '(' { // 栈顶是前括号
            top-- // 栈顶元素出栈
            j := index[top] // 获取最长合法括号子串的前一个字符的下标
            result = max(result, i - j) // 更新最长值
        } else { // 栈顶是后括号，直接入栈
            top++
            index[top] = i
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 思路3：DP
//  看到题解中还可以只用DP，在这方便继续挖掘了一下可以利用的信息
//  继续思考运行流程，可以发现思路1中的栈和DP中存在冗余的信息，所以可以使用DP数组也可以得到栈顶的信息
//  dp[i] 表示以 后括号s[i] 结尾的最长合法括号子串的长度（前括号记录无意义，默认为 0）
//  1. 当前字符是前括号：dp[i] = 0
//  2. 当前字符是后括号
//      (1) 前一个字符是前括号：匹配成功，dp[i] = dp[i - 2] + 2
//      (2) 前一个字符是后括号：s[i - dp[i]] 表示前一个字符匹配的最长合法括号子串的前一个字符
//          ① 若 s[i - dp[i]] 是前括号：匹配成功，dp[i] = dp[i - 1] + dp[i - dp[i] - 2] + 2
//          ② 若 s[i - dp[i]] 是后括号：则需要重复 2.(2)，由于子问题已处理这部分内容，所以不用处理
//  时间复杂度：O(n)，空间复杂度：O(n)

func longestValidParentheses(s string) int {
    length := len(s)
    dp := make([]int, length) // dp[i] 表示以 后括号s[i] 结尾的最长合法括号子串的长度（前括号记录无意义，默认为 0）

    result := 0
    for i := 1; i < length; i++ {
        if s[i] == '(' { // 前括号 则不处理
            continue
        }

        if s[i - 1] == '(' { // 前一个字符是 前括号：匹配成功
            if i - 2 >= 0 {
                dp[i] = dp[i - 2] + 2
            } else {
                dp[i] = 2
            }
        } else { // 前一个字符是 后括号
            j := i - 1 // 前一个字符下标
            if j - dp[j] >= 0 && s[j - dp[j]] == '(' { // 前一个字符匹配的最长合法括号子串的前一个字符是前括号：匹配成功
                if j - dp[j] - 1 >= 0 {
                    dp[i] = dp[j] + dp[j - dp[j] - 1] + 2
                } else {
                    dp[i] = dp[j] + 2
                }
            }
        }
        result = max(result, dp[i])
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 思路4：计数
//  题解最后还有一种更优的思路，挖掘到了更深层次的信息
//  继续思考使用栈时的运行流程，可以发现：
//  1. 每次更新答案的时候，必有这段左括号数量等于右括号数量
//  2. 若一个后括号没有匹配成功，则之前的必定不会时答案的一部分
//      即多余的左括号可以继续用来匹配右括号，而多余的右括号必定划分所有合法的匹配子串
//
//  因此，可以只记录左括号和右括号的数量，记为：left 和 right
//  1. 每次循环，先判断当前字符：若当前字符是左括号，则 left++；若当前字符是右括号，则 right++
//  2. 然后判断 left 和 right 的大小：
//      (1) left <  right：存在多余的右括号，重置 left 和 right 为 0
//      (2) left == right：刚好匹配成功，合法子串长度为：left + right
//  当然：只从左遍历一遍会存在问题，比如：(((((() 这样的串就无法得出正确答案，所以还需要从右遍历一遍
//  时间复杂度：O(n)，空间复杂度：O(1)

func longestValidParentheses(s string) int {
    length := len(s)

    left, right := 0, 0
    result := 0
    for i := 0; i < length; i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }

        if left == right {
            result = max(result, left << 1)
        } else if left < right {
            left, right = 0, 0
        }
    }

    left, right = 0, 0
    for i := length - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }

        if left == right {
            result = max(result, right << 1)
        } else if left > right {
            left, right = 0, 0
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}