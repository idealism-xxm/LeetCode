// 链接：https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// 题意：给定一个字符串和一个单词数组，若单词数组任意顺序连接后是字符串的子串，则记录其起始下标，求所有这样的起始下标

// 输入：s = "barfoothefoobarman", words = ["foo","bar"]
// 输出：[0,9]

// 输入： s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// 输出：[]

// 思路：双指针
//  刚开始没注意到所有单词的长度都一样，导致一直在举反例反驳自己的各种思路
//  后来没办法就开始看题解，一看第一句话就说明所有单词长度一样
//  然后瞬间就想到了如何去解答：由于长度一样，所以每次按照长度划分
//  就可以看成一个整数数组是否存在一个连续的区间，包含另一个整数集合的全部数字，且不含其他数字
//  大致做法和：0003 - Longest Substring Without Repeating Characters.go 这题类似，不过具体处理方式有点区别
//  变成字符串后，就枚举起点 [0, wordLen) ，然后进行上述操作即可
//
//  时间复杂度：O(wordLen * n)，空间复杂度：O(wordsLen * wordLen)

func findSubstring(s string, words []string) []int {
    sLen, wordsLen := len(s), len(words)
    if sLen == 0 || wordsLen == 0 || sLen < wordsLen * len(words[0]){
        return []int {}
    }

    // 统计每个单词出现次数
    wordCnt := make(map[string] int)
    for _, word := range words {
        wordCnt[word]++
    }

    wordLen := len(words[0])
    totalLen := wordsLen * wordLen
    var result []int
    // 遍历一个单词长度内所有可能的起点，后面所有的步长都是 wordLen
    for i := 0; i < wordLen; i++ {
        rMax := sLen - wordLen
        l, r := i, i // 双指针分别表示：第一个出现的单词的起始下标，最后一个出现的单词的起始下标
        for r <= rMax {
            word := s[r: r + wordLen] // 获取当前单词
            wordCnt[word]-- // 当前单词还可以使用，减去一次可使用次数（所有单词都减去，方便加回可用次数不用判断）
            r += wordLen // 移动 r

            for wordCnt[word] < 0 { // 当前单词起始不能使用时，添加第一个单词的可使用次数，并移动 l，直到当前单词可以使用
                wordCnt[s[l: l + wordLen]]++ // 添加单词可使用次数
                l += wordLen // 移动 l
            }

            if r - l == totalLen { // 若两个指针内包含了全部的单词，则记录起始下标
                result = append(result, l)
            }
        }

        for l < r { // 添加所有单词的可使用次数，并移动 l，直到 l == r，保证复原计数map
            wordCnt[s[l: l + wordLen]]++ // 添加单词可使用次数
            l += wordLen // 移动 l
        }
    }
    return result
}