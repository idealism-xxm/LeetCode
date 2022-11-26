// 链接：https://leetcode.com/problems/word-search/
// 题意：给定一个字符矩阵，求一个单词 word 是否出现在其中？
//      （每个位置的字符最多可使用一次，当前字符必须在下一个字符的上下左右）


// 数据限制：
//  m == board.length
//  n = board[i].length
//  1 <= m, n <= 6
//  1 <= word.length <= 15
//  board 和 word 仅含有英文大小写字母


// 输入： board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出： true
// 解释： |A||B||C| E
//        S  F |C| S
//        A |D||E| E

// 输入： board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出： true
// 解释：  A  B  C  E
//        S  F  C |S|
//        A  D |E||E|

// 输入： board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// 输出： false


// 思路： 递归/回溯/DFS
//
//      因为 word 在 board 中出现的方向没有任何限制，
//      所以只能使用递归/回溯/DFS 的方式遍历全部可能的路径。
//
//      首先，要进行剪枝，对绝对找不到的情况直接返回 false 。
//          1. board 不足以容纳 word 中全部字符，则直接返回 false
//          2. word 中某种字符的数量在 board 中不够，则直接返回 false
//
//      其次，要进一步优化： word 可以从左边或者右边搜，选择选择搜索空间更小的一端。
//
//      然后，需要遍历 baord 中每个位置 (r, c) 作为起点。
//
//      每次使用 dfs(board, word, cur, r, c) 进行处理，
//      判断在 board 在中以 (r, c) 为起点是否能找到 word 。
//          1. board: 题目给定的字符矩阵，同时复用记录 board 记录使用状态
//          2. word: 题目给定的单词，透传即可
//          3. cur: word 下一个该匹配的字符下标
//          4. r: board 中下一个该匹配的字符位置行坐标
//          5. c: board 中下一个该匹配的字符位置列坐标
//
//      在 dfs 中，按照如下流程处理即可：
//          1. 如果 board[r][c] != word[cur] ，则匹配失败，直接返回 false
//          2. 如果 cur == word.len() ，则最后一个字符匹配成功，直接返回 true
//          3. 遍历计算全部 4 个方向的下一个位置 (rr, cc) ，位置合法时，
//              递归调用 dfs(board, word, cur + 1, rr, cc) 。
//              如果递归返回 true ，则直接返回 true
//          4. 遍历完 4 个方向后还未返回时，则无法继续组成 word ，直接返回 false
//
//
//      设 word 长度为 l 。
//
//      时间复杂度： O(mn * 3 ^ min(mn, l))
//          1. 需要用 board 全部 O(mn) 个位置作为起点。
//              每次都需要使用 dfs 递归，递归每一层都能选 3 个方向（不能往回走），
//              递归共 O(min(mn, l)) 层，递归时间复杂度为 O(3 ^ min(mn, l)) 。
//      空间复杂度： O(mn + l)
//          1. 不复用 board 的话，需要维护 board 全部 O(mn) 个位置的使用状态
//          2. 需要维护递归栈，最差情况下需要递归 O(min(mn, l)) 层
//          3. 翻转 word 时需要维护全部 O(l) 个字符


// 每个方向的位置改变量
//  0: 向上走 1 步
//  1: 向右走 1 步
//  2: 向下走 1 步
//  3: 向左走 1 步
var DR = []int{-1, 0, 1, 0}
var DC = []int{0, 1, 0, -1}


func exist(board [][]byte, word string) bool {
	// 先进行剪枝和优化
	wordPtr := prune(board, word)
	// 如果 wordPtr 为空，则说明需要剪枝，直接返回 false
	if wordPtr == nil {
		return false
	}
	word = *wordPtr

	// 尝试以每一个位置为起点，有一个符合题意则直接返回 true
	for r := range board {
		for c := range board[0] {
			if dfs(board, word, 0, r, c) {
				return true
			}
		}
	}
	// 此时所有起点都无法找到 word ，返回 false
	return false
}

// 剪枝和优化，如果返回空，则说明需要剪枝；否则 word 就是搜索空间更小的
func prune(board [][]byte, word string) *string {
	m, n := len(board), len(board[0])
	// 如果 board 不足以容纳 word 中全部字符，则直接返回空
	if m * n < len(word) {
		return nil
	}
	
	// 统计 board 和 word 每种字符的数量
	cellToCnt := make(map[byte]int)
	for r := range board {
		for c := range board[r] {
			cellToCnt[board[r][c]] += 1
		}
	}
    chToCnt := make(map[byte]int)
	for i := range word {
		chToCnt[word[i]] += 1
	}
	// 如果 word 中某种字符的数量在 board 中不够，则直接返回空
    for ch, cnt := range chToCnt {
        if cellToCnt[ch] < cnt {
			return nil
        }
    }
	
	// 选择搜索空间更小的一端
    l, r := 0, len(word) - 1
	for l < r {
        lcnt, rcnt := chToCnt[word[l]], chToCnt[word[r]]
		if lcnt == rcnt {
			// 如果两端的字符出现数相等，则继续向内判断
			l += 1
			r -= 1
		} else {
			// 如果左侧的字符数出现更多，则选择以右边为起点，搜索空间更小
			if lcnt > rcnt {
                word = reverse(word)
			}

			break
		}
	}

	return &word
}

func dfs(board [][]byte, word string, cur, r, c int) bool {
	// 如果当前字符不匹配，则直接返回 false
	if board[r][c] != word[cur] {
		return false
	}
	// 如果 word 的最后一个字符已匹配成功，则已找到 word ，返回 true
	if cur == len(word) - 1 {
		return true
	}

	// 记录当前位置的原始字符，并标记当前位置的字符已使用
	originCh := board[r][c]
	board[r][c] = '#'

	m, n := len(board), len(board[0])
	// 遍历 4 个方向对应的偏移
	for i := range DR {
		// 计算该方向下一个位置的坐标
		rr, cc := r + DR[i], c + DC[i]
		// 如果下标在范围内，且后续的字符串满足题意，则直接返回 true
		if isValid(m, n, rr, cc) && 
			dfs(board, word, cur + 1, rr, cc) {
			return true
		}
	}

	// 复原当前位置的字符
	board[r][c] = originCh
	// 此时，所有方向都无法继续组成 word ，返回 false
	return false
}

// 判断 (r, c) 是否为 m * n 矩阵的合法坐标
func isValid(m, n, r, c int) bool {
	return 0 <= r && r < m && 0 <= c && c < n
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
