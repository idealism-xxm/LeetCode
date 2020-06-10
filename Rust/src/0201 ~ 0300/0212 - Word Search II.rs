// 链接：https://leetcode.com/problems/word-search-ii/
// 题意：给定一个字符矩阵和一些单词，求在矩阵中的所有单词？
//		（每个位置的字符最多可在一个单词中使用一次，
//        当前字符必须在下一个字符的上下左右）

// 输入：
// board = [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]
// words = ["oath","pea","eat","rain"]
// 输出：["eat","oath"]

// 思路：Trie
//
//		0079 的加强版
//
//		由于要同时搜索多个单词，所以需要建立 trie ，然后枚举矩阵的起点进行搜索即可
//      当遇到某一个节点是一个单词的终止字符时，就将对应的单词放入结果列表中
//
//      可以每次找到一个单词后将其从 trie 中删除，这样可以减少重复搜索

use std::cmp;
use std::collections::HashSet;
use std::iter::FromIterator;

static DR: [i32; 4] = [0, 1, 0, -1];
static DC: [i32; 4] = [1, 0, -1, 0];

// 简化创建节点
#[derive(Default)]
struct TrieNode {
    next: [Option<Box<TrieNode>>; 26],
    is_end: bool,
}

impl TrieNode {
    /** 从当前节点开始，搜索 word 是否存在  */
    fn search(&self, board: &Vec<Vec<char>>, used: &mut Vec<Vec<bool>>, r: usize, c: usize, word: &mut Vec<u8>, index: usize) -> HashSet<String> {
        // 结果集合，收集本层开始的单词
        let mut result: HashSet<String> = HashSet::new();

        // 获取对应的字符
        let ch = board[r][c] as u8;
        // 计算对应的子节点
        let i = (ch - b'a') as usize;
        // 如果该节点不存在，则直接返回空集合
        if self.next[i].is_none() {
            return result;
        }
        // 节点存在，获取对应的节点
        let cur = self.next[i].as_ref().unwrap();

        // 当前字符可以放入 word 中
        word[index] = ch;

        // 如果当前节点是一个单词的结束字符，则收集对应的单词
        if cur.is_end {
            result.insert(String::from_utf8(word[..=index].to_vec()).unwrap());
        }

        // 继续收集后续可能产生的单词
        let m = board.len() as i32;
        let n = board[0].len() as i32;
        for (dr, dc) in DR.iter().zip(DC.iter()) {
            // 计算下一个字符的位置
            let rr = (r as i32) + *dr;
            let cc = (c as i32) + *dc;
            // 如果位置不合法，则直接处理下一个
            if !TrieNode::is_valid(m, n, rr, cc) {
                continue;
            }
            // 如果未使用
            let rr = rr as usize;
            let cc = cc as usize;
            if !used[rr][cc] {
                // 标记该位置已使用
                used[rr][cc] = true;
                result = result.union(&cur.search(board, used, rr, cc, word, index + 1)).cloned().collect();
                // 取消使用标记
                used[rr][cc] = false;
            }
        }
        result
    }

    fn is_valid(m: i32, n: i32, r: i32, c: i32) -> bool {
        0 <= r && r < m && 0 <= c && c < n
    }
}

// 简化创建 Trie
#[derive(Default)]
struct Trie {
    root: TrieNode
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Trie {
    /** Initialize your data structure here. */
    fn new() -> Self {
        Self::default()
    }

    /** Adds a word into the data structure. */
    fn add_word(&mut self, word: &String) {
        let mut cur = &mut self.root;
        // 遍历 word 每一个字符
        for ch in word.bytes() {
            // 计算 next 数组中的下标
            let i = (ch - b'a') as usize;
            // 找到该字符对应的节点
            let next = &mut cur.next[i];
            // 若 next 不存在，则创建新的节点。
            // 返回目前 next 包含的节点的可变引用
            cur = next.get_or_insert_with(Box::<TrieNode>::default);
        }
        // 标记最后一个字符对应的节点为单词的结束节点
        cur.is_end = true;
    }

    fn search(&self, board: &Vec<Vec<char>>, word_max_len: usize) -> Vec<String> {
        // 初始化指定容量的单词 Vec
        let mut word = vec![0; word_max_len];
        // 标记字符是否已经使用过
        let mut used = vec![vec![false; board[0].len()]; board.len()];
        // 枚举矩阵的起点，搜索单词并收集结果
        let mut result: HashSet<String> = HashSet::new();
        for i in 0..board.len() {
            for j in 0..board[0].len() {
                // 标记该位置已使用
                used[i][j] = true;
                // 交给根节点进行搜索
                result = result.union(&self.root.search(board, &mut used, i, j, &mut word, 0)).cloned().collect();
                // 取消使用标记
                used[i][j] = false;
            }
        }

        // 结果集合转换为结果列表返回
        Vec::from_iter(result.into_iter())
    }
}

impl Solution {
    pub fn find_words(board: Vec<Vec<char>>, words: Vec<String>) -> Vec<String> {
        // 记录单词中最大的长度，方便后续处理
        let mut word_max_len: usize = 0;
        // 建立 trie
        let mut trie = Trie::new();
        words.iter().for_each(|word| {
            word_max_len = cmp::max(word_max_len, word.len());
            trie.add_word(word);
        });

        // 使用 trie 进行搜索
        trie.search(&board, word_max_len)
    }
}
