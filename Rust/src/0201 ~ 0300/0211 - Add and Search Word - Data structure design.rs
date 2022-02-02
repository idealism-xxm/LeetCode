// 链接：https://leetcode.com/problems/design-add-and-search-words-data-structure/
// 题意：实现一个数据结构，支持添加单词和搜索带 '.' 的单词，
//      搜索的单词只包含 '.' 和 a-z ， '.' 表示这一位可以匹配任意字母。

// 数据限制：
//  1 <= word.length <= 500
//  addWord 中的单词仅由英文小写字母组成
//  search 中的单词仅由英文小写字母和 '.' 组成
//  addWord 和 search 最多会被调用 50000 次

// 输入：["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
//      [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// 输出：[null,null,null,null,false,true,true,true]
// 解释：
//  addWord("bad")
//  addWord("dad")
//  addWord("mad")
//  search("pad") -> false
//  search("bad") -> true
//  search(".ad") -> true
//  search("b..") -> true

// 思路：Trie + 递归
//
//		本题是 0208 的加强版，引入了 '.' 可以匹配任意字符，
//      所以需要转成递归的 trie ，方便处理这种情况
//
//		建立 trie ，搜索的时候根据单词剩余部分的第一个字符决定相应逻辑：
//			1. '.': 遍历所有子节点递归，只有有一个子节点的结果是 true 就可以返回
//			2. a-z: 找到对应的子节点递归处理即可
//
//      遍历完单词的所有字符后，直接返回 node.is_end ：
//          1. 如果当前节点 node.is_end == true ，则说明单词存在，返回 true
//          2. 如果当前节点 node.is_end == false ，则说明单词不存在，返回 false
//
//
//      设 C 表示字符集大小，为常数 26 ， L 为单词最大长度， 
//      n 为 addWrod 调用次数， m 为 search 调用次数
//
//      时间复杂度：O(n * L + m * C ^ L)
//      空间复杂度：O(n * C * L)

// 简化创建节点
#[derive(Default)]
struct TrieNode {
    // 标记不同字母的子节点
    next: [Option<Box<TrieNode>>; 26],
    // 判断是否为一个单词终点
    is_end: bool,
}

impl TrieNode {
    pub fn search(&self, word: &[u8]) -> bool {
        // 如果已经遍历完所有字符，则只有当前节点 is_end == true 时，
        // 才存在这个单词，可以直接返回 self.is_end
        if word.len() == 0 {
            return self.is_end;
        }

        let remain = &word[1..];
        // 根据第一个字符判断是否需要遍历所有子节点
        match word[0] {
            // 如果第一个字符是 '.' ，则可以匹配任意字符，
            // 所以只要任意一个子节点中能找到就可以直接返回 true
            b'.' => self.next
                // 将子节点转成迭代器
                .iter()
                // 过滤存在的子节点
                .filter(|node| node.is_some())
                // 如果有任意一个子节点中能找到单词剩余的部分，则可以返回 true
                .any(|node| node.as_ref().unwrap().search(remain)),
            // 其余情况，第一个字符是 'a' - 'z' ，只需要找到对应子节点递归搜索即可
            _ => self.next[(word[0] - b'a') as usize]
                // 获取不可变引用
                .as_ref()
                // 如果子节点不存在，则直接返回 false
                // 如果子节点存在，则递归搜索单词剩余部分
                .map_or(false, |node| node.search(remain)),
        }
    }
}

// 简化创建节点
#[derive(Default)]
struct Trie {
    // Trie 的根节点
    root: TrieNode
}

impl Trie {
    pub fn add_word(&mut self, word: String) {
        word
            // 转成字节数组
            .as_bytes()
            // 转成迭代器处理
            .iter()
            // 使用 flod 累积
            .fold(
                // 初始节点是 trie 的根节点
                &mut self.root, 
                // 对于当前 trie 节点，返回当前字符 ch 对应的子节点
                |node, ch| 
                    // 返回 ch 对应的子节点 node.next[(ch - b'a') as usize] ，
                    // 如果不存在则创建
                    node.next[(ch - b'a') as usize].get_or_insert_with(Default::default),
            )
            // 设置最后一个子节点表明单词已结束，设置 is_end = true
            .is_end = true;
    }

    pub fn search(&self, word: String) -> bool {
        // 由于需要递归处理，我们直接转交给 trie 的根节点处理
        self.root.search(word.as_bytes())
    }
}

// 简化创建 WordDictionary
#[derive(Default)]
struct WordDictionary {
    // 使用 trie 来辅助
    trie: Trie
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl WordDictionary {

    fn new() -> Self {
        Self::default()
    }
    
    fn add_word(&mut self, word: String) {
        // 在 trie 中添加单词 word
        self.trie.add_word(word)
    }
    
    fn search(&self, word: String) -> bool {
        // 在 trie 中搜索单词 word
        self.trie.search(word)
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * let obj = WordDictionary::new();
 * obj.add_word(word);
 * let ret_2: bool = obj.search(word);
 */
