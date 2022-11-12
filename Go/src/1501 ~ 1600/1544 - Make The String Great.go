// 链接：https://leetcode.com/problems/make-the-string-great/
// 题意：给定一个字符串 s ，每次操作可以移除其中两个相邻且互为大小写的字母，
//      不断重复这个操作直至没有相邻且互为大小写的字母，返回最终字符串。


// 数据限制：
//  1 <= s.length <= 100
//  s 仅由英文大小写字母组成


// 输入： s = "leEeetcode"
// 输出： "leetcode"
// 解释： "leEeetcode" -> "leetcode"

// 输入： s = "abBAcC"
// 输出： ""
// 解释： "abBAcC" --> "aAcC" --> "cC" --> ""

// 输入： s = "s"
// 输出： "s"


// 思路： 栈
//
//      本题是 LeetCode 1047 的加强版，可移除的相邻字母从相同变成了互为大小写，
//      改变一下判定逻辑就可以直接复用。
//
//
//      我们可以直接用一个栈 stack 维护相邻但不互为大小写的字母。
//
//      然后遍历 s 中的每一个字母 ch ：
//          1. stack 为空 : 将 ch 入栈， stack 仍能保持相邻字母不互为大小写
//          2. stack.top 与 ch 不互为大小写: 将 ch 入栈， stack 仍能保持相邻字母不互为大小写
//          3. stack.top() 与 ch 互为大小写: 需要执行一次操作，以删除栈顶字母和 ch ，
//               即直接出栈，并丢弃 ch
//
//      最后 stack 中所有相邻的字母都不互为大小写，直接转成字符串返回即可。
//
//
//      时间复杂度： O(n)
//          1. 需要遍历 s 中全部 O(n) 个字母，每个字母都会入栈 1 次，最多出栈 1 次
//          2. 需要遍历 stack 中全部相邻但不互为大小写的字母，以形成结果字符串，
//              最差情况下全部 O(n) 个字母都是相邻但不互为大小写
//      空间复杂度： O(n)
//          1. 需要维护 stack 和结果字符串中全部相邻但不互为大小写的字母，
//              最差情况下全部 O(n) 个字母都是相邻但不互为大小写


func makeGood(s string) string {
    // stack 维护相邻但不互为大小写的字母
    stack := make([]rune, 0, len(s))
    // 遍历 s 中每个字母 ch
    for _, ch := range s {
        if len(stack) > 0 && needDelete(stack[len(stack) - 1], ch) {
            // 如果栈顶的字母和 ch 互为大小写，
            // 则需要执行一次操作，以删除栈顶字母和 ch ，
            // 即直接出栈，并丢弃 ch
            stack = stack[:len(stack) - 1]
        } else {
            // 此时将 ch 入栈后，仍能保持 stack 中相邻字母不互为大小写
            stack = append(stack, ch)
        }
    }

    // 转换成字符串后返回
    return string(stack)
}

func needDelete(preCh, ch rune) bool {
    // 保证 preCh <= ch ，方便后续判断
    if preCh > ch {
        preCh, ch = ch, preCh
    }

    // 只有当 ch 和 preCh 分别是同一个字母的小写和大写时，
    // 才需要删除这两个字母
    return ch - preCh == 'a' - 'A'
}
