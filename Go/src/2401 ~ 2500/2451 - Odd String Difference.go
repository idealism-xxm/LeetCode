// 链接：https://leetcode.com/problems/odd-string-difference/
// 题意：给定一个字符串数组 words ，每个字符串的长度都是 n 。
//      每一个字符串 words[i] 都能转换成一个长度为 n - 1 的差分数组 difference[i] 。
//      其中 difference[i][j] = words[i][j+1] - words[i][j] (0 <= j <= n - 2) 。
//
//      words 中除了一个字符串的差分数组不一样，其他都相同。
//      找到并返回这个差分数组不一样的字符串。


// 数据限制：
//  3 <= words.length <= 100
//  n == words[i].length
//  2 <= n <= 20
//  words[i] 仅由英文小写字母组成


// 输入： words = ["adc","wzy","abc"]
// 输出： "abc"
// 解释： "adc" 的差分数组为： [3 - 0, 2 - 3] = [3, -1]
//       "wzy" 的差分数组为： [25 - 22, 24 - 25]= [3, -1]
//       "abc" 的差分数组为： [1 - 0, 2 - 1] = [1, 1]

// 输入： words = ["aaa","bob","ccc","ddd"]
// 输出： "bob"
// 解释： "bob" 的差分数组为： [13, -13]
//       其他字符串的差分数组为： [0, 0]


// 思路： Map
//
//      我们可以用一个名为 keyToInfo 的 Map 来维护每一种差分数组的相关信息。
//      为了使用 Map ，我们需要将差分数组格式化成字符串 key ，以作为 Map 的键。
//
//      keyToInfo[key] 维护两个值，分别为 count 和 word ：
//          1. count: 能形成 key 的字符串数
//          2. word: words 中第一个形成 key 的字符串
//
//      我们可以遍历 words 中的每个字符串 word ，计算差分数组，并格式化成 key ，
//      对 keyToInfo[key] 做相应处理即可。
//
//      根据题意可知， keyToInfo 有且仅有 2 个不同的 key ，
//      所以也可以直接使用两个变量维护相关信息即可。
//
//
//      设字符串最长为 L 。
//
//      时间复杂度：O(nL)
//          1. 需要枚举全部 O(n) 个字符串，每次都需要枚举字符串全部 O(L) 个字符
//      空间复杂度：O(L)
//          1. 需要维护 2 个长度为 O(L) 的 key


func oddString(words []string) string {
    // keyToInfo[key] = (count, word) 
    //  key: 差分数组格式化形成的字符串，方便作为 Map 的键
    //  count: 能形成 key 的字符串数
    //  word: words 中第一个形成 key 的字符串
    keyToInfo := make(map[string]*Info, 2)
    for _, word := range words {
        // 计算差分，并转成字符串，方便拼接
        diff := make([]string, len(word) - 1)
        for i := 0; i < len(word) - 1; i++ {
            diff[i] = strconv.Itoa(int(word[i + 1] - word[i]))
        }
        // 将差分数组拼接成 key
        key := strings.Join(diff, "|")

        if info, exists := keyToInfo[key]; exists {
            // 能形成 key 的字符串数 +1
            info.count += 1
        } else {
            // 没有 key 对应的信息，则直接插入
            keyToInfo[key] = &Info{ 1, word }
        }
    }

    // 找到并返回差分数组只出现一次的字符串
    for _, info := range keyToInfo {
        if info.count == 1 {
            return info.word
        }
    }
    // 题目保证必定存在满足题意的字符串，所以不会走到这
    return ""
}

type Info struct {
    count int
    word string
}
