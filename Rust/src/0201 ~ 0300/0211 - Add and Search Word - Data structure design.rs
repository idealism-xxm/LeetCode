// 链接：https://leetcode.com/problems/add-and-search-word-data-structure-design/
// 题意：实现一个数据结构，支持添加单词和搜索带正则的单词，搜索的单词只包含 . 和 a-z ？

// addWord("bad")
// addWord("dad")
// addWord("mad")
// search("pad") -> false
// search("bad") -> true
// search(".ad") -> true
// search("b..") -> true

// 思路：Trie
//
//		0208 的加强版
//
//		建立 trie ，然后搜索的时候改用递归即可，
//			1. 遇到 . 时遍历所有子节点递归，
//			2. 遇到 a-z 时找到对应的子节点递归

// 简化创建节点
#[derive(Default)]
struct Node {
    next: [Option<Box<Node>>; 26],
    is_end: bool,
}

impl Node {
    /** 从当前节点开始，搜索 word 是否存在  */
    fn search(&self, word: &[u8]) -> bool {
        // 如果已经遍历完所有字符，则返回 self.is_end
        if word.len() == 0 {
            return self.is_end;
        }

        // 第一个字符，用于判断
        let ch = word[0];
        let remain = &word[1..];
        // 如果是通配符，则遍历所有存在的子节点递归查找
        if ch == '.' as u8 {
			// 遍历所有子节点
            self.next.iter()
				// 只保留存在的子节点
				.filter(|next| next.is_some())
				// 这些节点只要有一个能完成搜索剩余的字符即可返回 true
				.any(|next| next.as_ref().unwrap().search(remain))
        } else {
			// 如果是字母，则找到对应的节点判断即可
			let i = (ch - b'a') as usize;
			// 获取对应的子节点
			let next = &self.next[i];
			// 当该节点存在 且 能完成搜索剩余的字符，则返回 true
			next.is_some() && next.as_ref().unwrap().search(remain)
		}
    }
}

// 简化创建 WordDictionary
#[derive(Default)]
struct WordDictionary {
    root: Node
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl WordDictionary {
    /** Initialize your data structure here. */
    fn new() -> Self {
        Self::default()
    }

    /** Adds a word into the data structure. */
    fn add_word(&mut self, word: String) {
        let mut cur = &mut self.root;
        // 遍历 word 每一个字符
        for ch in word.bytes() {
            // 计算 next 数组中的下标
            let i = (ch - b'a') as usize;
            // 找到该字符对应的节点
            let mut next = &mut cur.next[i];
            // 若 next 不存在，则创建新的节点。
            // 返回目前 next 包含的节点的可变引用
            cur = next.get_or_insert_with(Box::<Node>::default);
        }
        // 标记最后一个字符对应的节点为单词的结束节点
        cur.is_end = true;
    }

    /** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
    fn search(&self, word: String) -> bool {
        // 交给根节点进行搜索
        self.root.search(word.as_bytes())
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * let obj = WordDictionary::new();
 * obj.add_word(word);
 * let ret_2: bool = obj.search(word);
 */
